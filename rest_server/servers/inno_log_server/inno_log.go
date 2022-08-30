package inno_log_server

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/inno_log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

var gLogServer *inno_log.Server

func GetInstance() *inno_log.Server {
	return gLogServer
}

func InitInnoLog(conf *config.ServerConfig) error {
	ServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.InnoLog.InternalpiDomain,
			ExtHostUri: conf.InnoLog.ExternalpiDomain,
			IntVer:     conf.InnoLog.InternalVer,
			ExtVer:     conf.InnoLog.ExternalVer,
		},
	}

	gLogServer = inno_log.NewServerInfo(ServerInfo)
	return nil
}
