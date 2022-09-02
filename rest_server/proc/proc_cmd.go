package proc

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/basenet"
)

const (
	MarketCmd_SendNFT uint32 = 0
)

type Cmd struct {
	procMgr *ProcManager
	command chan *basenet.CommandData
}

func NewCmd(procMgr *ProcManager) *Cmd {
	tokenCmd := new(Cmd)
	tokenCmd.procMgr = procMgr

	tokenCmd.command = make(chan *basenet.CommandData)
	return tokenCmd
}

func (o *Cmd) GetMarketCmdChannel() chan *basenet.CommandData {
	return o.command
}

func (o *Cmd) GetProc(data *basenet.CommandData) base.BaseResponse {
	if ch, exist := GetChanInstance().Get(ChannelName); exist {
		ch.(chan *basenet.CommandData) <- data
	}

	if data.Callback == nil {
		return base.BaseResponse{}
	}

	ticker := time.NewTicker(90 * time.Second)

	resp := base.BaseResponse{}
	select {
	case callback := <-data.Callback:
		ticker.Stop()
		msg, ok := callback.(*base.BaseResponse)
		if ok {
			resp = *msg
		}
	case <-ticker.C:
		ticker.Stop()
		resp = base.BaseResponseInternalServerError()
	}

	return resp
}

func (o *Cmd) StartCommand() {
	GetChanInstance().Put(ChannelName, o.command)

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)

		defer func() {
			ticker.Stop()
		}()

		for {
			select {
			case ch := <-o.command:
				o.CommandProc(ch)
			case <-ticker.C:
			}
		}
	}()
}

func (o *Cmd) CommandProc(data *basenet.CommandData) error {

	if data.Data != nil {
		//start := time.Now()
		switch data.CommandType {
		case MarketCmd_SendNFT:
			o.SendNFT(data.Data, data.Callback)
		}

		//end := time.Now()

		//log.Debug("cmd.kind:", data.CommandType, ",elapsed", end.Sub(start))
	}
	return nil
}

func (o *Cmd) SendNFT(data interface{}, cb chan interface{}) {
	//model.GetDB().AddAccountAuthLog(data.(*context.AccountAuthLog))
}
