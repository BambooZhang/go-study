package admin

import (
	"coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	common_models "common/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type adminMerchantStoreController struct {
	Ctx iris.Context
}

func AminMerchantStoreMvc(mvc *mvc.Application) {
	mvc.Handle(new(adminMerchantStoreController))
}

func (m *adminMerchantStoreController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/store/add", "Add")
	b.Handle("PUT", "/store/{id:long}", "Modify")
	b.Handle("DELETE", "/store/{id:long}", "Delete")
	b.Handle("GET", "/store/{id:long}", "Query")
	b.Handle("PUT", "/store/status/{id:long}", "Disable")
	b.Handle("GET", "/store/stores", "Stores")
}

/**
商户店铺管理-新增
*/
func (m *adminMerchantStoreController) Add() {
	store := &models.StoreAddParams{}
	if err := m.Ctx.ReadJSON(store); err != nil {
		clog.Errorf("商户店铺管理-新增失败:store=%+v,err=%+v", &store, err)
		return
	}
	service := &service.StoreService{}
	id, err := service.Add(store)
	if err != nil {
		clog.Infof("商户店铺管理-新增失败:store=%+v,err=%+v", &store, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("id", id))
}

/**
商户店铺管理-修改
*/
func (m *adminMerchantStoreController) Modify(id int64) {
	store := &models.StoreModifyParams{}
	if err := m.Ctx.ReadJSON(&store); err != nil {
		clog.Infof("商户店铺管理-修改失败:store=%+v,err=%+v", &store, err)
		return
	}
	store.Id = id
	service := &service.StoreService{}
	err := service.Modify(store)
	if err != nil {
		clog.Infof("商户店铺管理-修改失败:merchant=%+v,err=%+v", &store, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
商户店铺管理-删除
*/
func (m *adminMerchantStoreController) Delete(id int64) {
	store := &models.Store{}
	if err := m.Ctx.ReadJSON(&store); err != nil {
		clog.Infof("商户店铺管理-删除失败:store=%+v,err=%+v", &store, err)
		return
	}
	store.Id = id
	service := &service.StoreService{}
	err := service.Delete(store)
	if err != nil {
		clog.Infof("商户店铺管理-删除失败:store=%+v,err=%+v", &id, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
商户店铺管理-详情
*/
func (m *adminMerchantStoreController) Query(id int64) {
	merchantId, err := m.Ctx.URLParamInt64("merchantId")
	if err != nil {
		clog.Infof("商户店铺管理-详情失败,商户id传递错误:store=%+v,merchantId=%+v,err=%+v", id, merchantId, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	service := &service.StoreService{}
	result, err := service.Query(merchantId, id)
	if err != nil {
		clog.Infof("商户店铺管理-详情失败:store=%+v,err=%+v", &id, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("store", result))
}

/**
商户店铺管理-禁用
*/
func (m *adminMerchantStoreController) Disable(id int64) {
	store := &models.Store{}
	if err := m.Ctx.ReadJSON(&store); err != nil {
		clog.Infof("商户店铺管理-禁用失败:store=%+v,err=%+v", &store, err)
		return
	}
	store.Id = id
	service := &service.StoreService{}
	err := service.Disable(store)
	if err != nil {
		clog.Infof("商户店铺管理-禁用失败:store=%+v,err=%+v", &id, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
商户店铺管理-列表
*/
func (m *adminMerchantStoreController) Stores() {
	service := &service.StoreService{}
	pageParams := common_models.PageParams(m.Ctx)
	stores, total, err := service.AdminStores(pageParams)
	if err != nil {
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("page", common_models.NewPageResult(total, pageParams.Size, pageParams.Page, len(stores), stores)))
}
