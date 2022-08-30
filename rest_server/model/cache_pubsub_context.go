package model

import (
	"fmt"

	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

const (
	PubSub      = "pubsub"
	InternalCmd = "internal_cmd"
)

type PSHeader struct {
	Type string `json:"type"`
}

// pubsub cmd
const (
	PubSub_cmd_healthcheck  = "HealthCheck"
	PubSub_cmd_meta_refresh = "MetaRefresh"
	PubSub_cmd_product      = "Product"
)

func MakePubSubKey(val string) string {
	return fmt.Sprintf("%s:%s:%s", config.GetInstance().DBPrefix, PubSub, val)
}

type PSHealthCheck struct {
	PSHeader
	Value struct {
		Timestamp int64 `json:"ts"`
	} `json:"value"`
}

type PSMetaRefresh struct {
	PSHeader
	Value struct {
		Refresh       bool  `json:"refresh"`
		RefreshTarget int64 `json:"refresh_target"`
	} `json:"value"`
}
