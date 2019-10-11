package service

import (
	"coin.merchant/constant"
	"coin.merchant/models"
)

type MerchantCoinService struct {
}

func (service *MerchantCoinService) Add(merchantCoin *models.MerchantCoin) (id int64, err error) {
	// 检查该商户id下是否存在该积分类型的钱包地址
	total, err := merchantCoin.CheckExist(merchantCoin.MerchantId, merchantCoin.Coin)
	if err != nil || total > 0 {
		return total, constant.Merchant_coin_exist_error
	}
	if id, err = merchantCoin.Add(); err != nil {
		return
	}
	return
}
