package model

import (
	"errors"

	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	"github.com/labstack/gommon/log"
)

const (
	AirDropInfo = "AIRDROPINFO"
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
	}
	key := MakeOMZKey(AirDropInfo)
	return o.Cache.Set(key, params, -1)
}
