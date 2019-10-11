package controllers

import (
	"coin.merchant/constant"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type statController struct {
	Ctx iris.Context
}

func StatMvc(mvc *mvc.Application) {
	mvc.Handle(new(statController))
}

func (m *statController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{mid:int64}", "MerchantStat")
}

/**
 *  商户号 信息统计
 */
func (m *statController) MerchantStat(mid int64) {
	statService := &service.StatService{}
	stat, err := statService.MerchantStat(mid)
	if stat != nil {
		m.Ctx.JSON(code.DefaultSuccess().Add("stat", stat))
		return
	}
	if err == nil {
		clog.Errorf("统计商户收入信息 未知错误,mid=%+v", mid)
		m.Ctx.JSON(code.ReturnError(constant.Merchant_stat_error))
		return
	}
	m.Ctx.JSON(code.ReturnError(err))
}
