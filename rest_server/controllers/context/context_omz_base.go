package context

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

type OMZ_MyMission struct {
	MissionID        int64  `json:"mission_id"`
	MissionDesc      string `json:"mission_desc"`
	MissionCompleted bool   `json:"mission_completed"` // 미션 완료 여부
	Win              bool   `json:"win"`               // 당첨 여부
}

type OMZ_MyClaimNFT struct {
	NFTID     int64 `json:"nft_id"`
	NFTPackID int64 `json:"nft_pack_id"`
}

// channel used
type OMZ_NFTTransfer struct {
	AUID int64 `json:"au_id"`
}
