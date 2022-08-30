package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/datetime"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
)

func PostPSHealthCheck(ctx *context.OnbuffEventContext, req *context.ReqPSHealthCheck) error {
	resp := new(base.BaseResponse)
	resp.Success()

	msg := &model.PSHealthCheck{
		PSHeader: model.PSHeader{
			Type: model.PubSub_cmd_healthcheck,
		},
	}
	if req.Timestamp == 0 {
		msg.Value.Timestamp = datetime.GetTS2MilliSec()
	} else {
		msg.Value.Timestamp = req.Timestamp
	}

	if err := model.GetDB().PublishEvent(model.InternalCmd, msg); err != nil {
		log.Errorf("PublishEvent %v, type : %v, error : %v", model.InternalCmd, model.PubSub_cmd_healthcheck, err)
		resp.SetReturn(resultcode.Result_PubSub_InternalErr)
	}

	resp.Value = &context.ResPSHealthCheck{
		Timestamp: msg.Value.Timestamp,
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}
