package apis

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func initWalletHistoryAdmin() {
	mvc.Get("/admin/wallethistory", getWalletHistoryAdmin).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/admin/wallethistory/daily", getWalletDailyHistoryAdmin).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	mvc.Post("/admin/wallethistory", writeWalletHistoryAdmin).SetAuthorizer(&utilsMVC.Anonymous).SetWriter(&biApiWriter)
}

func getWalletHistoryAdmin(ctx *mvc.Context) {
	api := "/athena/admin/wallethistory"
	ApisToAthenaV2(ctx, api)
}

func getWalletDailyHistoryAdmin(ctx *mvc.Context) {
	api := "/athena/admin/wallethistory/daily"
	ApisToAthenaV2(ctx, api)
}

func writeWalletHistoryAdmin(ctx *mvc.Context) {
	api := "/athena/admin/wallethistory"
	ApisToAthenaV2(ctx, api)
}
