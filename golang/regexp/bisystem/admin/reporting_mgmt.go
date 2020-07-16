package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// 報表管理

	// 一般報表
	mvc.Get("/v1/admin/reports/generalreport", getGeneralReportAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/reports/generalreport/total", getGeneralReportTotalAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/reports/generalreport/{channelcode}", getGeneralReportChannelAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/reports/generalreport/total/trend", getGeneralReportTotalTrendAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/reports/generalreport/{channelcode}/trend", getGeneralReportChannelTrendAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 有效會員
	mvc.Get("/v1/reports/effectivemember", getEffectiveMemberReportEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 毛利分析
	mvc.Get("/admin/reporting/orderstats", getProfitReportingAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/orderstats/member", getProfitReportingByMemberAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/orderstats/agent", getProfitReportingByAgentAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/orderstats/product", getProfitReportingByProductAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 頻道報表
	mvc.Get("/v1/reports/order/categories", getOrderReportByCategoryEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/order/types", getOrderReportByTypeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/reports/order/members", getOrderReportByMemberEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 每日報表
	// mvc.Get("/admin/reports/dailyreporting", getDailyReportAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/dailyreporting/deposit", getDailyReportDepositAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/dailyreporting/withdraw", getDailyReportWithdrawAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 水位報表
	mvc.Get("/admin/reports/probability", getProbabilityReportEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/probability/productdict", getProductDictEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Put("/admin/reports/probability", updateProbabilityEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 彩票分析
	mvc.Get("/admin/reports/lotteryanalysis/lottery", getLotteryAnalysisLotteryEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/lotteryanalysis/lotteries/overview", getLotteryAnalysisLotteriesOverviewEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/lotteryanalysis/playcodes/overview", getLotteryAnalysisPlayCodesOverviewEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/lotteryanalysis/playcode", getLotteryAnalysisPlayCodeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/lotteryanalysis/lotteries/available", getLotteryAnalysisLotteriesAvailableEndPoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 紅包－>掃雷－>自訂群
	mvc.Get("/admin/reports/hbsl/{room_code}/memberlist", getHbslMemberListByRoom).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 管理首页->营销中心->手气红包
	mvc.Get("/admin/reports/hbsq/{id}/stats", getHbsqStatsByID).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/hbsq/{id}/rulestats", getHbsqRuleStatsByHbsqID).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

}

func getGeneralReportAdmEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/reports/generalreport"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportTotalAdmEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/reports/generalreport/total"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportChannelAdmEpt(ctx *mvc.Context) {
	channelCode := ctx.GetRouteValue("channelcode", "")
	api := "/v1/athena/admin/reports/generalreport/" + channelCode
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportTotalTrendAdmEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/reports/generalreport/total/trend"
	ApisToAthenaV2(ctx, api)
}

func getGeneralReportChannelTrendAdmEpt(ctx *mvc.Context) {
	channelCode := ctx.GetRouteValue("channelcode", "")
	api := "/v1/athena/admin/reports/generalreport/" + channelCode + "/trend"
	ApisToAthenaV2(ctx, api)
}

func getEffectiveMemberReportEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/effectivemember"
	ApisToAthenaV2(ctx, api)
}

func getProfitReportingAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/orderstats"
	ApisToAthenaV2(ctx, api)
}

func getProfitReportingByMemberAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/orderstats/member"
	ApisToAthenaV2(ctx, api)
}

func getProfitReportingByAgentAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/orderstats/agent"
	ApisToAthenaV2(ctx, api)
}

func getProfitReportingByProductAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/orderstats/product"
	ApisToAthenaV2(ctx, api)
}

func getOrderReportByMemberEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/order/members"
	ApisToAthenaV2(ctx, api)
}

func getOrderReportByTypeEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/order/types"
	ApisToAthenaV2(ctx, api)
}

func getOrderReportByCategoryEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reports/order/categories"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/dailyreporting"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportDepositAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/dailyreporting/deposit"
	ApisToAthenaV2(ctx, api)
}

func getDailyReportWithdrawAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/dailyreporting/withdraw"
	ApisToAthenaV2(ctx, api)
}

func getProbabilityReportEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/probability"
	ApisToAthenaV2(ctx, api)
}

func getProductDictEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/probability/productdict"
	ApisToAthenaV2(ctx, api)
}

func updateProbabilityEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/probability"
	ApisToAthenaV2(ctx, api)
}

func getLotteryAnalysisLotteryEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/lotteryanalysis/lottery"
	ApisToAthenaV2(ctx, api)
}

func getLotteryAnalysisLotteriesOverviewEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/lotteryanalysis/lotteries/overview"
	ApisToAthenaV2(ctx, api)
}

func getLotteryAnalysisPlayCodesOverviewEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/lotteryanalysis/playcodes/overview"
	ApisToAthenaV2(ctx, api)
}

func getLotteryAnalysisPlayCodeEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/lotteryanalysis/playcode"
	ApisToAthenaV2(ctx, api)
}

func getLotteryAnalysisLotteriesAvailableEndPoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/lotteryanalysis/lotteries/available"
	ApisToAthenaV2(ctx, api)
}

func getHbslMemberListByRoom(ctx *mvc.Context) {
	roomCode := ctx.GetRouteValue("room_code", "")
	api := "/athena/admin/reports/hbsl/" + roomCode + "/memberlist"
	ApisToAthenaV2(ctx, api)
}

func getHbsqStatsByID(ctx *mvc.Context) {
	id := ctx.GetRouteValue("id", "")
	api := "/athena/admin/reports/hbsq/" + id + "/stats"
	ApisToAthenaV2(ctx, api)
}

func getHbsqRuleStatsByHbsqID(ctx *mvc.Context) {
	id := ctx.GetRouteValue("id", "")
	api := "/athena/admin/reports/hbsq/" + id + "/rulestats"
	ApisToAthenaV2(ctx, api)
}
