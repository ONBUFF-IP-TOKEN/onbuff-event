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
			IntHostUri: conf.InnoMarket.InternalpiDomain,
			ExtHostUri: conf.InnoMarket.ExternalpiDomain,
			IntVer:     conf.InnoMarket.InternalVer,
			ExtVer:     conf.InnoMarket.ExternalVer,
		},
	}

	gMarket = market.NewServerInfo(ServerInfo)
	return nil
}
