package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
)

var once sync.Once
var currentConfig *ServerConfig

type OnbuffEvent struct {
	ApplicationName string `json:"application_name" yaml:"application_name"`
	APIDocs         bool   `json:"api_docs" yaml:"api_docs"`
	EnableOMZ       bool   `json:"enable_omz"`
}

type ApiAuth struct {
	AuthEnable    bool   `yaml:"auth_enable"`
	ApiAuthDomain string `json:"api_auth_domain" yaml:"api_auth_domain"`
	ApiAuthVerify string `json:"api_auth_verify" yaml:"api_auth_verify"`
}
type Symbol struct {
	Name             string `yaml:"name"`
	ParentWalletAddr string `yaml:"parent_wallet_address"`
	PassPhrase       string `yaml:"pass_phrase"`
	PrivateKey       string `yaml:"private_key"`
	IsBase           bool   `yaml:"is_base"`
}

type Octet struct {
	Name                 string   `yaml:"name"`
	ApiKey               string   `yaml:"api_key"`
	Host                 string   `yaml:"host"`
	Ver                  string   `yaml:"ver"`
	WebHookPort          string   `yaml:"webhook_port"`
	WalletReservedEnable bool     `yaml:"wallet_reserved_enable"`
	ReservedWalletCount  int64    `yaml:"reserved_wallet_count"`
	GenerateWalletCount  int64    `yaml:"generate_wallet_count"`
	Symbols              []Symbol `yaml:"symbols"`
}

type Token struct {
	Octet []Octet `yaml:"octet"`
}
type Wallets struct {
	Name             string `yaml:"name"`
	FeeWalletAddr    string `yaml:"fee_wallet"`
	ParentWalletAddr string `yaml:"parent_wallet"`
}

type ApiInno struct {
	InternalpiDomain string `yaml:"api_internal_domain"`
	ExternalpiDomain string `yaml:"api_external_domain"`
	InternalVer      string `yaml:"internal_ver"`
	ExternalVer      string `yaml:"external_ver"`
}

type Schedule struct {
	Name    string `yaml:"name"`
	TermSec int64  `yaml:"term_sec"`
	Enable  bool   `yaml:"schedule_enable"`
}
type ServerConfig struct {
	baseconf.Config `yaml:",inline"`

	OnbuffEvent OnbuffEvent `yaml:"onbuff_event"`

	MssqlDBAccountAll     baseconf.DBAuth `yaml:"mssql_db_account"`
	MssqlDBAccountRead    baseconf.DBAuth `yaml:"mssql_db_account_read"`
	MssqlDBMarketAll      baseconf.DBAuth `yaml:"mssql_db_market"`
	MssqlDBMarketRead     baseconf.DBAuth `yaml:"mssql_db_market_read"`
	MssqlDBMarketToolAll  baseconf.DBAuth `yaml:"mssql_db_market_tool"`
	MssqlDBMarketToolRead baseconf.DBAuth `yaml:"mssql_db_market_tool_read"`
	MssqlDBEvent          baseconf.DBAuth `yaml:"mssql_db_event"`
	MssqlDBEventRead      baseconf.DBAuth `yaml:"mssql_db_event_read"`

	Auth           ApiAuth `yaml:"api_auth"`
	TokenMgrServer ApiInno `yaml:"api_token_manager_server"`
	InnoLog        ApiInno `yaml:"inno_log"`
	PointMgrServer ApiInno `yaml:"api_point_manager_server"`
	InnoMarket     ApiInno `yaml:"inno_market"`
	WebInno        ApiInno `yaml:"web_inno_server"`

	Schedules   []Schedule `yaml:"schedules"`
	ScheduleMap map[string]Schedule
}

func GetInstance(filepath ...string) *ServerConfig {
	once.Do(func() {
		if len(filepath) <= 0 {
			panic(baseconf.ErrInitConfigFailed)
		}
		currentConfig = &ServerConfig{}
		if err := baseconf.Load(filepath[0], currentConfig); err != nil {
			currentConfig = nil
		} else {

			currentConfig.ScheduleMap = make(map[string]Schedule)
			for _, schedule := range currentConfig.Schedules {
				currentConfig.ScheduleMap[schedule.Name] = schedule
			}
			if os.Getenv("ASPNETCORE_PORT") != "" {
				port, _ := strconv.ParseInt(os.Getenv("ASPNETCORE_PORT"), 10, 32)
				currentConfig.APIServers[0].Port = int(port)
				currentConfig.APIServers[1].Port = int(port)
				fmt.Println(port)
			}
		}
	})

	return currentConfig
}
