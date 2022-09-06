package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/labstack/echo"
)

func (o *ExternalAPI) GetOMZAirDropInfo(c echo.Context) error {
	return commonapi.GetOMZAirDropInfo(c)
}

func (o *ExternalAPI) GetOMZMyMission(c echo.Context) error {
	ctx := base.GetContext(c).(*context.OnbuffEventContext)

	params := context.NewReqMyMission()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}
	if err := params.CheckValidate(ctx); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.GetOMZMyMission(ctx, params)
}

func (o *ExternalAPI) PostOMZClaimAirDrop(c echo.Context) error {
	ctx := base.GetContext(c).(*context.OnbuffEventContext)

	params := context.NewReqOMZClaimAirDrop()
	if err := ctx.EchoContext.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}
	if err := params.CheckValidate(ctx); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return commonapi.PostOMZClaimAirDrop(ctx, params)
}
