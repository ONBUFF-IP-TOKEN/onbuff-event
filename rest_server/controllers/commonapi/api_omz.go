package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/basenet"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/commonapi/inner"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/proc"
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
			if missionsLst, isClaimed, err := model.GetDB().USPE_GetList_AccountAirDropMissions(0); err != nil {
				resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
			} else {
				airDropInfo.Missions = &context.ResOMZMyMission{
					IsClaimed: isClaimed,
					MyMission: missionsLst,
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

	if missionLst, isClaimed, err := model.GetDB().USPE_GetList_AccountAirDropMissions(req.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
	} else {
		resMyMission := &context.ResOMZMyMission{
			IsClaimed: isClaimed,
			MyMission: missionLst,
		}
		resp.Value = resMyMission
	}

	return ctx.EchoContext.JSON(http.StatusOK, resp)
}

func PostOMZClaimAirDrop(ctx *context.OnbuffEventContext, req *context.ReqOMZClaimAirDrop) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// 1. 사용자 미션 정보 수집
	// 2. claim 진행 여부 판단
	// 3. market에 nft 전송 요청

	// 1. 사용자 미션 정보 수집
	if missionLst, isClaimed, err := model.GetDB().USPE_GetList_AccountAirDropMissions(req.AUID); err != nil {
		resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_Mission)
	} else {
		// 2. claim 진행 여부 판단
		// 2-1. alreay claimed
		if isClaimed {
			resp.SetReturn(resultcode.Result_Claimed_Already)
		} else {
			// 2-2. No winning information
			findWin := false
			for _, mission := range missionLst {
				if mission.MissionCompleted && mission.Win {
					findWin = true
				}
			}
			if !findWin {
				resp.SetReturn(resultcode.Result_Claimed_Not_Winning_Info)
			} else {
				// 3. market에 nft 전송 요청 / message queue
				data := &basenet.CommandData{
					CommandType: proc.OMZCmd_NFT_Transfer,
					Data: &context.OMZ_NFTTransfer{
						AUID: req.AUID,
					},
					Callback: nil, //콜백은 필요 없다.
				}
				inner.OMZProc(data)
			}
		}
	}

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
