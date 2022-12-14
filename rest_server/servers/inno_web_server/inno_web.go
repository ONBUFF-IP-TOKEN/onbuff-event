package inno_web_server

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/web_inno"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

var gServer *web_inno.Server

func GetInstance() *web_inno.Server {
	return gServer
}

func InitWebInno(conf *config.ServerConfig) error {
	ServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.WebInno.InternalpiDomain,
			ExtHostUri: conf.WebInno.ExternalpiDomain,
			IntVer:     conf.WebInno.InternalVer,
			ExtVer:     conf.WebInno.ExternalVer,
		},
	}

	gServer = web_inno.NewServerInfo(ServerInfo)
	return nil
}
