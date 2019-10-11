package service

import "coin.merchant/models"

type StatService struct {
}

func (service *StatService) MerchantStat(mid int64) (stat *models.MerchantStat, err error) {
	merchantStat := &models.MerchantStat{}
	return merchantStat.MerchantStat(mid)
}
