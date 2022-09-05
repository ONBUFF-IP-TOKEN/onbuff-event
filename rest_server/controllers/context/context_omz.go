package context

type OMZ_AirDrop struct {
	MissionStartSDT string `json:"mission_start_sdt"`
	MissionEndSDT   string `json:"mission_end_sdt"`
	ClaimStartSDT   string `json:"claim_start_sdt"`
	ClaimEndSDT     string `json:"claim_end_sdt"`
	IsDrawed        bool   `json:"is_drawed"`
	AirDropQuantity int64  `json:"airdrop_quantity"`
	ClaimQuantity   int64  `json:"claim_quantity"`
}
