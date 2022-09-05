package commonapi

import (
	"net/http"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/model"
	"github.com/labstack/echo"
)

func GetOMZAirDropInfo(c echo.Context) error {
	resp := new(base.BaseResponse)
	resp.Success()

	// redis 확인 후 없으면 db load
	if airDropInfo, err := model.GetDB().CacheOMZGetAirDropInfo(); err != nil {
		if airDropInfo, err := model.GetDB().USPE_Get_AirDrop(); err != nil {
			resp.SetReturn(resultcode.Result_Error_DB_OMZ_Get_AirdropInfo)
		} else {
			resp.Value = airDropInfo
			model.GetDB().CacheOMZSetAirDropInof(airDropInfo)
		}
	} else {
		resp.Value = airDropInfo
	}

	return c.JSON(http.StatusOK, resp)
}
