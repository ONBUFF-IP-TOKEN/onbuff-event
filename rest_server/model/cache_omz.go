package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/labstack/gommon/log"
)

const (
	AirDropInfo = "AIRDROPINFO"
	MyMissions  = "MYMISSIONS"
)

func (o *DB) CacheOMZGetAirDropInfo() (*context.OMZ_AirDrop, error) {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return nil, errors.New("redis disable")
	}

	key := MakeOMZKey(AirDropInfo)
	airDropInfo := &context.OMZ_AirDrop{}
	err := o.Cache.Get(key, airDropInfo)

	return airDropInfo, err
}

func (o *DB) CacheOMZSetAirDropInof(params *context.OMZ_AirDrop) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return errors.New("redis disable")
	}
	key := MakeOMZKey(AirDropInfo)
	return o.Cache.Set(key, params, -1)
}

func (o *DB) CacheOMZDelAirDropInof() error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return errors.New("redis disable")
	}
	key := MakeOMZKey(AirDropInfo)
	return o.Cache.Del(key)
}

func (o *DB) CacheOMZSetMyMissions(auid int64, params []*context.OMZ_MyMission) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return errors.New("redis disable")
	}

	key := MakeOMZKey(MyMissions + ":" + strconv.FormatInt(auid, 10))
	return o.Cache.Set(key, params, time.Duration(5*int64(time.Minute)))
}

func (o *DB) CacheOMZGetMyMissions(auid int64) ([]*context.OMZ_MyMission, error) {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return nil, errors.New("redis disable")
	}

	key := MakeOMZKey(MyMissions + ":" + strconv.FormatInt(auid, 10))
	params := []*context.OMZ_MyMission{}
	err := o.Cache.Get(key, params)
	return params, err
}

func (o *DB) CacheOMZDelMyMissions(auid int64) error {
	if !o.Cache.Enable() {
		log.Warnf("redis disable")
		return errors.New("redis disable")
	}
	key := MakeOMZKey(MyMissions + ":" + strconv.FormatInt(auid, 10))
	return o.Cache.Del(key)
}
