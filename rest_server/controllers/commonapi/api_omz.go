package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
	"github.com/labstack/echo"
)

func GetOMZAirDropInfo(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// redis 확인 후 없으면 db load
	if airDropInfo, err := model.GetDB().CacheOMZGetAirDropInfo(); err != nil {
		if airDropInfo, err := model.GetDB().USPE_Get_AirDrop(); err != nil {
			resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_AirdropInfo)
		} else {
			if missions, err := model.GetDB().USPE_GetList_AccountAirDropMissions(0); err != nil {
				resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
			} else {
				airDropInfo.Missions = missions
				resp.Value = airDropInfo
				model.GetDB().CacheOMZSetAirDropInof(airDropInfo)
			}
		}
	} else {
		resp.Value = airDropInfo
	}

	return c.JSON(http.StatusOK, resp)
}

func GetOMZMyMission(ctx *context.OnbuffEventContext, req *context.ReqOMZMyMisssion) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if myMissions, err := model.GetDB().CacheOMZGetMyMissions(req.AUID); err != nil {
		if myMissions, err := model.GetDB().USPE_GetList_AccountAirDropMissions(req.AUID); err != nil {
			resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
		} else {
			resp.Value = myMissions
			model.GetDB().CacheOMZSetMyMissions(req.AUID, myMissions)
		}
	} else {
		resp.Value = myMissions
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostOMZClaimAirDrop(ctx *context.OnbuffEventContext, req *context.ReqOMZClaimAirDrop) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}
