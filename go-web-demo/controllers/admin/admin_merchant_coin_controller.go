package admin

import (
	"coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type adminMerchantCoinController struct {
	Ctx iris.Context
}

func AminMerchantCoinMvc(mvc *mvc.Application) {
	mvc.Handle(new(adminMerchantCoinController))
}

func (m *adminMerchantCoinController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/coin/add", "Add")
}

/**
新增商户信息
*/
func (m *adminMerchantCoinController) Add() {
	var merchantCoin models.MerchantCoin
	m.Ctx.ReadJSON(&merchantCoin)

	service := &service.MerchantCoinService{}
	id, err := service.Add(&merchantCoin)
	if err != nil {
		clog.Infof("新增商户积分信息失败:merchantCoin=%+v,err=%+v", &merchantCoin, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("id", id))
}
