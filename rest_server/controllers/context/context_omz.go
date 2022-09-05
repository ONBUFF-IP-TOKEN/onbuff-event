package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
)

// 추첨 여부
const (
	IsDrawed_Incomplete = iota // 추첨 미완료
	IsDrawed_Complete          // 추첨 완료
)

// 미션 완료 여부
const (
	MissionIncomplete = iota // 미션 미완료
	MissionCompleted         // 미션 완료
)

// 당첨 여부
const (
	Loser = iota // 낙첨
	Win          // 당첨
)

// 청구 여부
const (
	ClaimStatus_NotClaim      = iota // 청구 안함
	ClaimStatus_Claiming             // 청구 중
	ClaimStatus_ClaimFail            // 청구 실패
	ClaimStatus_ClaimComplete        // 청구 완료
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
	Missions        []*OMZ_MyMission `json:"missions"`
}

// response GetOMZMyMission
type OMZ_MyMission struct {
	MissionID        int64  `json:"mission_id"`
	MissionDesc      string `json:"mission_desc"`
	MissionCompleted bool   `json:"mission_completed"` // 미션 완료 여부
	Win              bool   `json:"win"`               // 당첨 여부
	ClaimStatus      int64  `json:"claim_status"`      // 청구 여부
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
