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
			if missionsLst, winningQuantity, err := model.GetDB().USPE_GetList_AccountAirDropMissions(0); err != nil {
				resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
			} else {
				airDropInfo.Missions = &context.ResOMZMyMission{
					WinningQuantity: winningQuantity,
					MyMission:       missionsLst,
				}
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

	if resMyMission, err := model.GetDB().CacheOMZGetMyMissions(req.AUID); err != nil {
		if missionLst, winningQuantity, err := model.GetDB().USPE_GetList_AccountAirDropMissions(req.AUID); err != nil {
			resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
		} else {
			resMyMission := &context.ResOMZMyMission{
				WinningQuantity: winningQuantity,
				MyMission:       missionLst,
			}
			resp.Value = resMyMission
			model.GetDB().CacheOMZSetMyMissions(req.AUID, resMyMission)
		}
	} else {
		resp.Value = resMyMission
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostOMZClaimAirDrop(ctx *context.OnbuffEventContext, req *context.ReqOMZClaimAirDrop) error {
	resp := new(base.BaseResponse)
	resp.Success()

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PutOMZAirDropMission(ctx *context.OnbuffEventContext, req *context.ReqOMZAirDropMission) error {
	resp := new(base.BaseResponse)
	resp.Success()

	if err := model.GetDB().USPE_Add_AccountAirDropMissions(req.AUID, req.MissionID); err != nil {
		resp.SetReturn(resultcode.Result_Error_DB_OMZ_Add_AccountAirDropMissions)
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}
