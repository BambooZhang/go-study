package service

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"coin.merchant/utils"
	"common/clog"
	models2 "common/models"
	"encoding/json"
	"fmt"
	"strconv"
)

type StoreService struct {
}

/**
 * 新增商铺
 */
func (service *StoreService) Add(store *models.StoreAddParams) (id int64, err error) {
	sid, err := store.Add()
	if err != nil {
		clog.Errorf(" rank=%+v,生成的商铺id=%+v", store, sid)
		return
	}
	//读取商户支持的积分类型
	coinCofin := &models.MerchantFeeConfig{}
	coins, err := coinCofin.CoinConfig(store.MerchantId)
	if err != nil {
		clog.Errorf("新增商铺,同时生成二维码信息异常 rank=%+v,生成的商铺id=%+v", store, sid)
		return
	}

	// 生成商铺二维码信息
	code := &models.ShopQrCode{}
	for _, coin := range coins {
		//生成唯一二维码key
		key := utils.BuildQrCodeKey(sid)
		codeUrl := fmt.Sprintf(constant.QR_CODE_ACCESSING_URL, key)
		//组装二维码 包含参数信息
		data, err := WrapQrCodeParams(*coin, store.MerchantId, sid)
		qrId, err := code.Add(coin, key, codeUrl, data, store.MerchantId, sid)
		if err != nil {
			clog.Errorf("新增商铺,同时生成二维码信息异常 rank=%+v,生成的商铺id=%+v,二维码id=%+v", store, sid, qrId)
			return 0, nil
		}
	}
	return sid, nil
}

/**
 * 修改商铺信息
 */
func (service *StoreService) Modify(store *models.StoreModifyParams) (err error) {
	return store.Modify()
}

func (service *StoreService) Stores(store *models.Store, mid int64) (results []*models.StoreQueryResult, err error) {
	return store.Stores(mid)
}

func (service *StoreService) Delete(store *models.Store) (err error) {
	return store.DeleteStore()
}

func (service *StoreService) Query(merchantId int64, sid int64) (result *models.StoreQueryDetailResult, err error) {
	store := &models.Store{}
	return store.Query(merchantId, sid)
}

func (service *StoreService) Disable(store *models.Store) (err error) {
	return store.DisableStore()
}

func (service *StoreService) AdminStores(page *models2.Page) (stores []*models.StoreQueryDetailResult, total int64, err error) {
	total, err = (&models.Store{}).StoresTotal()
	if err != nil || total == 0 {
		return
	}
	stores, err = (&models.Store{}).AdminStores(page)
	return
}

/**
 * 封装 二维码内参数
 */
func WrapQrCodeParams(coin string, merchantId int64, storeId int64) (data string, err error) {
	params := map[string]string{}
	params["coin"] = coin
	params["merchantId"] = strconv.FormatInt(merchantId, 10)
	params["storeId"] = strconv.FormatInt(storeId, 10)
	marshal, err := json.Marshal(params)
	if err != nil {
		clog.Errorf("创建二维码参数失败")
		return "", err
	}
	return string(marshal), nil
}
