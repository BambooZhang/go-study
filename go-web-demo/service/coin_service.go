package service

import (
	"coin.merchant/models"
	"common/clog"
	models2 "common/models"
)

type CoinService struct {
}

func (service CoinService) AllIncomes(mid int64, page *models2.Page) (coinDetailQueryResult []*models.MerchantCoinDetailQueryResult, err error) {
	result := models.MerchantCoinDetailQueryResult{}
	results, err := result.AllIncomes(mid, page)
	if err != nil {
		return nil, err
	}
	for _, coinDetail := range results {
		coinDetail.UserPhone = PhoneSensitive(coinDetail.UserPhone, 3, 7)
	}
	return results, nil
}

func (service CoinService) Incomes(mid int64, storeId string, page *models2.Page) (coinDetailQueryResult []*models.MerchantCoinDetailQueryResult, total int64, err error) {
	result := models.MerchantCoinDetailQueryResult{}
	total, err = result.Total(mid, storeId)
	if err != nil {
		clog.Errorf("分页查询商户收入总条数异常,merchant_id=%+v,err=%+v", mid, err)
		return nil, 0, err
	}
	results, err := result.Incomes(mid, storeId, page)
	if err != nil {
		return nil, total, err
	}
	for _, coinDetail := range results {
		coinDetail.UserPhone = PhoneSensitive(coinDetail.UserPhone, 3, 7)
	}
	return results, total, nil
}

/**
 * 查询积分收入总数
 */
func (service CoinService) Total(mid int64) (total float64, err error) {
	detail := &models.MerchantCoinDetail{}
	return detail.Total(mid)
}

/**
 * 手机号码脱敏处理
 */
func PhoneSensitive(phone string, start int, end int) (result string) {
	if len(phone) < end || start < 0 || end < start {
		return phone
	}
	one := phone[0:start]
	two := phone[start:end]
	three := phone[end:]
	Symbol := ""
	for a := 0; a < len(two); a++ {
		Symbol += "*"
	}
	one += Symbol
	one += three
	return one
}
