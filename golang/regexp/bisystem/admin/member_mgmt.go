package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// 會員管理

	// 會員管理-會員分析
	//下注分析
	mvc.Get("/admin/reports/betanalysis", getBetAnalysisAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/betanalysisdetail", getBetAnalysisDetailAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 出入款統計
	mvc.Get("/admin/reporting/revenue", getRevenueReportingAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 優惠統計
	mvc.Get("/admin/reports/memberdiscountanalysis", getMemberDiscountAnalysis).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 會員管理-輸贏監控
	mvc.Get("/v1/reports/winlost", getWinlostReportEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
}

func getBetAnalysisAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/betanalysis"
	ApisToAthenaV2(ctx, api)
}

func getBetAnalysisDetailAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/betanalysisdetail"
	ApisToAthenaV2(ctx, api)
}

func getRevenueReportingAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/revenue"
	ApisToAthenaV2(ctx, api)
}

func getMemberDiscountAnalysis(ctx *mvc.Context) {
	api := "/athena/admin/reports/memberdiscountanalysis"
	ApisToAthenaV2(ctx, api)
}

func getWinlostReportEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/winlost"
	ApisToAthenaV2(ctx, api)
}
