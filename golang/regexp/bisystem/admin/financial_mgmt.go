package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}
	// 現金系統

	// 現金系統-現金流水
	mvc.Get("/admin/transaction", getTransactions).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/transaction/daily", getTransactionsDaily).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 現金系統-財務報表
	mvc.Get("/v1/reports/financial/overview", getReportOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 公司入款
	mvc.Get("/v1/reports/financial/deposititems", getDepositItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/deposititems/{account}", getDepositItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 線上支付手續費
	mvc.Get("/v1/reports/financial/webdepositfeeitems", getWebDepositFeeItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/webdepositfeeitems/{merchant_id}", getWebDepositFeeItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 人工存入
	mvc.Get("/v1/reports/financial/manualdeposititems", getManualDepositItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/manualdeposititems/{action_code}", getManualDepositItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 人工提出
	mvc.Get("/v1/reports/financial/manualwithdrawitems", getManualWithDrawItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/manualwithdrawitems/{action_code}", getManualWithDrawItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 線上支付
	mvc.Get("/v1/reports/financial/webdeposititems", getWebDepositItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/webdeposititems/{merchant_id}", getWebDepositItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 出款申請
	mvc.Get("/v1/reports/financial/withdrawitems", getWithDrawItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/withdrawitems/{group}", getWithDrawItemsByGroupEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/withdrawitems/{group}/{account}", getWithDrawItemsAccountDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 會員出款被扣除金額
	mvc.Get("/v1/reports/financial/withdrawfeeitems", getWithDrawFeeItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/withdrawfeeitems/{charge_type}", getWithDrawFeeItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 給予優惠
	mvc.Get("/v1/reports/financial/discountitems", getDiscountItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/discountitems/{deposit_type}", getDiscountItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/discountitems/{deposit_type}/{action_code}", getDiscountItemsDetailsByActionCodeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 給予反水
	mvc.Get("/v1/reports/financial/rebateitems", getRebateItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/rebateitems/{subtype_code}", getRebateItemsDetailsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/rebate/overview", getRebateOverviewEndPoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/rebate/realtime", getRebateRealtimeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/rebate/realtime/items", getRebateRealtimeItemsEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/financial/rebate/ladder", getRebateLadderItmesEndPoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 現金系統-差額報表
	mvc.Get("/admin/reports/diffreporting/overview", getDiffReportingOverviewEndPoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/diffreporting/records", getDiffReportingEndPoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 入款管理
	// 入款管理-入款總覽
	mvc.Get("/v1/admin/reports/depositoverview", getDepositOverviewAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 入款管理-公司入款
	mvc.Get("/admin/deposit", getDeposit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/deposit/{id}/image/", getDepositImageByID).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 入款管理-在線入款
	mvc.Get("/admin/webdeposit", webDepositItems).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 入款管理-人工存入
	mvc.Get("/admin/mandeposit", getMandeposit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/mandeposit/actioncode", getMandepositActionCode).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/mandepositaudit", getMandepositAudit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 入款管理-自動入帳
	mvc.Get("/athena/admin/auto_accept", getAutoDeposit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 銀行流水
	mvc.Get("/athena/admin/bank_list", getBankList).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/athena/admin/deposit/{deposit_id}/bank_list", getBankListByID).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 出款管理
	// 出款總覽
	mvc.Get("/v1/admin/reports/withdrawoverview", getWithdrawOverviewAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 出款申請
	mvc.Get("/admin/withdraw", getWithdraw).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 在線出款
	mvc.Get("/admin/webwithdraw", getWebWithdrawEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 人工提出
	mvc.Get("/admin/manwithdraw", getManWithdraw).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 外接平台
	// cmtxn
	mvc.Get("/admin/cmtransaction", cmTransactionEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/cmtransaction/latest", getLatestCMTransactionEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// cmwallet
	mvc.Get("/admin/cmwallet", GetCMWalletAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/cmwallet/latest", SyncCMWalletAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/cmwallet/check", checkThresholdAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/cmwallet/{walletid}/check", checkThresholdByIDAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

}

func getTransactions(ctx *mvc.Context) {
	api := "/athena/admin/transaction"
	ApisToAthenaV2(ctx, api)
}

func getTransactionsDaily(ctx *mvc.Context) {
	api := "/athena/admin/transaction/daily"
	ApisToAthenaV2(ctx, api)
}

func getReportOverview(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/overview"
	ApisToAthenaV2(ctx, api)
}

func getDepositItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/deposititems"
	ApisToAthenaV2(ctx, api)
}

func getDepositItemsDetailsEndpoint(ctx *mvc.Context) {
	account := ctx.GetRouteValue("account", "")
	api := "/v1/athena/reports/financial/deposititems/" + account
	ApisToAthenaV2(ctx, api)
}

func getWebDepositFeeItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/webdepositfeeitems"
	ApisToAthenaV2(ctx, api)
}

func getWebDepositFeeItemsDetailsEndpoint(ctx *mvc.Context) {
	merchant_id := ctx.GetRouteValue("merchant_id", "")
	api := "/v1/athena/reports/financial/webdepositfeeitems/" + merchant_id
	ApisToAthenaV2(ctx, api)
}

func getManualDepositItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/manualdeposititems"
	ApisToAthenaV2(ctx, api)
}

func getManualDepositItemsDetailsEndpoint(ctx *mvc.Context) {
	action_code := ctx.GetRouteValue("action_code", "")
	api := "/v1/athena/reports/financial/manualdeposititems/" + action_code
	ApisToAthenaV2(ctx, api)
}

func getManualWithDrawItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/manualwithdrawitems"
	ApisToAthenaV2(ctx, api)
}

func getManualWithDrawItemsDetailsEndpoint(ctx *mvc.Context) {
	action_code := ctx.GetRouteValue("action_code", "")
	api := "/v1/athena/reports/financial/manualwithdrawitems/" + action_code
	ApisToAthenaV2(ctx, api)
}

func getWebDepositItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/webdeposititems"
	ApisToAthenaV2(ctx, api)
}

func getWebDepositItemsDetailsEndpoint(ctx *mvc.Context) {
	merchant_id := ctx.GetRouteValue("merchant_id", "")
	api := "/v1/athena/reports/financial/webdeposititems/" + merchant_id
	ApisToAthenaV2(ctx, api)
}

func getWithDrawItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/withdrawitems"
	ApisToAthenaV2(ctx, api)
}

func getWithDrawItemsByGroupEndpoint(ctx *mvc.Context) {
	group := ctx.GetRouteValue("group", "")
	api := "/v1/athena/reports/financial/withdrawitems/" + group
	ApisToAthenaV2(ctx, api)
}

func getWithDrawItemsAccountDetailsEndpoint(ctx *mvc.Context) {
	group := ctx.GetRouteValue("group", "")
	account := ctx.GetRouteValue("account", "")
	api := "/v1/athena/reports/financial/withdrawitems/" + group + "/" + account
	ApisToAthenaV2(ctx, api)
}

func getWithDrawFeeItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/withdrawfeeitems"
	ApisToAthenaV2(ctx, api)
}

func getWithDrawFeeItemsDetailsEndpoint(ctx *mvc.Context) {
	chargeType := ctx.GetRouteValue("charge_type", "")
	api := "/v1/athena/reports/financial/withdrawfeeitems/" + chargeType
	ApisToAthenaV2(ctx, api)
}

func getDiscountItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/discountitems"
	ApisToAthenaV2(ctx, api)
}

func getDiscountItemsDetailsEndpoint(ctx *mvc.Context) {
	depositType := ctx.GetRouteValue("deposit_type", "")
	api := "/v1/athena/reports/financial/discountitems/" + depositType
	ApisToAthenaV2(ctx, api)
}

func getDiscountItemsDetailsByActionCodeEndpoint(ctx *mvc.Context) {
	depositType := ctx.GetRouteValue("deposit_type", "")
	actionCode := ctx.GetRouteValue("action_code", "")
	api := "/v1/athena/reports/financial/discountitems/" + depositType + "/" + actionCode
	ApisToAthenaV2(ctx, api)
}

func getRebateItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/rebateitems"
	ApisToAthenaV2(ctx, api)
}

func getRebateItemsDetailsEndpoint(ctx *mvc.Context) {
	subtype_code := ctx.GetRouteValue("subtype_code", "")
	api := "/v1/athena/reports/financial/rebateitems/" + subtype_code
	ApisToAthenaV2(ctx, api)
}

func getRebateOverviewEndPoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/rebate/overview"
	ApisToAthenaV2(ctx, api)
}

func getRebateRealtimeEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/rebate/realtime"
	ApisToAthenaV2(ctx, api)
}

func getRebateRealtimeItemsEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/rebate/realtime/items"
	ApisToAthenaV2(ctx, api)
}

func getRebateLadderItmesEndPoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/financial/rebate/ladder"
	ApisToAthenaV2(ctx, api)
}

func getDiffReportingOverviewEndPoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/diffreporting/overview"
	ApisToAthenaV2(ctx, api)
}

func getDiffReportingEndPoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/diffreporting/records"
	ApisToAthenaV2(ctx, api)
}

func getDepositOverviewAdmEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/reports/depositoverview"
	ApisToAthenaV2(ctx, api)
}

func getDeposit(ctx *mvc.Context) {
	api := "/athena/admin/deposit"
	ApisToAthenaV2(ctx, api)
}

func getDepositImageByID(ctx *mvc.Context) {
	id := ctx.GetRouteValue("id", "")
	api := "/athena/admin/deposit/" + id + "/image/"
	ApisToAthenaV2(ctx, api)
}

func webDepositItems(ctx *mvc.Context) {
	api := "/athena/admin/webdeposit"
	ApisToAthenaV2(ctx, api)
}

func getMandeposit(ctx *mvc.Context) {
	api := "/athena/admin/mandeposit"
	ApisToAthenaV2(ctx, api)
}

func getMandepositActionCode(ctx *mvc.Context) {
	api := "/athena/admin/mandeposit/actioncode"
	ApisToAthenaV2(ctx, api)
}

func getMandepositAudit(ctx *mvc.Context) {
	api := "/athena/admin/mandepositaudit"
	ApisToAthenaV2(ctx, api)
}

func getAutoDeposit(ctx *mvc.Context) {
	api := "/athena/admin/auto_accept"
	ApisToAthenaV2(ctx, api)
}

func getBankList(ctx *mvc.Context) {
	api := "/athena/admin/bank_list"
	ApisToAthenaV2(ctx, api)
}

func getBankListByID(ctx *mvc.Context) {
	depositID := ctx.GetRouteValue("deposit_id", "")
	api := "/athena/admin/deposit/" + depositID + "/bank_list"
	ApisToAthenaV2(ctx, api)
}

func getWithdrawOverviewAdmEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/reports/withdrawoverview"
	ApisToAthenaV2(ctx, api)
}

func getWithdraw(ctx *mvc.Context) {
	api := "/athena/admin/withdraw"
	ApisToAthenaV2(ctx, api)
}

func getWebWithdrawEpt(ctx *mvc.Context) {
	api := "/athena/admin/webwithdraw"
	ApisToAthenaV2(ctx, api)
}

func getManWithdraw(ctx *mvc.Context) {
	api := "/athena/admin/manwithdraw"
	ApisToAthenaV2(ctx, api)
}
func cmTransactionEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/cmtransaction"
	ApisToAthenaV2(ctx, api)
}

func getLatestCMTransactionEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/cmtransaction/latest"
	ApisToAthenaV2(ctx, api)
}

func auditCMTransactionEndpoint(ctx *mvc.Context) {
	transactionid := ctx.GetRouteValue("transactionid", "")
	api := "/athena/admin/cmtransaction/" + transactionid + "/audit"
	ApisToAthenaV2(ctx, api)
}

func GetCMWalletAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/cmwallet"
	ApisToAthenaV2(ctx, api)
}

func SyncCMWalletAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/cmwallet/latest"
	ApisToAthenaV2(ctx, api)
}

func checkThresholdByIDAdmEpt(ctx *mvc.Context) {
	walletIDStr := ctx.GetRouteValue("walletid", "")
	api := "/athena/admin/cmwallet/" + walletIDStr + "/check"
	ApisToAthenaV2(ctx, api)
}

func checkThresholdAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/cmwallet/check"
	ApisToAthenaV2(ctx, api)
}
