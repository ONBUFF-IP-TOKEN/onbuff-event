package model

import "github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"

func MakeOMZKey(key string) string {
	return config.GetInstance().DBPrefix + ":OMZ:" + key
}
