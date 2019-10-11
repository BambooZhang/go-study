package admin

import (
	merchant_models "coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"common/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type adminMerchantController struct {
	Ctx iris.Context
}

func AminMerchantMvc(mvc *mvc.Application) {
	mvc.Handle(new(adminMerchantController))
}

func (m *adminMerchantController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/merchants", "Merchants")
	b.Handle("POST", "/add", "Add")
	b.Handle("PUT", "/{id:int64}/status", "UpdateStatus")
}

/**
新增商户信息
*/
func (m *adminMerchantController) Add() {
	var merchant merchant_models.Merchant
	m.Ctx.ReadJSON(&merchant)

	service := &service.MerchantService{}
	id, err := service.Add(&merchant)
	if err != nil {
		clog.Infof("新增商户信息失败:merchant=%+v,err=%+v", &merchant, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("id", id))
}

/**
查询商户列表
*/
func (m *adminMerchantController) Merchants() {
	service := &service.MerchantService{}
	pageParams := models.PageParams(m.Ctx)
	merchants, total, err := service.Merchants(pageParams)
	if err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("page", models.NewPageResult(total, pageParams.Size, pageParams.Page, len(merchants), merchants)))
}

/**
更新状态
*/
func (m *adminMerchantController) UpdateStatus(id int64) {
	service := &service.MerchantService{}
	body := merchant_models.MerchantUpdateStatusUpdateParams{}
	if err := m.Ctx.ReadJSON(&body); err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	body.Id = id
	if err := service.UpdateStatus(body); err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}
