package models

import (
	"coin/conn"
	"common/clog"
)

// 确认支付业务参数
type QrPayServiceParam struct {
	Uid       int64   `json:"uid"`       // 支付用户
	StoreId   int64   `json:"storeId"`   // 支付用户
	ScanCode  string  `json:"scanCode"`  // 扫码唯一凭证
	Money     float64 `json:"money"`     // 支付金额
	Amount    float64 `json:"amount"`    // 支付积分数
	Coin      string  `json:"coin"`      // 支付金额
	CoinPrice float64 `json:"CoinPrice"` // 支付金额
	FeeUnit   string  `json:"feeUnit"`   // 用户手续费积分数
	FeePrice  float64 `json:"feePrice"`  // 用户手续费积分数
	PayPwd    string  `json:"payPwd"`    // 支付密码
	Pf        string  `json:"pf"`        // 来源
	PhoneVer  string  `json:"phoneVer"`  // 手机版本
	AppVer    string  `json:"appVer"`    // app版本
	Sign      string  `json:"sign"`      // 签名
}

// 支付业务对象
type QrPayServiceBo struct {
	ScanCode     string  `json:"scanCode"`
	StoreId      int64   `json:"storeId"`
	MerchantId   int64   `json:"merchantId"`
	MerchantName string  `json:"merchantName"`
	Coin         string  `json:"coin"`
	CoinPrice    float64 `json:"coinPrice"`
	FeeUnit      string  `json:"feeUnit"`
	FeePrice     float64 `json:"feePrice"`
}

/**
 * 二维码信息 读取结果
 */
type QrCodeInfoResult struct {
	ScanCode     string  `json:"scanCode"`
	StoreId      int64   `json:"storeId" db:"store_id"`
	MerchantId   int64   `json:"merchantId" db:"merchant_id"`
	MerchantName string  `json:"merchantName" db:"merchant_name"`
	HeadImage    string  `json:"headImage" db:"head_image"`
	Coin         string  `json:"coin"`
	CoinPrice    float64 `json:"coinPrice"`
	FeeUnit      string  `json:"feeUnit" db:"fee_coin"`
	FeePrice     float64 `json:"feePrice"`
}

/**
 * 支付结果
 */
type QrCodePayResult struct {
	Status       string `json:"status"`       //支付结果   success 成功  fail 失败
	OrderNo      int64  `json:"orderNo" `     //订单号
	MerchantId   int64  `json:"merchantId"`   //收款方id
	MerchantName string `json:"merchantName"` //收款方姓名  需要脱敏
	Coin         string `json:"coin"`
	PayType      string `json:"payType"` //支付方式
}

/**
 * 通过key 查询二维码的支付信息
 */
func (m QrCodeInfoResult) Info(key string) (result *QrCodeInfoResult, err error) {
	result = &QrCodeInfoResult{}
	err = conn.DB.Get(result, "SELECT t.store_id,t.merchant_id,tm.merchant_name,t.coin,tm.head_image,tmc.fee_coin "+
		"FROM t_store_qr_code t LEFT JOIN t_merchant tm on t.merchant_id=tm.id LEFT JOIN t_merchant_fee_config tmc "+
		"ON t.merchant_id=tmc.merchant_id AND t.coin=tmc.coin WHERE t.qr_code =? ", key)
	if err != nil {
		clog.Errorf("查询二维信息异常,err=%+v")
		return nil, err
	}
	return
}
