package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// [order] general report 一般報表
	mvc.Get("/v1/reseller/reports/generalreport", getGeneralReportResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/generalreport/total", getGeneralReportTotalResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/generalreport/{channelcode}", getGeneralReportChannelResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/generalreport/total/trend", getGeneralReportTotalTrendResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/generalreport/{channelcode}/trend", getGeneralReportChannelTrendResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)

	// 每日報表
	mvc.Get("/v1/reseller/reports/dailyreporting", getDailyReportResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/dailyreporting/deposit", getDailyReportDepositResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/dailyreporting/firstdeposit", getDailyReportFirstDepositResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	mvc.Get("/v1/reseller/reports/dailyreporting/withdraw", getDailyReportWithdrawResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)

}

func getGeneralReportResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/generalreport"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportTotalResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/generalreport/total"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportChannelResEpt(ctx *mvc.Context) {
	channelCode := ctx.GetRouteValue("channelcode", "")
	api := "/v1/athena/reseller/reports/generalreport/" + channelCode
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportTotalTrendResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/generalreport/total/trend"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportChannelTrendResEpt(ctx *mvc.Context) {
	channelCode := ctx.GetRouteValue("channelcode", "")
	api := "/v1/athena/reseller/reports/generalreport/" + channelCode + "/trend"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/dailyreporting"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportDepositResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/dailyreporting/deposit"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportFirstDepositResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/dailyreporting/firstdeposit"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportWithdrawResEpt(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/dailyreporting/withdraw"
	ApisToAthenaV2(ctx, api)
}
