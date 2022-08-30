package resultcode

const (
	Result_Success            = 0
	ResultInternalServerError = 500

	Result_Require_Symbol          = 60000
	Result_Require_ToAddress       = 60001
	Result_Require_Amount          = 60002
	Result_Require_Transactioninfo = 60003
	Result_Require_FromAddress     = 60004
	Result_Require_BaseCoinSymbol  = 60005
	Result_Require_WebhookName     = 60006
	Result_Require_WebhookURL      = 60007
	Result_Require_AUID            = 60008
	Result_Require_WebhookID       = 60009
	Result_Require_MUID            = 60010
	Result_Require_NFTID           = 60011

	Result_Error_octet = 61000

	Result_RedisError_Lock_fail = 62000 // redis lock error
	Result_Error_Redis_read     = 62001

	Result_PubSub_InternalErr = 62100 // pubsub error

	//외부서버오류
	Result_PointManager_Network_Err = 62201
	Result_Coin_Load_Err            = 62202
	Result_TokenManager_Network_Err = 62211
	Result_AuthServer_Network_Err   = 62221

	Result_Auth_RequireMessage    = 20000
	Result_Auth_RequireSign       = 20001
	Result_Auth_InvalidLoginInfo  = 20002
	Result_Auth_DontEncryptJwt    = 20003
	Result_Auth_InvalidJwt        = 20004
	Result_Auth_InvalidWalletType = 20005

	//DB Failed
	Result_Error_DB_Normal_fail            = 63000
	Result_Error_NFTPack_Insert_fail       = 63001
	Result_Error_NFTPack_Update_fail       = 63002
	Result_Error_NFTPack_Get_fail          = 64003
	Result_Error_NFTPack_AppID_Insert_fail = 64004

	Result_Error_Product_Insert_fail             = 64011
	Result_Error_Product_Update_fail             = 64012
	Result_Error_ProductStatus_Update_fail       = 64013
	Result_Error_Product_Add_NFTlists_failStatus = 64014
	Result_Error_Product_Merge_priceLists_fail   = 64015
	Result_Error_Product_GetLists_fail           = 64016
	Result_Error_Product_Get_BidDeposit_fail     = 64017
	Result_Error_Coin_Get_fail                   = 64018 // 내 코인 리스트 정보 로드 실패
	Result_Error_Product_Purchase_Start_fail     = 64019 // 상품 구매 시작 에러
	Result_Error_Product_Purchase_Bid_fail       = 64020
	Result_Error_Product_Get_BidList_fail        = 64021
	Result_Error_Mypage_GetList_NFT_fail         = 64022
	Result_Error_Mypage_GetList_Bid_fail         = 64023
	Result_Error_Mypage_GetList_Win_fail         = 64024
	Result_Error_BidDeposit_Purchase_fail        = 64025
	Result_Error_Mypage_GetList_Purchase_fail    = 64026
	Result_Error_Get_Member_Nft_List_fail        = 64027
	Result_Error_Mod_Nft_Customproperties_fail   = 64028
	Result_Error_GetList_NFT_fail                = 64029
	Result_Error_Get_ProductHighestBid           = 64030

	Result_Error_NFT_Cache_Get_fail = 64041

	Result_Error_NFT_Insert_fail   = 64051
	Result_Error_NFT_Scan_fail     = 64052
	Result_Error_NFTPACK_Not_Found = 64053

	Result_Error_App_WebHook_Insert_fail = 64060
	Result_Error_App_WebHook_Get_fail    = 64061
	Result_Error_App_WebHook_Update_fail = 64062
	Result_Error_App_WebHook_Delete_fail = 64063

	Result_Error_Cache_Set_ProductPruchase = 64100 // redis error

	//Params Invalid
	Result_Error_Invalid_MintQuantity  = 64101
	Result_Error_Invalid_NFTPackID     = 64102
	Result_Error_Invalid_TxHash        = 64103
	Result_Error_Invalid_Result        = 64104
	Result_Error_Invalid_PageInfo      = 64105
	Result_Error_Invalid_ProductStatus = 64106
	Result_Error_Invalid_Visible       = 64107
	Result_Error_Invalid_ProductID     = 64108
	Result_Error_Invalid_AppID         = 64109
	Result_Error_Invalid_BaseCoinID    = 64110
	Result_Error_Invalid_CoinID        = 64111
	Result_Error_Invalid_Quantity      = 64112
	Result_Error_Invalid_SectionID     = 64113
	Result_Error_Invalid_BrandID       = 64114
	Result_Error_Invalid_SearchType    = 64115
	Result_Error_Invalid_SearchValue   = 64116

	//입찰,구매처리 오류
	Result_Error_Product_IsNotAuction        = 64201 //경매상품이아니다(상품 상태 이상)
	Result_Error_Product_IsNotPaymentDeposit = 64202 //입찰보증금 납무대상 상품이아니다(상품 상태 이상)
	Result_Error_Product_Refresh             = 64203 //상품정보가 업데이트되었으니 다시시도
	Result_Error_NotEnough_Coin              = 64204 //돈부족
	Result_Error_NotEnough_Gas               = 64205 //가스비 부족
	Result_Error_SoldOut                     = 64206 // 상품 판매 완료 : sold out
	Result_Error_Product_IsNotBid            = 64207 //입찰가능상태가아니다
	Result_Error_Product_Load_Error          = 64208 //상품 로딩 실패
	Result_Error_GasFee_Load_Error           = 64209 //가스비 로딩 실패(코인정보 로딩 실패)
)

var ResultCodeText = map[int]string{
	Result_Success:            "success",
	ResultInternalServerError: "internal server error",

	Result_Require_Symbol:          "Requires symbol information.",
	Result_Require_ToAddress:       "Requires ToAddress information.",
	Result_Require_Amount:          "Requires amount information.",
	Result_Require_Transactioninfo: "Requires transaction information.",
	Result_Require_FromAddress:     "Requires FromAddress information.",
	Result_Require_BaseCoinSymbol:  "Requires basecoinsymbol informatiln",
	Result_Require_WebhookName:     "Requires webhook name information.",
	Result_Require_WebhookURL:      "Requires webhook url information.",
	Result_Require_AUID:            "Requires auid information.",
	Result_Require_WebhookID:       "Requires webhook id information.",
	Result_Require_MUID:            "Requires muid information.",
	Result_Require_NFTID:           "Requires nftid information.",

	Result_Error_octet: "exteral coin net error",

	Result_RedisError_Lock_fail: "redis lock error",
	Result_Error_Redis_read:     "redis read error",

	Result_PubSub_InternalErr: "Internal pubsub error",

	Result_PointManager_Network_Err: "Server Nework Error",
	Result_Coin_Load_Err:            "Coin Load Error",
	Result_TokenManager_Network_Err: "Server Nework Error",
	Result_AuthServer_Network_Err:   "Server Nework Error",

	Result_Auth_RequireMessage:    "Message is required",
	Result_Auth_RequireSign:       "Sign info is required",
	Result_Auth_InvalidLoginInfo:  "Invalid login info",
	Result_Auth_DontEncryptJwt:    "Auth token create fail",
	Result_Auth_InvalidJwt:        "Invalid jwt token",
	Result_Auth_InvalidWalletType: "Invalid wallet type",
	//DB Failed
	Result_Error_DB_Normal_fail:            "DB Error",
	Result_Error_NFTPack_Insert_fail:       "Nft Pack Insert Failed",
	Result_Error_NFTPack_Update_fail:       "Nft Pack Update Failed",
	Result_Error_NFTPack_Get_fail:          "Nft Pack Get Failed",
	Result_Error_NFTPack_AppID_Insert_fail: "Nft Pack AppID Insert Failed",

	Result_Error_Product_Insert_fail:             "Product Insert Failed",
	Result_Error_Product_Update_fail:             "Product Update Failed",
	Result_Error_ProductStatus_Update_fail:       "Product Status Update Failed",
	Result_Error_Product_Add_NFTlists_failStatus: "Product NFTlists Matching Failed",
	Result_Error_Product_Merge_priceLists_fail:   "Product PriceLists Merge Failed",
	Result_Error_Product_GetLists_fail:           "Product Lists Load Failed",
	Result_Error_Product_Get_BidDeposit_fail:     "Product Get Bid Deposit Load Failed",
	Result_Error_Coin_Get_fail:                   "Coin load failed",
	Result_Error_Product_Purchase_Start_fail:     "Product Purchase Failed",
	Result_Error_Product_Purchase_Bid_fail:       "Product Purchase Bid Failed",
	Result_Error_Product_Get_BidList_fail:        "Product Bid List Load Failed",
	Result_Error_Mypage_GetList_NFT_fail:         "MyPage Load Failed",
	Result_Error_Mypage_GetList_Bid_fail:         "MyPage Load Failed",
	Result_Error_Mypage_GetList_Win_fail:         "MyPage Load Failed",
	Result_Error_BidDeposit_Purchase_fail:        "Product Purchase Deposit Failed",
	Result_Error_Mypage_GetList_Purchase_fail:    "MyPage Load Failed",
	Result_Error_Get_Member_Nft_List_fail:        "Failed to load member nft list",
	Result_Error_Mod_Nft_Customproperties_fail:   "Failed to modify nft attribute information",
	Result_Error_GetList_NFT_fail:                "Failed to Get Nft List",
	Result_Error_Get_ProductHighestBid:           "Failed to Get Product HighestBid",

	Result_Error_NFT_Insert_fail:   "Nft Insert Failed",
	Result_Error_NFT_Scan_fail:     "Nft Scan Failed",
	Result_Error_NFTPACK_Not_Found: "NFTINFO Not Found",

	Result_Error_NFT_Cache_Get_fail: "Cache Get TokenID From TxHash Failed",

	Result_Error_App_WebHook_Insert_fail: "WebHook Insert Failed",
	Result_Error_App_WebHook_Get_fail:    "WebHook Get Failed",
	Result_Error_App_WebHook_Update_fail: "WebHook Update Failed",
	Result_Error_App_WebHook_Delete_fail: "WebHook Delete Failed",

	Result_Error_Cache_Set_ProductPruchase: "Purchase caching Failed",

	Result_Error_Invalid_MintQuantity:  "Invalid Parameter 'mint_quantity'",
	Result_Error_Invalid_NFTPackID:     "Invalid Parameter 'pack_id'",
	Result_Error_Invalid_TxHash:        "Invalid Parameter 'txhash'",
	Result_Error_Invalid_Result:        "Invalid Parameter 'result'",
	Result_Error_Invalid_PageInfo:      "Invalid Parameter 'page_offset' or 'page_size'",
	Result_Error_Invalid_ProductStatus: "Invalid Parameter 'product_status' is only 1 or 2 or 7",
	Result_Error_Invalid_Visible:       "Invalid Parameter 'visible'",
	Result_Error_Invalid_ProductID:     "Invalid Parameter 'product_id'",
	Result_Error_Invalid_AppID:         "Invalid Parameter 'app_id'",
	Result_Error_Invalid_BaseCoinID:    "Invalid Parameter `basecoin_id'",
	Result_Error_Invalid_CoinID:        "Invalid Parameter `coin_id`",
	Result_Error_Invalid_Quantity:      "Invalid Parameter `quantity`",
	Result_Error_Invalid_SectionID:     "Invalid Parameter `section_id`",
	Result_Error_Invalid_BrandID:       "Invalid Parameter `brand_id'",
	Result_Error_Invalid_SearchType:    "Invalid Parameter `search_type'",
	Result_Error_Invalid_SearchValue:   "Invalid Parameter `search_value'",

	Result_Error_Product_IsNotAuction:        "This is not an auction item",
	Result_Error_Product_IsNotPaymentDeposit: "This is not payment deposit",
	Result_Error_Product_Refresh:             "Product information has been updated. try again",
	Result_Error_NotEnough_Coin:              "Not Enough Coin Quantity",
	Result_Error_NotEnough_Gas:               "Not Enough Gas Quantity",
	Result_Error_SoldOut:                     "Sold Out",
	Result_Error_Product_IsNotBid:            "Bidding is not available",
	Result_Error_Product_Load_Error:          "Product information load failed",
	Result_Error_GasFee_Load_Error:           "Token information load failed",
}
