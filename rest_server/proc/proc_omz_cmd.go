package proc

import (
	"encoding/json"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/market"
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/basenet"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_market_server"
)

const (
	OMZCmd_NFT_Transfer = iota
)

type Cmd struct {
	procMgr *ProcManager
	command chan *basenet.CommandData
}

func NewCmd(procMgr *ProcManager) *Cmd {
	cmd := new(Cmd)
	cmd.procMgr = procMgr

	cmd.command = make(chan *basenet.CommandData)
	return cmd
}

func (o *Cmd) GetCmdChannel() chan *basenet.CommandData {
	return o.command
}

func (o *Cmd) GetProc(data *basenet.CommandData) base.BaseResponse {
	if ch, exist := GetChanInstance().Get(OMZChannel); exist {
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
	GetChanInstance().Put(OMZChannel, o.command)

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)

		defer func() {
			ticker.Stop()
		}()

		for {
			select {
			case ch := <-o.command:
				o.CommandProc(ch)
				log.Infof("NFT Send reserved Queue Size : %v", len(o.command))
			case <-ticker.C:
			}
		}
	}()
}

func (o *Cmd) CommandProc(data *basenet.CommandData) error {

	if data.Data != nil {
		start := time.Now()
		switch data.CommandType {
		case OMZCmd_NFT_Transfer:
			o.OMZNFTTransfer(data.Data, data.Callback)
		}

		end := time.Now()
		log.Infof("cmd.kind:", data.CommandType, ",elapsed", end.Sub(start))
	}
	return nil
}

func (o *Cmd) OMZNFTTransfer(data interface{}, cb chan interface{}) {
	chanData := data.(*context.OMZ_NFTTransfer)

	// market send request
	params := market.ReqPostEventNFTTransfer{
		EventType: "OMZ",
		AUID:      chanData.AUID,
	}

	if res, err := inno_market_server.GetInstance().PostEventNFTTransfer(&params); err != nil {
		log.Errorf("inno_market PostEventNFTTransfer err : %v", err)
	} else {
		if res.Return != 0 {
			log.Errorf("inno_market response return : %v, message : %v", res.Return, res.Message)
		} else {
			jsonbytes, _ := json.Marshal(res.Value.NFTLists)
			log.Infof("auid:%v , claimquantity:%v,  nftlists:%v", params.AUID, res.Value.ClaimQuantity, string(jsonbytes))

			// meta redis delete
			model.GetDB().CacheOMZDelAirDropInof()
		}
	}
}
