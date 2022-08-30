package inno_point_manager

import (
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/context"
	"github.com/ONBUFF-IP-TOKEN/baseInnoClient/point_manager"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
)

var gPointServer *point_manager.Server

func GetInstance() *point_manager.Server {
	return gPointServer
}

func InitPointManager(conf *config.ServerConfig) error {
	ServerInfo := &context.ServerInfo{
		HostInfo: context.HostInfo{
			IntHostUri: conf.PointMgrServer.InternalpiDomain,
			ExtHostUri: conf.PointMgrServer.ExternalpiDomain,
			IntVer:     conf.PointMgrServer.InternalVer,
			ExtVer:     conf.PointMgrServer.ExternalVer,
		},
	}

	gPointServer = point_manager.NewServerInfo(ServerInfo)
	return nil
}
