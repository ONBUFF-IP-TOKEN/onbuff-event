package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-event/rest_server/controllers/auth"
)

// InnoLogContext API의 Request Context
type OnbuffEventContext struct {
	*base.BaseContext
	VerifyValue *auth.VerifyAuthToken
}

// common
type PageInfo struct {
	PageOffset int64 `json:"page_offset" query:"page_offset"`
	PageSize   int64 `json:"page_size" query:"page_size"`
}

func NewOnbuffEventContext(baseCtx *base.BaseContext) interface{} {
	if baseCtx == nil {
		return nil
	}

	ctx := new(OnbuffEventContext)
	ctx.BaseContext = baseCtx
	ctx.VerifyValue = new(auth.VerifyAuthToken)

	return ctx
}

// AppendRequestParameter BaseContext 이미 정의되어 있는 ReqeustParameters 배열에 등록
func AppendRequestParameter() {
}

func (o *OnbuffEventContext) SetVerifyAuthToken(value *auth.VerifyAuthToken) {
	o.VerifyValue = value
}

func (o *OnbuffEventContext) GetValue() *auth.VerifyAuthToken {
	return o.VerifyValue
}
