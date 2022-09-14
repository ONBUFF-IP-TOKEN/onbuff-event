package inno_market_server

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/market"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

var gMarket *market.Server

func GetInstance() *market.Server {
	return gMarket
}

func InitInnoMarket(conf *config.ServerConfig) error {
	ServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.InnoLog.InternalpiDomain,
			ExtHostUri: conf.InnoLog.ExternalpiDomain,
			IntVer:     conf.InnoLog.InternalVer,
			ExtVer:     conf.InnoLog.ExternalVer,
		},
	}

	gMarket = market.NewServerInfo(ServerInfo)
	return nil
}
