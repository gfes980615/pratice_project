package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// 投注紀錄
	mvc.Get("/v1/my/order/{wallet_code}", getOrderByWalletCodeMemEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Member).SetWriter(&biApiWriter)
	mvc.Get("/v1/my/order/{wallet_code}/{channel_code}", getOrderByChannelCodeAndWalletCodeMemEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Member).SetWriter(&biApiWriter)

	mvc.Get("/v1/statistics", getStatisticsEndpoint).SetAuthorizer(&utilsMVC.Anonymous).SetWriter(&biApiWriter)

	// 入款紀錄
	mvc.Get("/v1/recharge/record/", getRechargeRecord).SetWriter(&biApiWriter)
	mvc.Get("/v1/recharge/record/{id}", getRechargeRecordById).SetWriter(&biApiWriter)
	// 出款紀錄
	mvc.Get("/v1/withdrawal/record/", getWithdrawalRecord).SetWriter(&biApiWriter)
	mvc.Get("/v1/withdrawal/record/{id}", getWithdrawalRecordById).SetWriter(&biApiWriter)

	// 交易流水
	mvc.Get("/v1/my/transaction", getTransactionMemEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Member).SetWriter(&biApiWriter)
}

func getOrderByWalletCodeMemEpt(ctx *mvc.Context) {
	wallet_code := ctx.GetRouteValue("wallet_code", "")

	// 因前端 hard code cp 的邏輯，所以加上這段 code 給紅包用。
	// 目前有寫新的 API 不讓前端 hard code，所以這段 code 以後不需要再更新邏輯，
	// 但為了舊版 APP 的相容問題，就保留這段邏輯。

	if wallet_code == "cp" && ctx.GetValue("channelcode", "") == "" {
		ctx.Request.Params["channelcode"] = []string{"fctc,game,hk,js,sf,sport,ssc,tv,wf"}
	} else if wallet_code == "hb" {
		wallet_code = "cp"
		ctx.Request.RouteParams["wallet_code"] = wallet_code
		ctx.Request.Params["channelcode"] = []string{"hb"}
	}

	api := "/v1/athena/my/order/" + wallet_code
	ApisToAthenaV2(ctx, api)
}

func getOrderByChannelCodeAndWalletCodeMemEpt(ctx *mvc.Context) {
	walletCode := ctx.GetRouteValue("wallet_code", "")
	channelCode := ctx.GetRouteValue("channel_code", "")
	api := "/v1/athena/my/order/" + walletCode + "/" + channelCode
	ApisToAthenaV2(ctx, api)
}

func getStatisticsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/statistics"
	ApisToAthenaV2(ctx, api)
}

func getTransactionMemEpt(ctx *mvc.Context) {
	api := "/v1/athena/my/transaction"
	ApisToAthenaV2(ctx, api)
}

func getRechargeRecord(ctx *mvc.Context) {
	api := "/v1/athena/member/recharge/record"
	ApisToAthenaV2(ctx, api)
}

func getRechargeRecordById(ctx *mvc.Context) {
	id := ctx.GetRouteValue("id", "")
	api := "/v1/athena/member/recharge/record/" + id
	ApisToAthenaV2(ctx, api)
}

func getWithdrawalRecord(ctx *mvc.Context) {
	api := "/v1/athena/member/withdrawal/record"
	ApisToAthenaV2(ctx, api)
}

func getWithdrawalRecordById(ctx *mvc.Context) {
	id := ctx.GetRouteValue("id", "")
	api := "/v1/athena/member/withdrawal/record/" + id
	ApisToAthenaV2(ctx, api)
}
