package model

import (
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/basedb"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/config"

	baseconf "github.com/ONBUFF-IP-TOKEN/baseapp/config"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var gDB *DB

type DB struct {
	MssqlAccountAll     *basedb.Mssql
	MssqlAccountRead    *basedb.Mssql
	MssqlMarketAll      *basedb.Mssql
	MssqlMarketRead     *basedb.Mssql
	MssqlMarketToolAll  *basedb.Mssql
	MssqlMarketToolRead *basedb.Mssql

	Cache *basedb.CacheV8

	RedSync *redsync.Redsync
}

func GetDB() *DB {
	return gDB
}

func InitDB(conf *config.ServerConfig) (err error) {
	cache := basedb.GetCacheV8(&conf.Cache)
	gDB = &DB{
		Cache: cache,
	}
	pool := goredis.NewPool(cache.GetDB().RedisClient())
	gDB.RedSync = redsync.New(pool)

	if err := ConnectAllDB(conf); err != nil {
		log.Errorf("InitDB Error : %v", err)
		return err
	}

	go func() {
		for {
			timer := time.NewTimer(5 * time.Second)
			<-timer.C
			timer.Stop()

			// DB ping 체크 후 오류 시 재 연결
			if db := CheckPingDB(gDB.MssqlAccountAll, &conf.MssqlDBAccountAll); db != nil {
				gDB.MssqlAccountAll = db
			}
			if db := CheckPingDB(gDB.MssqlAccountRead, &conf.MssqlDBAccountRead); db != nil {
				gDB.MssqlAccountRead = db
			}
			if db := CheckPingDB(gDB.MssqlMarketAll, &conf.MssqlDBMarketAll); db != nil {
				gDB.MssqlMarketAll = db
			}
			if db := CheckPingDB(gDB.MssqlMarketRead, &conf.MssqlDBMarketRead); db != nil {
				gDB.MssqlMarketRead = db
			}
			if db := CheckPingDB(gDB.MssqlMarketToolAll, &conf.MssqlDBMarketToolAll); db != nil {
				gDB.MssqlMarketAll = db
			}
			if db := CheckPingDB(gDB.MssqlMarketToolRead, &conf.MssqlDBMarketToolRead); db != nil {
				gDB.MssqlMarketRead = db
			}
		}
	}()

	LoadDBMeta()

	schedule := conf.ScheduleMap["redis_pubsub_keepalive"]
	go gDB.ListenSubscribeEvent(schedule.Enable, schedule.TermSec)

	return nil
}
func ConnectAllDB(conf *config.ServerConfig) error {
	var err error
	gDB.MssqlAccountAll, err = gDB.ConnectDB(&conf.MssqlDBAccountAll)
	if err != nil {
		return err
	}

	gDB.MssqlAccountRead, err = gDB.ConnectDB(&conf.MssqlDBAccountRead)
	if err != nil {
		return err
	}

	gDB.MssqlMarketAll, err = gDB.ConnectDB(&conf.MssqlDBMarketAll)
	if err != nil {
		return err
	}

	gDB.MssqlMarketRead, err = gDB.ConnectDB(&conf.MssqlDBMarketRead)
	if err != nil {
		return err
	}

	gDB.MssqlMarketToolAll, err = gDB.ConnectDB(&conf.MssqlDBMarketToolAll)
	if err != nil {
		return err
	}

	gDB.MssqlMarketToolRead, err = gDB.ConnectDB(&conf.MssqlDBMarketToolRead)
	if err != nil {
		return err
	}
	return nil
}
func (o *DB) ConnectDB(conf *baseconf.DBAuth) (*basedb.Mssql, error) {
	port, _ := strconv.ParseInt(conf.Port, 10, 32)
	mssqlDB, err := basedb.NewMssql(conf.Database, "", conf.ID, conf.Password, conf.Host, int(port),
		conf.ApplicationIntent, conf.Timeout, conf.ConnectRetryCount, conf.ConnectRetryInterval)
	if err != nil {
		log.Errorf("err: %v, val: %v, %v, %v, %v, %v, %v",
			err, conf.Host, conf.ID, conf.Password, conf.Database, conf.PoolSize, conf.IdleSize)
		return nil, err
	}

	idleSize, _ := strconv.ParseInt(conf.IdleSize, 10, 32)
	mssqlDB.GetDB().SetMaxOpenConns(int(idleSize))
	mssqlDB.GetDB().SetMaxIdleConns(int(idleSize))

	return mssqlDB, nil
}

func CheckPingDB(db *basedb.Mssql, conf *baseconf.DBAuth) *basedb.Mssql {
	// 연결이 안되어있거나, DB Connection이 끊어진 경우에는 재연결 시도
	if db == nil || !db.Connection.IsConnect {
		var err error
		newDB, err := gDB.ConnectDB(conf)
		if err == nil {
			log.Errorf("connect DB OK")
		}
		return newDB
	}

	// 연결이 되어있는 상태면 ping
	if db.Connection.IsConnect {
		if err := db.GetDB().Ping(); err != nil {
			// 재시도 횟수
			db.Connection.RetryCount += 1
			log.Errorf("%v DB Ping err RetryCount(%v)", conf.Database, db.Connection.RetryCount)
			// ping 2회 시도해도 안되면 close
			if db.Connection.RetryCount >= 2 {
				db.Connection.IsConnect = false
				// DB Close
				if err = db.GetDB().Close(); err == nil {
					log.Errorf("DB Closed (RetryCount >=2)")
				}
			}
		}
	}
	return nil
}

func LoadDBMeta() {
}
