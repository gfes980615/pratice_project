package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}
	// 其他報表相關

	// 下注分析
	mvc.Get("/v1/reseller/reports/betanalysis", getBetAnalysisResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/betanalysisdetail", getBetAnalysisDetailResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	// 會員折扣分析
	mvc.Get("/v1/reseller/reports/memberdiscountanalysis", getMemberDiscountAnalysisResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	// 頻道報表
	mvc.Get("/v1/reseller/reports/order/categories", getChannelReportingByCategoryResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/order/types", getChannelReportingByTypeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/order/members", getChannelReportingByMemberResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	// 輸贏監控
	mvc.Get("/v1/reseller/reports/winlost", getWinlostReportResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	// 會員收入
	mvc.Get("/v1/reseller/reports/revenue", getRevenueReportingResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)

	// 現金流水
	mvc.Get("/reseller/transaction", getTransactionsEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
}

func getBetAnalysisResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/betanalysis"
	ApisToAthenaV2(ctx, api)
}

func getBetAnalysisDetailResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/betanalysisdetail"
	ApisToAthenaV2(ctx, api)
}

func getMemberDiscountAnalysisResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/memberdiscountanalysis"
	ApisToAthenaV2(ctx, api)
}

func getChannelReportingByCategoryResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/order/categories"
	ApisToAthenaV2(ctx, api)
}

func getChannelReportingByTypeEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/order/types"
	ApisToAthenaV2(ctx, api)
}

func getChannelReportingByMemberResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/order/members"
	ApisToAthenaV2(ctx, api)
}

func getWinlostReportResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/winlost"
	ApisToAthenaV2(ctx, api)
}

func getRevenueReportingResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/revenue"
	ApisToAthenaV2(ctx, api)
}

func getTransactionsEpt(ctx *mvc.Context) {
	api := "/athena/reseller/transaction"
	ApisToAthenaV2(ctx, api)
}
