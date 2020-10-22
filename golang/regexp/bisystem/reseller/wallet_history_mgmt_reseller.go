package apis

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func initWalletHistoryReseller() {
	mvc.Get("/reseller/wallethistory", getWalletHistoryReseller).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_Agent).SetWriter(&biApiWriter)
}

func getWalletHistoryReseller(ctx *mvc.Context) {
	api := "/athena/reseller/wallethistory"
	ApisToAthenaV2(ctx, api)
}
