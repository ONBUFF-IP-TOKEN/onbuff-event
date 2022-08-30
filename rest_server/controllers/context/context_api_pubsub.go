package context

// pubsub api healthcheck
type ReqPSHealthCheck struct {
	Timestamp int64 `json:"ts"`
}

func NewReqPSHealthCheck() *ReqPSHealthCheck {
	return new(ReqPSHealthCheck)
}

type ResPSHealthCheck struct {
	Timestamp int64 `json:"ts"`
}

/////////////////////////////////////////////////////////

// pubsub api meta db refresh
type ReqPSMetaRefresh struct {
	Refresh       bool  `json:"refresh"`
	RefreshTarget int64 `json:"refresh_target"`
}

func NewReqPSMetaRefresh() *ReqPSMetaRefresh {
	return new(ReqPSMetaRefresh)
}

/////////////////////////////////////////////////////////

// pubsub prouduct info change alarm
type ReqPSProduct struct {
	Type      int64 `json:"type"`
	ProductID int64 `json:"product_id"`
}

func NewReqPSProduct() *ReqPSProduct {
	return new(ReqPSProduct)
}
