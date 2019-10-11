package models

import (
	"coin.merchant/conn"
	"coin.merchant/constant"
	"common/clog"
	"common/parse"
	"time"
)

/**
 * 商铺 二维码
 */
type ShopQrCode struct {
	Id         int64              `json:"id"`
	MerchantId int64              `json:"merchantId"`
	StoreId    int64              `json:"storeId"`
	QrCode     string             `json:"qrCode"`
	Coin       string             `json:"coin"`
	QrCodeInfo string             `json:"qrCodeInfo"`
	Utime      parse.JsonDateTime `json:"utime"`
	Ctime      parse.JsonDateTime `json:"ctime"`
}

type ShopQrCodeQueryResult struct {
	Id         int64              `json:"id"`
	MerchantId int64              `json:"merchantId" db:"merchant_id"`
	StoreId    int64              `json:"storeId" db:"store_id"`
	QrCode     string             `json:"qrCode" db:"qr_code"`
	Coin       string             `json:"coin" db:"coin"`
	QrCodeInfo string             `json:"qrCodeInfo" db:"qr_code_info"`
	Utime      parse.JsonDateTime `json:"utime"`
	Ctime      parse.JsonDateTime `json:"ctime"`
}

type QrCodeManageResult struct {
	Id        int64  `json:"id"`
	StoreId   int64  `json:"storeId" db:"store_id"`
	StoreName string `json:"storeName" db:"store_name"`
	QrCode    string `json:"qrCode" db:"qr_code"`
	Coin      string `json:"coin" db:"coin"`
	QrCodeUrl string `json:"qrCodeUrl" db:"qr_code_url"`
}

func (store *ShopQrCode) StoreQrCodeQuery(qrCode string) (result *ShopQrCodeQueryResult, err error) {
	result = &ShopQrCodeQueryResult{}
	err = conn.DB.Get(result, "select id,merchant_id,store_id,qr_code,coin,qr_code_info,utime,ctime from t_store_qr_code where qr_code = ?", qrCode)
	if err != nil {
		clog.Errorf("查询商户店铺二维码失败,qrCode=%+v,err=%+v", store.QrCode, err)
		return
	}
	if result == nil {
		clog.Errorf("查询商户店铺二维码失败,qrCode=%+v,err=%+v", &qrCode, err)
		return nil, constant.Query_store_qr_code_error
	}
	return
}

/**
 * 生成二维码信息
 */
func (store *ShopQrCode) Add(coin *string, key string, codeUrl string, info string, mid int64, sid int64) (qrId int64, err error) {
	result, err := conn.DB.Exec("insert into t_store_qr_code (merchant_id,store_id,qr_code,qr_code_url,coin,qr_code_info,ctime,utime)values (?,?,?,?,?,?,?,?)", mid, sid, key, codeUrl, coin, info, time.Now(), time.Now())
	if err != nil {
		clog.Errorf("插入商铺二维码信息失败,coin=%+v,mid=%+v,sid=%+v,err=%+v", &coin, mid, sid, err)
		return 0, err
	}
	i, err := result.LastInsertId()
	if err != nil {
		clog.Errorf("插入商铺二维码信息失败,coin=%+v,mid=%+v,sid=%+v,err=%+v", &coin, mid, sid, err)
		return 0, err
	}
	return i, nil
}
