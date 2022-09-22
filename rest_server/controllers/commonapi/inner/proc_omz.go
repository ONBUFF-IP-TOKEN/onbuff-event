package inner

import (
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/basenet"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/proc"
)

func OMZProc(data *basenet.CommandData) base.BaseResponse {
	if ch, exist := proc.GetChanInstance().Get(proc.OMZChannel); exist {
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
