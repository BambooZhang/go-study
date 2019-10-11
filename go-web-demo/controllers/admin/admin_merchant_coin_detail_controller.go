package admin

import (
	models2 "coin.merchant/models"
	"coin.merchant/service"
	"common/code"
	"common/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type adminMerchantCoinDetailController struct {
	Ctx iris.Context
}

func AminMerchantCoinDetailMvc(mvc *mvc.Application) {
	mvc.Handle(new(adminMerchantCoinDetailController))
}

func (m *adminMerchantCoinDetailController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/coin/details", "Details")
	b.Handle("GET", "/coin/pay/stat", "CoinPayStat")
}

func (m *adminMerchantCoinDetailController) Details() {
	service := &service.MerchantCoinDetailService{}
	pageParams := models.PageParams(m.Ctx)
	params := &models2.AdminMerchantCoinDetailQueryParams{
		BeginTime:    m.Ctx.URLParam("beginTime"),
		EndTime:      m.Ctx.URLParam("endTime"),
		MerchantName: m.Ctx.URLParam("merchantName"),
		Coin:         m.Ctx.URLParam("coin"),
		UserPhone:    m.Ctx.URLParam("userPhone"),
		Type:         m.Ctx.URLParam("type"),
	}
	merchantCoinDetails, total, _ := service.MerchantCoinDetails(params, pageParams)
	m.Ctx.JSON(code.DefaultSuccess().Add("page", models.NewPageResult(total, pageParams.Size, pageParams.Page, len(merchantCoinDetails), merchantCoinDetails)))
}

func (m *adminMerchantCoinDetailController) CoinPayStat() {
	service := &service.MerchantCoinDetailService{}
	pageParams := models.PageParams(m.Ctx)
	params := &models2.AdminCoinPayTotalQueryParams{
		Coin:         m.Ctx.URLParam("coin"),
		PayTimeMonth: m.Ctx.URLParam("payTimeMonth"),
		FormatDate:   m.Ctx.URLParam("formatDate"),
	}
	coinPayStat, total, _ := service.CoinPayStat(params, pageParams)
	m.Ctx.JSON(code.DefaultSuccess().Add("page", models.NewPageResult(total, pageParams.Size, pageParams.Page, len(coinPayStat), coinPayStat)))
}
