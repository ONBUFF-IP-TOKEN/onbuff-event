package model

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/baseutil/datetime"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
)

func (o *DB) PublishEvent(channel string, val interface{}) error {
	msg, _ := json.Marshal(val)
	return o.Cache.GetDB().Publish(MakePubSubKey(channel), string(msg))
}

func (o *DB) ListenSubscribeEvent(bKeepAlive bool, termSec int64) error {
	receiveCh := make(chan basedb.PubSubMessageV8)
	defer close(receiveCh)

	channel := MakePubSubKey(InternalCmd)
	log.Infof("ListenSubscribeEvent() has been started channel : %v", channel)
	rch, err := o.Cache.GetDB().Subscribe(receiveCh, channel)
	if err != nil {
		log.Error(err)
		return err
	}

	defer func(ch string, bKeepAlive bool, termSec int64) {
		o.Cache.GetDB().Unsubscribe(ch)
		o.Cache.GetDB().ClosePubSub()
		if recver := recover(); recver != nil {
			log.Error("Recoverd in listenPubSubEvent()", recver)
		}
		go o.ListenSubscribeEvent(bKeepAlive, termSec)
	}(channel, bKeepAlive, termSec)

	if bKeepAlive {
		go func() {
			ticker := time.NewTicker(time.Duration(termSec) * time.Second)

			for {
				msg := &PSHealthCheck{
					PSHeader: PSHeader{
						Type: PubSub_cmd_healthcheck,
					},
				}
				msg.Value.Timestamp = datetime.GetTS2MilliSec()

				if err := o.PublishEvent(InternalCmd, msg); err != nil {
					log.Errorf("pubsub health check err : %v", err)
				}
				<-ticker.C
			}

		}()
	}

	for {
		msg, ok := <-rch
		if msg == nil || !ok {
			log.Errorf("redis pubsub channel rev faile [ok:%v]", ok)
			break
			//continue
		}

		if strings.Contains(msg.Channel, MakePubSubKey(InternalCmd)) {
			o.PubSubCmdByInternal(msg)
		}

		log.Debugf("subscribe channel: %v, val: %v", msg.Channel, msg.Payload)
	}

	return nil
}

func (o *DB) PubSubCmdByInternal(msg basedb.PubSubMessageV8) error {

	header := &PSHeader{}
	json.Unmarshal([]byte(msg.Payload), header)

	log.Infof("pubsub cmd : %v", header.Type)

	switch header.Type {
	case PubSub_cmd_healthcheck:
		o.PubSubCmd_HealthCheck(msg)
	}

	return nil
}

func (o *DB) PubSubCmd_HealthCheck(msg basedb.PubSubMessageV8) {
	psPacket := &PSHealthCheck{}
	json.Unmarshal([]byte(msg.Payload), psPacket)
	log.Infof("pubsub healthcheck : %v ", psPacket.Value.Timestamp)
}
