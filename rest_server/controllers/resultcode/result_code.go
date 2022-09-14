package resultcode

const (
	Result_Success            = 0
	ResultInternalServerError = 500

	// db fail
	Result_Error_DB_Normal_fail                    = 70000
	Result_Error_DB_OMZ_Get_AirdropInfo            = 70001
	Result_Error_DB_OMZ_Get_Mission                = 70002
	Result_Error_DB_OMZ_Add_AccountAirDropMissions = 70003

	// redis error
	Result_PubSub_InternalErr = 71000 // pubsub error

	// auth error
	Result_Auth_InvalidJwt = 71100 // auth error

	// param error
	Result_Require_AUID = 71200 // require auid

	// logic error
	Result_Claimed_Already          = 71300 // already claimed
	Result_Claimed_Not_Winning_Info = 71301 // No winning information
	Result_Claimed_Succes           = 71302 // 	claim success
)

var ResultCodeText = map[int]string{
	Result_Success:            "success",
	ResultInternalServerError: "internal server error",

	Result_Error_DB_Normal_fail:                    "DB Error",
	Result_Error_DB_OMZ_Get_AirdropInfo:            "Airdrop information inquiry failed",
	Result_Error_DB_OMZ_Get_Mission:                "Mission lookup failed",
	Result_Error_DB_OMZ_Add_AccountAirDropMissions: "Failed to add mission",

	Result_PubSub_InternalErr: "redis pubusb error",

	Result_Auth_InvalidJwt: "Invalid jwt token",

	Result_Require_AUID: "Requires auid information.",

	Result_Claimed_Already:          "already claimed",
	Result_Claimed_Not_Winning_Info: "No winning information",
	Result_Claimed_Succes:           "claim success",
}
