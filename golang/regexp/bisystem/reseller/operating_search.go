package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// 有效會員
	mvc.Get("/v1/reseller/reports/effectivemember", getResellerEffectivememberReportEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
	// 注單查詢
	mvc.Get("/reseller/order", orderResEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
}

func getResellerEffectivememberReportEndpoint(ctx *mvc.Context) {
	api := "/v1/athena/reseller/reports/effectivemember"
	ApisToAthenaV2(ctx, api)
}

func orderResEpt(ctx *mvc.Context) {
	api := "/athena/reseller/order"
	ApisToAthenaV2(ctx, api)
}
