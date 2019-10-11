package service

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"common/clog"
	models2 "common/models"
)

type MerchantFeeService struct {
}

/**
新增商户积分费率配置
*/
func (service *MerchantFeeService) Add(merchantFee *models.MerchantFeeConfig) (id int64, err error) {
	// 检查该商户id下是否存在该积分类型的钱包地址
	total, err := merchantFee.CheckExist(merchantFee.MerchantId, merchantFee.Coin)
	if err != nil || total > 0 {
		return total, constant.Merchant_coin_fee_exist_error
	}
	// 新增积分费率配置
	if id, err = merchantFee.Add(); err != nil {
		return
	}
	return
}

/**
查询商户积分费率配置列表
*/
func (service *MerchantFeeService) MerchantFees(page *models2.Page) (merchantFees []*models.MerchantFeesQueryResult, total int64, err error) {
	total, err = (&models.MerchantFeeConfig{}).MerchantFeesTotal()
	if err != nil || total == 0 {
		return
	}
	merchantFees, err = (&models.MerchantFeeConfig{}).MerchantFees(page)
	return
}

/**
积分费率配置-修改
*/
func (service *MerchantFeeService) UpdateMerchantFee(merchantFee models.MerchantFeeUpdateUpdateParams) (err error) {
	if err = merchantFee.UpdateMerchantFee(); err != nil {
		return
	}
	return
}

/**
根据id获取交易手续费详情
*/
func (service *MerchantFeeService) MerchantFeeId(id int64) (merchantFee *models.MerchantFeesQueryResult, err error) {
	merchantFee, err = (&models.MerchantFeeConfig{}).MerchantFeeId(id)
	if err != nil {
		clog.Errorf("根据Id查询merchantFee信息失败:id=%d,err=%+v", id, err)
		return nil, constant.Query_merchant_fee_config_error
	}
	return
}

/**
根据merchantId和积分类型获取积分费率配置详情
*/
func (service *MerchantFeeService) MerchantFeeQuery(merchantId int64, coin string) (merchantFee *models.MerchantFeesQueryResult, err error) {
	merchantFee, err = (&models.MerchantFeeConfig{}).MerchantFeeQuery(merchantId, coin)
	if err != nil {
		clog.Errorf("根据Id查询merchantFee信息失败:merchantId=%d,err=%+v", merchantId, err)
		return nil, constant.Merchant_fee_not_exist_error
	}
	return
}
