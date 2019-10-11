package controllers

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type merchantController struct {
	Ctx iris.Context
}

func MerchantMvc(mvc *mvc.Application) {
	mvc.Handle(new(merchantController))
}

func (m *merchantController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{mid:int64}", "Query")
	b.Handle("PUT", "/{mid:int64}", "Modify")
	b.Handle("GET", "/qrcodes/{mid:int64}}", "QrCodes")
}

/**
 * 查询商户信息
 */
func (m *merchantController) Query(mid int64) {
	merchantService := &service.MerchantService{}
	result, err := merchantService.Query(mid)
	if err != nil {
		clog.Infof("查询商户信息失败:mid=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("merchant", result))
}

/**
 * 修改商铺信息
 */
func (m *merchantController) Modify(mid int64) {
	merchant := &models.Merchant{}
	if err := m.Ctx.ReadJSON(merchant); err != nil {
		clog.Infof("修改商户失败:mid=%+v,err=%+v", mid, err)
		return
	}
	merchant.Id = mid
	merchantService := &service.MerchantService{}
	result, err := merchantService.Modify(merchant)
	if err != nil {
		clog.Infof("修改商户失败:mid=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	if !result {
		clog.Infof("修改商户失败:mid=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(constant.Modify_merchant_error))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("merchant", result))
}

/***
 * 二维码管理
 */
func (m *merchantController) QrCodes(mid int64) {
	merchantService := &service.MerchantService{}
	result, err := merchantService.QrCodes(mid)
	if err != nil {
		clog.Infof("修改商户失败:mid=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("QrCodes", result))
}
