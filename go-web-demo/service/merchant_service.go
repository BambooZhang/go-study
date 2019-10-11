package service

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	models2 "common/models"
)

type MerchantService struct {
}

/**
 *新增商户信息
 */
func (service *MerchantService) Add(merchant *models.Merchant) (id int64, err error) {
	total, err := merchant.CheckMerchantName(merchant.MerchantName)
	if err != nil || total > 0 {
		return total, constant.Merchant_name_exist_error
	}

	if id, err = merchant.Add(); err != nil {
		return
	}

	return
}

/**
查询商户列表
*/
func (service *MerchantService) Merchants(page *models2.Page) (merchants []*models.MerchantQueryResult, total int64, err error) {
	total, err = (&models.Merchant{}).MerchantsTotal()
	if err != nil || total == 0 {
		return
	}
	merchants, err = (&models.Merchant{}).Merchants(page)
	return
}

/**
修改商户状态
*/
func (service *MerchantService) UpdateStatus(merchant models.MerchantUpdateStatusUpdateParams) (err error) {
	if err = merchant.UpdateStatus(); err != nil {
		return
	}
	return
}

/**
 *  查询商户信息
 */
func (service *MerchantService) Query(mid int64) (result *models.MerchantQueryResult, err error) {
	merchant := &models.Merchant{}
	return merchant.Query(mid)
}

func (service *MerchantService) Modify(merchant *models.Merchant) (result bool, err error) {
	return merchant.Modify()
}

/**
 * 获取所有商铺的二维码信息
 */
func (service *MerchantService) QrCodes(mid int64) (result []*models.QrCodeManageResult, err error) {
	merchant := &models.Merchant{}
	return merchant.QrCodes(mid)
}
