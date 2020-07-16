package athena

import (
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc/view/json"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
)

func init() {
	biApiWriter = &json.StringWriter{}

	// 風控-產品分析
	mvc.Get("/v1/admin/risk/productanalysis/profitandloss", getProductProfitAndLossEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/profitandloss/chart", getProductProfitAndLossChartEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/time/profitandloss", getTimeProductProfitAndLossEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/new/profitandloss", getNewProductProfitAndLossEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/new/profitandloss/chart", getNewProductProfitAndLossChartEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/contribution/{contribute}", getProductContributionEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/distribution", getProductAnalysisDistributionEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/filter", getProductAnalysisFilterEpt).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/productanalysis/profit/detail", getProductAnalysisProfitDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-入款分析
	mvc.Get("/v1/admin/risk/depositanalysis/chart", riskDepositChart).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/chart/group", riskDepositChartGroup).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail", riskDepositDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/ranking", riskDepositRanking).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/amount", riskDepositDetailAmountRange).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/amount/custom", riskDepositDetailAmountRangeCustom).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/amount/overview", riskDepositDetailAmountRangeOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/transfer", riskWebdepositDetailTransfer).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/merchant", riskWebdepositDetailMerchant).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/depositanalysis/detail/bank", riskDepositDetailBank).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/deposit_and_withdraw/member_detail", riskDepositAndWithdraw).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/deposit/member_detail", riskDepositMemberDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/deposit/member_detail/transfer_method", riskDepositMemberDetailTransferMethod).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/deposit/overview", riskDepositOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-出款分析
	mvc.Get("/v1/admin/risk/withdrawanalysis/chart", riskWithdrawChart).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/chart/group", riskWithdrawChartGroup).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/detail", riskWithdrawDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/ranking", riskWithdrawRanking).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/detail/amount", riskWithdrawDetailAmountRange).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/detail/amount/custom", riskWithdrawDetailAmountRangeCustom).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/detail/amount/overview", riskWithdrawDetailAmountRangeOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdrawanalysis/detail/merchant", riskWebwithdrawDetailMerchant).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdraw/member_detail", riskWebwithdrawMemberDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/withdraw/overview", riskWithdrawOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-會員分析
	mvc.Get("/v1/admin/risk/memberanalysis/chart/age/", riskMemberChartAge).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/distribution/order/", riskMemberDistributionOrder).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/distribution/profit/", riskMemberDistributionProfit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/distribution/deposit/", riskMemberDistributionDeposit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/distribution/withdraw/", riskMemberDistributionWithdraw).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/product/", riskMemberProduct).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/transaction/", riskMemberTransaction).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/chart/", riskMemberChart).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/chart/activate/", riskMemberChartActivate).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/detail/activate/", riskMemberDetailActivate).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/memberanalysis/analysis/", riskMemberAnalysis).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/member_active/member_detail/", riskMemberActiveDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-平台帳戶
	mvc.Get("/v1/admin/risk/platformanalysis/chart/balance", riskPlatformChartBalance).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/platformanalysis/chart/profit", riskPlatformChartProfit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/platformanalysis/detail/profit", riskPlatformDetailProfit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/platform_cost/member_detail", riskPlatformCostDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-實時看板
	mvc.Get("/v1/admin/risk/realtime/order", getRealtimeOrder).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/realtime/earning", getRealtimeEarning).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/realtime/product", getRealtimeProduct).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/realtime/deposit", getRealtimeDeposit).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/realtime/withdraw", getRealtimeWithdraw).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/realtime/member", getRealtimeMember).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 風控-告警
	mvc.Get("/v1/admin/risk/notification/overview", getNotificationOverview).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/notification/{event_id}", getNotificationEvent).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/notification/detail/{event_id}", getNotificationEventDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Put("/v1/admin/risk/notification/{event_id}", updateNotificationThreshold).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	mvc.Get("/v1/athena/admin/risk/memberanalysis/new/trace/", riskMemberTrace).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/memberanalysis/new/trace/detail", riskMemberTraceDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/memberanalysis/new/trace/range", riskMemberTraceRange).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/memberanalysis/new/trace/self_range", riskMemberTraceSelfRange).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 營銷活動
	mvc.Get("/v1/admin/risk/activity/members/", riskGetActivityMembers).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/v1/admin/risk/activity/create_group/", riskCreateActivityGroup).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/activity/group_list/", riskGetActivityGroupList).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/admin/risk/activity/group_member_list/", riskGetActivityGroupMemberList).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/v1/athena/admin/risk/activity/new_member/create_group/", riskActivityNewMemberCreateGroup).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/activity/new_member/count/", riskActivityNewMemberCount).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/v1/athena/admin/risk/activity/overview/{category}/create_group/", riskActivityOverviewCreateGroupEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/activity/overview/{category}/member_count/", riskActivityOverviewMemberCountEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	// 平台總覽
	mvc.Get("/v1/athena/admin/risk/{page}/{category}", riskItemEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/chart/", riskItemChartEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/chart/detail", riskItemChartDetailEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/v1/athena/admin/risk/{page}/{category}/chart/detail/activity/create_group", riskActivityOverviewCreateGroupEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/chart/detail/activity/member_count", riskActivityOverviewMemberCountEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/detail/", riskItemDetailEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/distribution/", riskItemDistributionEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/range/", riskItemRangeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/range/detail/", riskItemRangeDetailEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/self_range/", riskItemSelfRangeEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/distribution/{type}", riskMemberDistributionEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/distribution/{type}/detail", riskMemberDistributionDetailEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Post("/v1/athena/admin/risk/{page}/{category}/distribution/{type}/activity/create_group", riskActivityMemberDistributionCreateGroupEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/distribution/{type}/activity/member_count", riskActivityMemberDistributionMemberCountEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
	mvc.Get("/v1/athena/admin/risk/{page}/{category}/self/distribution/{type}", riskMemberSelfDistributionEndpoint).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)

	mvc.Get("/v1/athena/admin/risk/transaction_member/withdraw_member/profit/detail", riskWithdrawMemberProfitDetail).SetItem(utilsMVC.API_Level, utilsMVC.API_Level_User).SetWriter(&biApiWriter)
}

func getProductProfitAndLossEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/profitandloss"
	ApisToAthenaV2(ctx, api)
}

func getProductProfitAndLossChartEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/profitandloss/chart"
	ApisToAthenaV2(ctx, api)
}

func getTimeProductProfitAndLossEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/time/profitandloss"
	ApisToAthenaV2(ctx, api)
}

func getNewProductProfitAndLossEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/new/profitandloss"
	ApisToAthenaV2(ctx, api)
}

func getNewProductProfitAndLossChartEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/new/profitandloss/chart"
	ApisToAthenaV2(ctx, api)
}

func getProductContributionEpt(ctx *mvc.Context) {
	contribute := ctx.GetRouteValue("contribute", "channel_code")
	api := "/v1/athena/admin/risk/productanalysis/contribution/" + contribute
	ApisToAthenaV2(ctx, api)
}

func getProductAnalysisDistributionEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/distribution"
	ApisToAthenaV2(ctx, api)
}

func getProductAnalysisFilterEpt(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/filter"
	ApisToAthenaV2(ctx, api)
}

func getProductAnalysisProfitDetail(ctx *mvc.Context) {
	api := "/v1/athena/admin/risk/productanalysis/profit/detail"
	ApisToAthenaV2(ctx, api)
}

// ------------------------------------------------------------------------------------------------------

func riskDepositChart(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/chart/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositChartGroup(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/chart/group/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositRanking(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/ranking/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositDetailAmountRange(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/amount/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositDetailAmountRangeCustom(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/amount/custom/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositDetailAmountRangeOverview(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/amount/overview/"
	ApisToAthenaV2(ctx, api)
}

func riskWebdepositDetailTransfer(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/transfer/"
	ApisToAthenaV2(ctx, api)
}

func riskWebdepositDetailMerchant(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/merchant/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositDetailBank(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/detail/bank/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositAndWithdraw(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit_and_withdraw/member_detail/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositMemberDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/member_detail/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositMemberDetailTransferMethod(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/member_detail/transfer_method/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawChart(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/chart/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawChartGroup(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/chart/group/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/detail/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawRanking(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/ranking/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawDetailAmountRange(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/detail/amount/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawDetailAmountRangeCustom(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/detail/amount/custom/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawDetailAmountRangeOverview(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/detail/amount/overview/"
	ApisToAthenaV2(ctx, api)
}

func riskWebwithdrawDetailMerchant(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/detail/merchant/"
	ApisToAthenaV2(ctx, api)
}

func riskWebwithdrawMemberDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/member_detail/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberChartAge(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/chart/age/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionOrder(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/distribution/order/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionProfit(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/distribution/profit/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionDeposit(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/distribution/deposit/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionWithdraw(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/distribution/withdraw/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberProduct(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/product/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberTransaction(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/transaction/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberChart(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/chart/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberChartActivate(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/chart/activate/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDetailActivate(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/detail/activate/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberAnalysis(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/analysis/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberActiveDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/member_active/member_detail/"
	ApisToAthenaV2(ctx, api)
}

func riskPlatformChartBalance(ctx *mvc.Context) {
	api := "/athena/admin/risk/platform/chart/balance/"
	ApisToAthenaV2(ctx, api)
}

func riskPlatformChartProfit(ctx *mvc.Context) {
	api := "/athena/admin/risk/platform/chart/profit/"
	ApisToAthenaV2(ctx, api)
}

func riskPlatformDetailProfit(ctx *mvc.Context) {
	api := "/athena/admin/risk/platform/detail/profit/"
	ApisToAthenaV2(ctx, api)
}

func riskPlatformCostDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/platform_cost/member_detail/"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeOrder(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/order"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeEarning(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/earning"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeProduct(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/product"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeDeposit(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/deposit"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeWithdraw(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/withdraw"
	ApisToAthenaV2(ctx, api)
}

func getRealtimeMember(ctx *mvc.Context) {
	api := "/athena/admin/risk/realtime/member/"
	ApisToAthenaV2(ctx, api)
}

func getNotificationOverview(ctx *mvc.Context) {
	api := "/athena/admin/risk/notification/overview/"
	ApisToAthenaV2(ctx, api)
}

func getNotificationEvent(ctx *mvc.Context) {
	event := ctx.GetRouteValue("event_id", "")
	api := "/athena/admin/risk/notification/" + event
	ApisToAthenaV2(ctx, api)
}

func getNotificationEventDetail(ctx *mvc.Context) {
	event := ctx.GetRouteValue("event_id", "")
	api := "/athena/admin/risk/notification/detail/" + event
	ApisToAthenaV2(ctx, api)
}

func updateNotificationThreshold(ctx *mvc.Context) {
	event := ctx.GetRouteValue("event_id", "")
	api := "/athena/admin/risk/notification/" + event
	ApisToAthenaV2(ctx, api)
}

func riskMemberTrace(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/new/trace/"
	ApisToAthenaV2(ctx, api)
}

func riskDepositOverview(ctx *mvc.Context) {
	api := "/athena/admin/risk/deposit/overview/"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawOverview(ctx *mvc.Context) {
	api := "/athena/admin/risk/withdraw/overview/"
	ApisToAthenaV2(ctx, api)
}

func riskGetActivityMembers(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/members/"
	ApisToAthenaV2(ctx, api)
}

func riskCreateActivityGroup(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/create_group/"
	ApisToAthenaV2(ctx, api)
}

func riskGetActivityGroupList(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/group_list/"
	ApisToAthenaV2(ctx, api)
}

func riskGetActivityGroupMemberList(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/group_member_list/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberTraceDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/new/trace/detail"
	ApisToAthenaV2(ctx, api)
}

func riskMemberTraceRange(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/new/trace/range"
	ApisToAthenaV2(ctx, api)
}

func riskMemberTraceSelfRange(ctx *mvc.Context) {
	api := "/athena/admin/risk/member/new/trace/self_range"
	ApisToAthenaV2(ctx, api)
}

func riskActivityNewMemberCreateGroup(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/new_member/create_group/"
	ApisToAthenaV2(ctx, api)
}

func riskActivityNewMemberCount(ctx *mvc.Context) {
	api := "/athena/admin/risk/activity/new_member/count/"
	ApisToAthenaV2(ctx, api)
}

func riskItemEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/"
	ApisToAthenaV2(ctx, api)
}

func riskItemChartEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/chart/"
	ApisToAthenaV2(ctx, api)
}

func riskItemChartDetailEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/chart/detail"
	ApisToAthenaV2(ctx, api)
}

func riskItemDetailEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/detail/"
	ApisToAthenaV2(ctx, api)
}

func riskItemDistributionEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/distribution/"
	ApisToAthenaV2(ctx, api)
}

func riskItemRangeEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/range/"
	ApisToAthenaV2(ctx, api)
}

func riskItemSelfRangeEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/self_range/"
	ApisToAthenaV2(ctx, api)
}

func riskItemRangeDetailEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/range/detail"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	t := ctx.GetRouteValue("type", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/distribution/" + t + "/"
	ApisToAthenaV2(ctx, api)
}

func riskMemberDistributionDetailEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	t := ctx.GetRouteValue("type", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/distribution/" + t + "/detail"
	ApisToAthenaV2(ctx, api)
}

func riskMemberSelfDistributionEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	t := ctx.GetRouteValue("type", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/self/distribution/" + t + "/"
	ApisToAthenaV2(ctx, api)
}

func riskActivityOverviewCreateGroupEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "overview")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/chart/detail/activity/create_group"
	ApisToAthenaV2(ctx, api)
}

func riskActivityOverviewMemberCountEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "overview")
	category := ctx.GetRouteValue("category", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/chart/detail/activity/member_count"
	ApisToAthenaV2(ctx, api)
}

func riskWithdrawMemberProfitDetail(ctx *mvc.Context) {
	api := "/athena/admin/risk/transaction_member/withdraw_member/profit/detail"
	ApisToAthenaV2(ctx, api)
}

func riskActivityMemberDistributionCreateGroupEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	t := ctx.GetRouteValue("type", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/distribution/" + t + "/activity/create_group"
	ApisToAthenaV2(ctx, api)
}

func riskActivityMemberDistributionMemberCountEndpoint(ctx *mvc.Context) {
	page := ctx.GetRouteValue("page", "")
	category := ctx.GetRouteValue("category", "")
	t := ctx.GetRouteValue("type", "")
	api := "/athena/admin/risk/" + page + "/" + category + "/distribution/" + t + "/activity/member_count"
	ApisToAthenaV2(ctx, api)
}
