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
	USPE_ClmStrt_AirDrop                = "[dbo].[USPE_ClmStrt_AirDrop]"
	USPE_ClmCmplt_AirDrop               = "[dbo].[USPE_ClmCmplt_AirDrop]"
	USPE_Add_AccountAirDropMissions     = "[dbo].[USPE_Add_AccountAirDropMissions]"
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

func (o *DB) USPE_GetList_AccountAirDropMissions(auid int64) ([]*context.OMZ_MyMission, bool, error) {
	ProcName := USPE_GetList_AccountAirDropMissions
	isClaimed := false
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlEventRead.GetDB().QueryContext(originCtx.Background(), ProcName,
		sql.Named("AUID", auid),
		sql.Named("IsClaimed", sql.Out{Dest: &isClaimed}),
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return nil, isClaimed, err
	}

	defer rows.Close()

	myMissions := []*context.OMZ_MyMission{}
	for rows.Next() {
		myMission := &context.OMZ_MyMission{}
		if err := rows.Scan(&myMission.MissionID,
			&myMission.MissionDesc,
			&myMission.MissionCompleted,
			&myMission.Win); err == nil {
			myMissions = append(myMissions, myMission)
		} else {
			log.Errorf(ProcName+" Scan error : %v", err)
		}
	}

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return nil, isClaimed, errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return myMissions, isClaimed, nil
}

func (o *DB) USPE_ClmStrt_AirDrop(auid int64) ([]*context.OMZ_MyClaimNFT, int64, error) {
	ProcName := USPE_ClmStrt_AirDrop
	claimQuantity := int64(0)
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlEvent.GetDB().QueryContext(originCtx.Background(), ProcName,
		sql.Named("AUID", auid),
		sql.Named("ClaimQuantity", sql.Out{Dest: &claimQuantity}),
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return nil, 0, err
	}

	defer rows.Close()

	myClaims := []*context.OMZ_MyClaimNFT{}
	for rows.Next() {
		myClaim := &context.OMZ_MyClaimNFT{}
		if err := rows.Scan(&myClaim.NFTID,
			&myClaim.NFTPackID); err == nil {
			myClaims = append(myClaims, myClaim)
		} else {
			log.Errorf(ProcName+" Scan error : %v", err)
		}
	}

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return nil, 0, errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return myClaims, claimQuantity, nil
}

func (o *DB) USPE_ClmCmplt_AirDrop(auid int64, transferID int64, nftID int64, txHash string) error {
	ProcName := USPE_ClmCmplt_AirDrop
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlEvent.GetDB().QueryContext(originCtx.Background(), ProcName,
		sql.Named("AUID", auid),
		sql.Named("TransferID", transferID),
		sql.Named("NFTID", nftID),
		sql.Named("TxHash", txHash),
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return err
	}

	defer rows.Close()

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return nil
}

func (o *DB) USPE_Add_AccountAirDropMissions(auid int64, missionID int64) error {
	ProcName := USPE_Add_AccountAirDropMissions
	var rs orginMssql.ReturnStatus
	rows, err := o.MssqlEvent.GetDB().QueryContext(originCtx.Background(), ProcName,
		sql.Named("AUID", auid),
		sql.Named("MissionID", missionID),
		&rs)
	if err != nil {
		log.Errorf(ProcName+" QueryContext err : %v", err)
		return err
	}

	defer rows.Close()

	if rs != 1 {
		log.Errorf(ProcName+" returnvalue error : %v", rs)
		return errors.New(ProcName + " returnvalue error " + strconv.Itoa(int(rs)))
	}

	return nil
}
