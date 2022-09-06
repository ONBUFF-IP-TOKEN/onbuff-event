package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
)

// response GetOMZAirDropInfo
type OMZ_AirDrop struct {
	MissionStartSDT string           `json:"mission_start_sdt"`
	MissionEndSDT   string           `json:"mission_end_sdt"`
	ClaimStartSDT   string           `json:"claim_start_sdt"`
	ClaimEndSDT     string           `json:"claim_end_sdt"`
	IsDrawed        bool             `json:"is_drawed"` // 추첨 여부
	AirDropQuantity int64            `json:"airdrop_quantity"`
	ClaimQuantity   int64            `json:"claim_quantity"`
	Missions        *ResOMZMyMission `json:"missions"`
}

// request GetOMZMyMission
type ReqOMZMyMisssion struct {
	AUID int64 `json:"au_id"`
}

func NewReqMyMission() *ReqOMZMyMisssion {
	return new(ReqOMZMyMisssion)
}

func (o *ReqOMZMyMisssion) CheckValidate(ctx *OnbuffEventContext) *base.BaseResponse {
	if o.AUID == 0 && ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	} else {
		return base.MakeBaseResponse(resultcode.Result_Require_AUID)
	}

	return nil
}

// response GetOMZMyMission
type ResOMZMyMission struct {
	WinningQuantity int64            `json:"winning_quantity"`
	MyMission       []*OMZ_MyMission `json:"my_missions"`
}

// request claim
type ReqOMZClaimAirDrop struct {
	AUID int64 `json:"au_id"`
}

func NewReqOMZClaimAirDrop() *ReqOMZClaimAirDrop {
	return new(ReqOMZClaimAirDrop)
}

func (o *ReqOMZClaimAirDrop) CheckValidate(ctx *OnbuffEventContext) *base.BaseResponse {
	if o.AUID == 0 && ctx.GetValue() != nil {
		o.AUID = ctx.GetValue().AUID
	} else {
		return base.MakeBaseResponse(resultcode.Result_Require_AUID)
	}

	return nil
}

// request PutOMZAirDropMission
type ReqOMZAirDropMission struct {
	AUID      int64 `json:"au_id"`
	MissionID int64 `json:"mission_id"`
}

func NewReqOMZAirDropMission() *ReqOMZAirDropMission {
	return new(ReqOMZAirDropMission)
}

func (o *ReqOMZAirDropMission) CheckValidate(ctx *OnbuffEventContext) *base.BaseResponse {
	return nil
}
