package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.paradise-soft.com.tw/data-glob/injection/injection"
	"gitlab.paradise-soft.com.tw/data/metis/models/dto"
	"gitlab.paradise-soft.com.tw/data/metis/service"
	"gitlab.paradise-soft.com.tw/data/metis/transport/http/common"
	"gitlab.paradise-soft.com.tw/data/metis/transport/http/middleware"
)

func init() {
	injection.AutoRegister(&AdminFinancialController{})
}

// AdminFinancialController ...
type AdminFinancialController struct {
	AdminService service.IAdminFinancialService `injection:"AdminFinancialService"`
}

// SetupRouter ...
func (ctl *AdminFinancialController) SetupRouter(router *gin.Engine) {
	metis := router.Group("v1/metis")
	metis.Use(middleware.Authenticate())
	{
		// 現金系統-財務報表
		metis.GET("/reports/financial/overview", ctl.Overview)
		// 公司入款
		metis.GET("/reports/financial/deposititems", ctl.DepositItems)
		// metis.GET("/reports/financial/deposititems/:account", ctl.Overview)
		// 線上支付手續費
		metis.GET("/reports/financial/webdepositfeeitems", ctl.WebDepositFeeItems)
		// metis.GET("/reports/financial/webdepositfeeitems/{merchant_id}", ctl.Overview)
		// 人工存入
		metis.GET("/reports/financial/manualdeposititems", ctl.MandepositItems)
		// metis.GET("/reports/financial/manualdeposititems/{action_code}", ctl.Overview)
		// 人工提出
		metis.GET("/reports/financial/manualwithdrawitems", ctl.ManwithdrawItems)
		// metis.GET("/reports/financial/manualwithdrawitems/{action_code}", ctl.Overview)
		// 線上支付
		metis.GET("/reports/financial/webdeposititems", ctl.WebDepositItems)
		// metis.GET("/reports/financial/webdeposititems/{merchant_id}", ctl.Overview)
		// 出款申請
		metis.GET("/reports/financial/withdrawitems", ctl.WithdrawItems)
		// metis.GET("/reports/financial/withdrawitems/{group}", ctl.Overview)
		// metis.GET("/reports/financial/withdrawitems/{group}/{account}", ctl.Overview)
		// 會員出款被扣除金額
		metis.GET("/reports/financial/withdrawfeeitems", ctl.WithdrawFeeItems)
		// metis.GET("/reports/financial/withdrawfeeitems/{charge_type}", ctl.Overview)
		// 給予優惠
		metis.GET("/reports/financial/discountitems", ctl.Discount)
		// metis.GET("/reports/financial/discountitems/{deposit_type}", ctl.Overview)
		// metis.GET("/reports/financial/discountitems/{deposit_type}/{action_code}", ctl.Overview)
		// 給予反水
		metis.GET("/reports/financial/rebate/overview", ctl.RebateOverview)
		metis.GET("/reports/financial/rebate/realtime", ctl.RebateRealtime)
		// metis.GET("/reports/financial/rebate/realtime/items", ctl.Overview)
		metis.GET("/reports/financial/rebate/ladder", ctl.RebateLadder)
		metis.GET("/reports/financial/rebate/promo", ctl.RebatePromo)
	}
}

// Overview ...
func (ctl *AdminFinancialController) Overview(ctx *gin.Context) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	var ret interface{}
	if conds.IsExportMode == 1 {
		ret, err = ctl.AdminService.OverviewExport(conds.Brand, conds)
	} else {
		ret, err = ctl.AdminService.Overview(conds.Brand, conds)
	}

	if err != nil {
		common.Error(ctx, err)
		return
	}

	common.Send(ctx, ret)
}

// Items ...
func (ctl *AdminFinancialController) Items(ctx *gin.Context, category string) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	conds.Category = category

	ret, err := ctl.AdminService.Items(conds.Brand, conds)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	if ret == nil {
		common.Send(ctx, []string{})
		return
	}

	common.Send(ctx, ret)
}

// DepositItems ...
func (ctl *AdminFinancialController) DepositItems(ctx *gin.Context) {
	ctl.Items(ctx, "deposit")
}

// WebDepositFeeItems ...
func (ctl *AdminFinancialController) WebDepositFeeItems(ctx *gin.Context) {
	ctl.Items(ctx, "webdepositfee")
}

// MandepositItems ...
func (ctl *AdminFinancialController) MandepositItems(ctx *gin.Context) {
	ctl.Items(ctx, "mandeposit")
}

// ManwithdrawItems ...
func (ctl *AdminFinancialController) ManwithdrawItems(ctx *gin.Context) {
	ctl.Items(ctx, "manwithdraw")
}

// WebDepositItems ...
func (ctl *AdminFinancialController) WebDepositItems(ctx *gin.Context) {
	ctl.Items(ctx, "webdeposit")
}

// WithdrawItems ...
func (ctl *AdminFinancialController) WithdrawItems(ctx *gin.Context) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	conds.Category = "withdraw"

	ret, err := ctl.AdminService.WithdrawItems(conds.Brand, conds)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	if ret == nil {
		common.Send(ctx, []string{})
		return
	}

	common.Send(ctx, ret)
}

// WithdrawFeeItems ...
func (ctl *AdminFinancialController) WithdrawFeeItems(ctx *gin.Context) {
	ctl.Items(ctx, "withdrawfee")
}

// Discount ...
func (ctl *AdminFinancialController) Discount(ctx *gin.Context) {
	ctl.Items(ctx, "discount")
}

// RebateOverview ...
func (ctl *AdminFinancialController) RebateOverview(ctx *gin.Context) {
	ctl.Items(ctx, "rebate")
}

// RebateRealtime ...
func (ctl *AdminFinancialController) RebateRealtime(ctx *gin.Context) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	ret, err := ctl.AdminService.RebateRealtime(conds.Brand, conds)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	if ret == nil {
		common.Send(ctx, []string{})
		return
	}

	common.Send(ctx, ret)
}

// RebateLadder ...
func (ctl *AdminFinancialController) RebateLadder(ctx *gin.Context) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	ret, err := ctl.AdminService.RebateLadder(conds.Brand, conds)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	if ret == nil {
		common.Send(ctx, []string{})
		return
	}

	common.Send(ctx, ret)
}

// RebatePromo ...
func (ctl *AdminFinancialController) RebatePromo(ctx *gin.Context) {
	conds := &dto.FinancialReportingConds{}
	err := conds.Init(ctx)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	ret, err := ctl.AdminService.RebatePromo(conds.Brand, conds)
	if err != nil {
		common.Error(ctx, err)
		return
	}

	if ret == nil {
		common.Send(ctx, []string{})
		return
	}

	common.Send(ctx, ret)
}
