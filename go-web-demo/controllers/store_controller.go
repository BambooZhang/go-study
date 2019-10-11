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

type storeController struct {
	Ctx iris.Context
}

func StoreMvc(mvc *mvc.Application) {
	mvc.Handle(new(storeController))
}

func (m *storeController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Create")
	b.Handle("PUT", "/{id:long}", "Modify")
	b.Handle("GET", "/{mid:long}", "Stores")
	b.Handle("DELETE", "/{sid:long}", "Delete")
	b.Handle("GET", "/detail/{sid:long}", "Query")
}

/**
 * 创建商铺
 */
func (m *storeController) Create() {
	store := &models.StoreAddParams{}
	if err := m.Ctx.ReadJSON(store); err != nil {
		clog.Errorf("新增商铺失败:store=%+v,err=%+v", &store, err)
		m.Ctx.JSON(code.ReturnError(constant.Params_error))
		return
	}
	service := &service.StoreService{}
	id, err := service.Add(store)
	if err != nil {
		clog.Infof("新增商铺失败:store=%+v,err=%+v", &store, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("id", id))
}

/**
 * 修改商铺
 */
func (m *storeController) Modify(id int64) {
	store := &models.StoreModifyParams{}
	if err := m.Ctx.ReadJSON(&store); err != nil {
		clog.Infof("修改商铺失败:store=%+v,err=%+v", &store, err)
		return
	}
	store.Id = id

	service := &service.StoreService{}
	err := service.Modify(store)
	if err != nil {
		clog.Infof("新增商户信息失败:merchant=%+v,err=%+v", &store, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
 *  店铺列表
 */
func (m *storeController) Stores(mid int64) {
	store := &models.Store{}
	service := &service.StoreService{}
	stores, err := service.Stores(store, mid)
	if err != nil {
		clog.Infof("新增商户信息失败:merchantId=%+v,err=%+v", mid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("list", stores))
}

/**
 * 删除店铺
 * @param sid 商铺id
 */
func (m *storeController) Delete(sid int64) {
	store := &models.Store{}
	if err := m.Ctx.ReadJSON(&store); err != nil {
		clog.Infof("删除商铺失败:store=%+v,err=%+v", &store, err)
		return
	}
	store.Id = sid
	service := &service.StoreService{}
	err := service.Delete(store)
	if err != nil {
		clog.Infof("删除商铺失败:store=%+v,err=%+v", &sid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
 * 商户查询自己名下 单个商铺详细信息
 * @param sid 商铺id
 */
func (m *storeController) Query(sid int64) {
	merchantId, err := m.Ctx.URLParamInt64("merchantId")
	if err != nil {
		clog.Infof("查询单个商铺详情失败,商户id传递错误:store=%+v,merchantId=%+v,err=%+v", sid, merchantId, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	service := &service.StoreService{}
	result, err := service.Query(merchantId, sid)
	if err != nil {
		clog.Infof("查询单个商铺详情失败:store=%+v,err=%+v", &sid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}

	m.Ctx.JSON(code.DefaultSuccess().Add("store", result))
}
