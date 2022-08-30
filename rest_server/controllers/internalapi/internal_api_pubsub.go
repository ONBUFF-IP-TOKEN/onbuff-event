package internalapi

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/labstack/echo"
)

func (o *InternalAPI) PostPSHealthCheck(c echo.Context) error {
	ctx := base.GetContext(c).(*context.OnbuffEventContext)

	params := context.NewReqPSHealthCheck()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	return commonapi.PostPSHealthCheck(ctx, params)
}
