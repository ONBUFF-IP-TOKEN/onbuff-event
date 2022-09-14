package app

import (
	"fmt"
	"sync"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/auth"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/externalapi"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/internalapi"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/proc"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_log_server"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_market_server"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_point_manager"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_token_manager"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/servers/inno_web_server"
)

type ServerApp struct {
	base.BaseApp
	conf *config.ServerConfig

	procMgr *proc.ProcManager
}

func (o *ServerApp) Init(configFile string) (err error) {
	o.conf = config.GetInstance(configFile)
	base.AppendReturnCodeText(&resultcode.ResultCodeText)
	context.AppendRequestParameter()

	auth.InitHttpClient()
	o.InitTokenManagerServer(o.conf)
	o.InitLogServer(o.conf)
	o.InitPointManagerServer(o.conf)
	o.InitMarketServer(o.conf)
	o.InitWebInnoServer(o.conf)

	if err := o.NewDB(o.conf); err != nil {
		return err
	}

	if err := o.InitProcManager(); err != nil {
		return err
	}

	return err
}

func (o *ServerApp) CleanUp() {
	fmt.Println("CleanUp")
}

func (o *ServerApp) Run(wg *sync.WaitGroup) error {
	return nil
}

func (o *ServerApp) GetConfig() *baseconf.Config {
	return &o.conf.Config
}

func NewApp() (*ServerApp, error) {
	app := &ServerApp{}

	intAPI := internalapi.NewAPI()
	extAPI := externalapi.NewAPI()

	if err := app.BaseApp.Init(app, intAPI, extAPI); err != nil {
		return nil, err
	}

	return app, nil
}

func (o *ServerApp) NewDB(conf *config.ServerConfig) error {
	return model.InitDB(conf)
}

func (o *ServerApp) InitTokenManagerServer(conf *config.ServerConfig) error {
	return inno_token_manager.InitTokenManager(conf)
}

func (o *ServerApp) InitLogServer(conf *config.ServerConfig) error {
	return inno_log_server.InitInnoLog(conf)
}

func (o *ServerApp) InitPointManagerServer(conf *config.ServerConfig) error {
	return inno_point_manager.InitPointManager(conf)
}

func (o *ServerApp) InitMarketServer(conf *config.ServerConfig) error {
	return inno_market_server.InitInnoMarket(conf)
}

func (o *ServerApp) InitWebInnoServer(conf *config.ServerConfig) error {
	return inno_web_server.InitWebInno(conf)
}

func (o *ServerApp) InitProcManager() error {
	o.procMgr = proc.NewProcManager()

	if err := o.procMgr.Init(); err != nil {
		return err
	}
	return nil
}
