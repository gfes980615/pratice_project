package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	//運營管理

	// 運營管理-注單中心
	// 注單中心
	mvc.Get("/admin/order", orderEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 派彩中心
	mvc.Get("/admin/order/searchpayout", searchPayout).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// 注單校驗
	mvc.Get("/admin/order/validation", orderValidationQueryEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/order/validation/{number}", orderValidationEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	// tool
	mvc.Get("/admin/cache/order", cacheServerOrderEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/cache/searchpayout", cacheServerSearchPayout).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 運營管理-即時注單
	mvc.Get("/admin/reports/realtimeorder", getRealtimeReportAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/realtimeorder/overview", getRealtimeReportOverviewAdmEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/realtimeorder/lotteries/available", getAvailableLotteriesEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/reports/realtimeorder/lotteries/overview", getLotteriesOverviewEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

}

func orderEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/order"
	ApisToAthenaV2(ctx, api)
}

func searchPayout(ctx *mvc.Context) {
	api := "/athena/admin/order/searchpayout"
	ApisToAthenaV2(ctx, api)
}

func orderValidationQueryEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/order/validation"
	ApisToAthenaV2(ctx, api)
}

func orderValidationEndpoint(ctx *mvc.Context) {
	number := ctx.GetRouteValue("number", "")
	api := "/athena/admin/order/validation/" + number
	ApisToAthenaV2(ctx, api)
}

func cacheServerOrderEndpoint(ctx *mvc.Context) {
	api := "/athena/admin/cache/order"
	ApisToAthenaV2(ctx, api)
}

func cacheServerSearchPayout(ctx *mvc.Context) {
	api := "/athena/admin/cache/searchpayout"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeReportAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/realtimeorder"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeReportOverviewAdmEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/realtimeorder/overview"
	ApisToAthenaV2(ctx, api)
}

func getAvailableLotteriesEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/realtimeorder/lotteries/available"
	ApisToAthenaV2(ctx, api)
}

func getLotteriesOverviewEpt(ctx *mvc.Context) {
	api := "/athena/admin/reports/realtimeorder/lotteries/overview"
	ApisToAthenaV2(ctx, api)
}
