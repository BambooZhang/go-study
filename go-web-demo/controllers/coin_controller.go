package controllers

import (
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"common/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type coinController struct {
	Ctx iris.Context
}

func CoinMvc(mvc *mvc.Application) {
	mvc.Handle(new(coinController))
}

func (m *coinController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/incomes/{mid:int64}", "Incomes")
	b.Handle("GET", "/total/{mid:int64}", "Total")

}

func (m *coinController) Incomes(mid int64) {
	storeId := m.Ctx.URLParamDefault("storeId", "")
	pageParams := models.PageParams(m.Ctx)
	coinService := &service.CoinService{}

	//查询指定商铺的流水
	coinDetails, total, err := coinService.Incomes(mid, storeId, pageParams)
	if err != nil {
		clog.Errorf("商户查询指定商铺 积分收入信息异常:mid=%+v,storeId=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	//m.Ctx.JSON(code.DefaultSuccess().Add("coins", coinDetails))
	m.Ctx.JSON(code.DefaultSuccess().Add("page", models.NewPageResult(total, pageParams.Size, pageParams.Page, len(coinDetails), coinDetails)))
}

func (m *coinController) Total(mid int64) {
	coinService := &service.CoinService{}

	//查询指定商铺的流水
	total, err := coinService.Total(mid)
	if err != nil {
		clog.Errorf("商户积分总收入查询异常:mid=%+v,storeId=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	//m.Ctx.JSON(code.DefaultSuccess().Add("coins", coinDetails))
	m.Ctx.JSON(code.DefaultSuccess().Add("total", total))
}
