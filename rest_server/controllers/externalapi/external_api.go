package externalapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/auth"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/commonapi"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/labstack/echo"
)

type ExternalAPI struct {
	base.BaseController

	conf    *config.ServerConfig
	apiConf *baseconf.APIServer
	echo    *echo.Echo
}

func PreCheck(c echo.Context) base.PreCheckResponse {
	conf := config.GetInstance()
	if err := base.SetContext(c, &conf.Config, context.NewOnbuffEventContext); err != nil {
		log.Error(err)
		return base.PreCheckResponse{
			IsSucceed: false,
		}
	}

	// auth token 검증

	if conf.Auth.AuthEnable {
		author, ok := c.Request().Header["Authorization"]
		if !ok {
			// auth token 오류 리턴
			res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)

			return base.PreCheckResponse{
				IsSucceed: false,
				Response:  res,
			}
		}

		if ret, value, err := auth.CheckAuthToken(author[0][7:]); err != nil || !ret {
			res := base.MakeBaseResponse(resultcode.Result_Auth_InvalidJwt)
			return base.PreCheckResponse{
				IsSucceed: false,
				Response:  res,
			}
		} else {
			base.GetContext(c).(*context.OnbuffEventContext).SetVerifyAuthToken(value)
			log.Debugf("from : [companyid:%v][appid:%v][logintype:%v]", value.CompanyID, value.AppID, value.LoginType)
		}
	}

	return base.PreCheckResponse{
		IsSucceed: true,
	}
}

func (o *ExternalAPI) Init(e *echo.Echo) error {
	o.echo = e
	o.BaseController.PreCheck = PreCheck

	if err := o.MapRoutes(o, e, o.apiConf.Routes); err != nil {
		return err
	}

	// serving documents for RESTful APIs
	if o.conf.OnbuffEvent.APIDocs {
		e.Static("/docs", "docs/ext")
	}

	return nil
}

func (o *ExternalAPI) GetConfig() *baseconf.APIServer {
	o.conf = config.GetInstance()
	o.apiConf = &o.conf.APIServers[1]
	return o.apiConf
}

func NewAPI() *ExternalAPI {
	return &ExternalAPI{}
}

func (o *ExternalAPI) GetHealthCheck(c echo.Context) error {
	return commonapi.GetHealthCheck(c)
}

func (o *ExternalAPI) GetVersion(c echo.Context) error {
	return commonapi.GetVersion(c, o.BaseController.MaxVersion)
}

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
