package service

import "coin.merchant/models"

type StoreQrCodeService struct {
}

func (service *StoreQrCodeService) StoreQrCodeQuery(qrCode string) (result *models.ShopQrCodeQueryResult, err error) {
	storeQrCode := &models.ShopQrCode{}
	return storeQrCode.StoreQrCodeQuery(qrCode)
}
