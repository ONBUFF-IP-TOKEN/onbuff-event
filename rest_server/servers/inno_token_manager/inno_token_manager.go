package inno_token_manager

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/token_manager"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

var gTokenServer *token_manager.Server

func GetInstance() *token_manager.Server {
	return gTokenServer
}

func InitTokenManager(conf *config.ServerConfig) error {
	tokenServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.TokenMgrServer.InternalpiDomain,
			ExtHostUri: conf.TokenMgrServer.ExternalpiDomain,
			IntVer:     conf.TokenMgrServer.InternalVer,
			ExtVer:     conf.TokenMgrServer.ExternalVer,
		},
	}

	gTokenServer = token_manager.NewServerInfo(tokenServerInfo)
	return nil
}
