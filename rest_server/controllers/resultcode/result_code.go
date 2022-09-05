package resultcode

const (
	Result_Success            = 0
	ResultInternalServerError = 500

	// db fail
	Result_Error_DB_Normal_fail         = 70000
	Result_Error_DB_OMZ_Get_AirdropInfo = 70001

	// redis error
	Result_PubSub_InternalErr = 71000 // pubsub error

	// auth error
	Result_Auth_InvalidJwt = 71100 // auth error
)

var ResultCodeText = map[int]string{
	Result_Success:            "success",
	ResultInternalServerError: "internal server error",

	Result_Error_DB_Normal_fail:         "DB Error",
	Result_Error_DB_OMZ_Get_AirdropInfo: "Airdrop information inquiry failed",

	Result_PubSub_InternalErr: "redis pubusb error",

	Result_Auth_InvalidJwt: "Invalid jwt token",
}
