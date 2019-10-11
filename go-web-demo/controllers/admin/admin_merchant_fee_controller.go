package admin

import (
	merchant_fee_models "coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"common/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type adminMerchantFeeController struct {
	Ctx iris.Context
}

func AminMerchantFeeMvc(mvc *mvc.Application) {
	mvc.Handle(new(adminMerchantFeeController))
}

func (m *adminMerchantFeeController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/fee/config/{id:long}", "MerchantFeeId")
	b.Handle("GET", "/fee/configs", "MerchantFees")
	b.Handle("POST", "/fee/config", "Add")
	b.Handle("PUT", "/fee/config/{id:long}", "UpdateMerchantFee")
	b.Handle("GET", "/fee/config/{merchantId:long}/{coin:string}", "MerchantFeeQuery")
}

/**
新增商户积分费率配置
*/
func (m *adminMerchantFeeController) Add() {
	var merchantFee merchant_fee_models.MerchantFeeConfig
	m.Ctx.ReadJSON(&merchantFee)

	service := &service.MerchantFeeService{}
	id, err := service.Add(&merchantFee)
	if err != nil {
		clog.Infof("新增商户积分费率配置失败:merchantFee=%+v,err=%+v", &merchantFee, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("id", id))
}

/**
查询商户积分费率配置
*/
func (m *adminMerchantFeeController) MerchantFees() {
	service := &service.MerchantFeeService{}
	pageParams := models.PageParams(m.Ctx)
	merchantFees, total, err := service.MerchantFees(pageParams)
	if err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("page", models.NewPageResult(total, pageParams.Size, pageParams.Page, len(merchantFees), merchantFees)))
}

/**
积分费率配置-修改
*/
func (m *adminMerchantFeeController) UpdateMerchantFee(id int64) {
	service := &service.MerchantFeeService{}
	body := merchant_fee_models.MerchantFeeUpdateUpdateParams{}
	if err := m.Ctx.ReadJSON(&body); err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	body.Id = id
	if err := service.UpdateMerchantFee(body); err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
积分费率配置-详情
*/
func (m *adminMerchantFeeController) MerchantFeeId(id int64) {
	service := &service.MerchantFeeService{}
	feeData, err := service.MerchantFeeId(id)
	if err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("result", feeData))
}

/**
根据merchantId和积分类型获取积分费率配置详情
*/
func (m *adminMerchantFeeController) MerchantFeeQuery(merchantId int64, coin string) {
	service := &service.MerchantFeeService{}
	feeData, err := service.MerchantFeeQuery(merchantId, coin)
	if err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("result", feeData))
}
