package service

import (
	"coin.merchant/models"
	models2 "common/models"
)

type MerchantCoinDetailService struct {
}

func (service *MerchantCoinDetailService) MerchantCoinDetails(params *models.AdminMerchantCoinDetailQueryParams, page *models2.Page) (merchantCoinDetails []*models.AdminMerchantCoinDetailQueryResult, total int64, err error) {
	total, err = params.MerchantCoinDetailsTotal()
	if err != nil || total == 0 {
		return
	}
	merchantCoinDetails, err = params.MerchantCoinDetails(page)
	if err != nil {
		return
	}
	return
}

func (service *MerchantCoinDetailService) CoinPayStat(params *models.AdminCoinPayTotalQueryParams, page *models2.Page) (coinPayStat []*models.AdminCoinPayStatQueryResult, total int64, err error) {
	total, err = params.CoinPayTotal()
	if err != nil || total == 0 {
		return
	}
	coinPayStat, err = params.CoinPayStat(page)
	if err != nil {
		return
	}
	return
}
