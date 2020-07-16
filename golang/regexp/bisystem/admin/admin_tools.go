package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// order agg and txn agg by day
	mvc.Post("/admin/reports/aggbyday", aggByDayEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Del("/admin/reports/aggbyday", deleteAggByDay).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// [sync] 同步分析表ＡＰＩ analysis table
	mvc.Post("/admin/reports/analysistable/{sync_code}/sync", syncAnalysisTableAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/admin/reports/analysistable/{sync_code}/dailyautosync", syncAnalysisTableDailyautosyncAdmEpt).SetAuthorizer(&utilsMVC.Anonymous).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/analysistable/{analysis_table}/status", getAnalysisTableStatusAdmEpt).SetAuthorizer(&utilsMVC.Anonymous).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/analysistable/syncfailur", getAnalysisTableSyncFailure).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	mvc.Get("/admin/reporting/financial/verify/{source}/source", VerifyFinancialAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/financial/verify/es", VerifyFinancialESAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/financial/verify/analysis", VerifyFinancialAnalysisAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reporting/financial/verify", VerifyFinancialESAndAnalysisAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// db es compare
	mvc.Get("/admin/reports/dbescompare", getDbEsCompareEndPoint).SetAuthorizer(&utilsMVC.Anonymous).SetWriter(&biApiWriter)
}

func aggByDayEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/aggbyday"
	ApisToAthenaV2(ctx, api)
}

func deleteAggByDay(ctx *mvc.Context) {
	api := "/athena/admin/reports/aggbyday"
	ApisToAthenaV2(ctx, api)
}

func getAnalysisTableStatusAdmEpt(ctx *mvc.Context) {
	analysisTable := ctx.GetRouteValue("analysis_table", "")
	api := "/athena/admin/reports/analysistable/" + analysisTable + "/status"
	ApisToAthenaV2(ctx, api)
}

func syncAnalysisTableAdmEpt(ctx *mvc.Context) {
	sync_code := ctx.GetRouteValue("sync_code", "")
	api := "/athena/admin/reports/analysistable/" + sync_code + "/sync"
	ApisToAthenaV2(ctx, api)
}

func syncAnalysisTableDailyautosyncAdmEpt(ctx *mvc.Context) {
	sync_code := ctx.GetRouteValue("sync_code", "")
	api := "/athena/admin/reports/analysistable/" + sync_code + "/dailyautosync"
	ApisToAthenaV2(ctx, api)
}

func getAnalysisTableSyncFailure(ctx *mvc.Context) {
	api := "/athena/admin/reports/analysistable/syncfailur"
	ApisToAthenaV2(ctx, api)
}

func VerifyFinancialAdmEpt(ctx *mvc.Context) {
	source := ctx.GetRouteValue("source", "")
	api := "/athena/admin/reporting/financial/verify/" + source + "/source"
	ApisToAthenaV2(ctx, api)
}

func VerifyFinancialESAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/financial/verify/es"
	ApisToAthenaV2(ctx, api)
}

func VerifyFinancialAnalysisAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/financial/verify/analysis"
	ApisToAthenaV2(ctx, api)
}

func VerifyFinancialESAndAnalysisAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reporting/financial/verify"
	ApisToAthenaV2(ctx, api)
}

func getDbEsCompareEndPoint(ctx *mvc.Context) {
	api := "/athena/admin/reports/dbescompare"
	ApisToAthenaV2(ctx, api)
}
