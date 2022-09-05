package model

import (
	originCtx "context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/context"
	orginMssql "github.com/denisenkom/go-mssqldb"
)

const (
	USPE_Get_AirDrop                    = "[dbo].[USPE_Get_AirDrop]"
	USPE_GetList_AccountAirDropMissions = "[dbo].[USPE_GetList_AccountAirDropMissions]"
)

func (o *DB) USPE_Get_AirDrop() (*context.OMZ_AirDrop, error) {
	ProcName := USPE_Get_AirDrop
	var rs orginMssql.ReturnStatus
	airDrop := &context.OMZ_AirDrop{}
	rows, err := o.MssqlEvent.GetDB().QueryContext(originCtx.Background(), ProcName,
		sql.Named("MissionStartSDT", sql.Out{Dest: &airDrop.MissionStartSDT}),
		sql.Named("MissionEndSDT", sql.Out{Dest: &airDrop.MissionEndSDT}),
		sql.Named("ClaimStartSDT", sql.Out{Dest: &airDrop.ClaimStartSDT}),
		sql.Named("ClaimEndSDT", sql.Out{Dest: &airDrop.ClaimEndSDT}),
		sql.Named("IsDrawed", sql.Out{Dest: &airDrop.IsDrawed}),
		sql.Named("AirDropQuantity", sql.Out{Dest: &airDrop.AirDropQuantity}),
		sql.Named("ClaimQuantity", sql.Out{Dest: &airDrop.ClaimQuantity}),
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return nil, err
	}

	defer rows.Close()

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return nil, errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	if t, err := time.Parse("Jan 2 2006 3:4PM", airDrop.MissionStartSDT); err == nil {
		airDrop.MissionStartSDT = t.Format("2006-01-02T15:04:05Z")
	}
	if t, err := time.Parse("Jan 2 2006 3:4PM", airDrop.MissionEndSDT); err == nil {
		airDrop.MissionEndSDT = t.Format("2006-01-02T15:04:05Z")
	}
	if t, err := time.Parse("Jan 2 2006 3:4PM", airDrop.ClaimStartSDT); err == nil {
		airDrop.ClaimStartSDT = t.Format("2006-01-02T15:04:05Z")
	}
	if t, err := time.Parse("Jan 2 2006 3:4PM", airDrop.ClaimEndSDT); err == nil {
		airDrop.ClaimEndSDT = t.Format("2006-01-02T15:04:05Z")
	}

	return airDrop, nil
}
