package service

import (
	"coin.merchant/config"
	"coin.merchant/constant"
	"coin.merchant/models"
	"coin.merchant/repository"
	"coin.merchant/utils"
	"common/clog"
	"common/http"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net/url"
	"rdgo"
	"strconv"
	"time"
)

type MerchantQrPayCheckService struct {
}

// 预支付校验
func (s *MerchantQrPayCheckService) PrepaymentCheck(qrCodeStr string) (qrCode *models.ShopQrCodeQueryResult, merchant *models.MerchantQueryResult, err error) {
	// 根据二维码id获取店铺信息
	qrCode, err = (&StoreQrCodeService{}).StoreQrCodeQuery(qrCodeStr)
	if err != nil {
		return
	}
	// 根据商户id获取商户
	merchant, err = (&MerchantService{}).Query(qrCode.MerchantId)
	if err != nil {
		return
	}
	if merchant.Status != 1 {
		return nil, nil, constant.Merchant_service_not_start_error
	}
	if time.Now().Before(time.Time(merchant.BeginTime)) {
		return nil, nil, constant.Merchant_service_not_begin_error
	}
	if time.Now().After(time.Time(merchant.EndTime)) {
		return nil, nil, constant.Merchant_service_expire_error
	}
	return
}

/**
 * 读取二维码key包含的信息
 */
func (s *MerchantQrPayCheckService) Info(code string, uid int64) (result *models.QrCodeInfoResult, err error) {
	models := &models.QrCodeInfoResult{}
	infoResult, err := models.Info(code)
	if err != nil {
		clog.Errorf("查询二维信息异常,err=%+v", err)
		return nil, err
	}
	//获取积分价格
	price, err := repository.PriceFromRpc(infoResult.Coin)
	if err != nil {
		clog.Errorf("rpc查询 支付积分价格失败,err=%+v", err)
		return nil, err
	}
	infoResult.CoinPrice = price
	feePrice, err := repository.PriceFromRpc(infoResult.FeeUnit)
	if err != nil {
		clog.Errorf("rpc查询 手续费积分价格失败,err=%+v", err)
		return nil, err
	}
	infoResult.FeePrice = feePrice
	infoResult.ScanCode = utils.BuildQrCodeKey(uid)
	// 存放redis  扫码key
	redis := rdgo.Pool.Get()
	defer redis.Close()
	data, _ := json.Marshal(infoResult)
	if _, err := redis.Do("SET", constant.ParseKey(constant.MERCHANT_QR_PAY_CODE, infoResult.ScanCode), data, "EX", 60*60); err != nil {
		clog.Errorf("二维码信息获取失败,redis缓存扫码信息失败,err=%+v", err)
		return nil, err
	}

	return infoResult, nil
}

/**
 * 校验 scanCode
 */
func (s *MerchantQrPayCheckService) CheckScanCode(scanCode string) (success bool, bo *models.QrPayServiceBo, err error) {
	rd := rdgo.Pool.Get()
	defer rd.Close()
	data, err := redis.String(rd.Do("GET", constant.ParseKey(constant.MERCHANT_QR_PAY_CODE, scanCode)))
	if err != nil {
		clog.Errorf("二维码信息获取失败,redis缓存扫码信息失败,err=%+v", err)
		return false, nil, err
	}
	if len(data) == 0 {
		clog.Errorf("二维码信息获取失败,redis缓存scan_code信息已经过期,err=%+v", err)
		return false, nil, constant.Pay_scan_code_invalid
	}
	result := &models.QrCodeInfoResult{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		clog.Errorf("二维码信息获取失败,缓存信息解密失败,err=%+v", err)
		return false, nil, constant.Pay_scan_code_invalid
	}
	bo = &models.QrPayServiceBo{}
	bo.Coin = result.Coin
	bo.MerchantId = result.MerchantId
	bo.MerchantName = result.MerchantName
	bo.ScanCode = scanCode
	bo.CoinPrice = result.CoinPrice
	bo.FeePrice = result.FeePrice
	bo.FeeUnit = result.FeeUnit
	bo.StoreId = result.StoreId
	return true, bo, nil
}

/**
 * 校验 sign
 */
func (s *MerchantQrPayCheckService) CheckSign(param *models.QrPayServiceParam) (success bool, err error) {
	//解签
	//base64 decode
	decodeString, err := base64.StdEncoding.DecodeString(param.Sign)
	decrypt, err := RsaDecrypt(decodeString)
	clientMd5 := string(decrypt)
	if err != nil {
		clog.Errorf("签名验证失败,解签失败,err=%+v", err)
		return false, err
	}
	//通过元数据 生成md5
	serviceMd5, err := getSignCheckContent(param)
	if err != nil {
		clog.Errorf("签名验证失败,原始数据拼接异常,err=%+v", err)
		return false, err
	}
	if clientMd5 != serviceMd5 {
		clog.Errorf("签名验证失败,签名验证不匹配")
		return false, constant.Pay_sign_error
	}
	return true, nil
}

/**
 * 校验交易密码
 */
func (s *MerchantQrPayCheckService) CheckPwd(param *models.QrPayServiceParam) (err error) {
	payPwdParams := map[string]interface{}{"uid": param.Uid, "payPwd": param.PayPwd}
	url := config.Config.Servers.Account_url + fmt.Sprintf("account/users/%d/payPwd/auth", param.Uid)
	return http.Post(url, payPwdParams, nil)
}

/**
 * 参数 拼接
 */
func getSignCheckContent(param *models.QrPayServiceParam) (sign string, err error) {
	data := url.Values{}
	data.Set("uid", strconv.FormatInt(param.Uid, 10))
	data.Set("scanCode", param.ScanCode)
	data.Set("storeId", strconv.FormatInt(param.StoreId, 10))
	data.Set("money", strconv.FormatFloat(param.Money, 'f', 4, 64))
	data.Set("coin", param.Coin)
	data.Set("coinPrice", strconv.FormatFloat(param.CoinPrice, 'f', 4, 64))
	data.Set("feeUnit", param.FeeUnit)
	data.Set("feePrice", strconv.FormatFloat(param.FeePrice, 'f', 4, 64))
	data.Set("payPwd", param.PayPwd)

	s, err := url.QueryUnescape(data.Encode())
	//md5
	sum := md5.Sum([]byte(s))
	//md5 第二次
	md5Two := md5.Sum([]byte(fmt.Sprintf("%x", sum)))
	serviceMd5 := fmt.Sprintf("%x", md5Two)
	//生成签名
	return serviceMd5, nil
	//float 转string  指数
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(config.Config.Sign.PubSign))
	if block == nil {
		return nil, nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

var pri = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA58A/jfDrLNdV4Vcbd4NLcuBMSSb/hLjk/6OxX0xndogtawT1
NnmuwzcNdUJAooWjyVxc3cKsrLM2DLrOBt/uBfLKp7rxJvrS1baL7WuR1O0ekjRu
FYnMKBFkIXDaYhE7hC+MLUcID+4R5040igy9GALkzF6JGPPVKVO8kp/TfXs4zDvR
lmbJTRK/FgKe8OvIZkIzZISmBwRqjKnB2pRmskGLTVbnzxD8cW0/hVOn8Q+whlL0
dSQHfMX6c2CsvTQUYN1mkeeGOiLu0QwYSZy22DViXkQO0cgiCavl+r1ma5uNG44l
ARPuWh81iFGBQtaHKVeE6QLVTOwkF8EaiXE8ZwIDAQABAoIBADzr4MkjZ+8lvEG8
cE/+h7rvE563TbxKDojVMy9mGlyid64GY5+qZTKUKkmE3RDcKK4qRY9WOaY8hhza
joZoH14Y8QUes34XuYzMrAQBnxhmLP8qITYwPybZS4Uu8XmOJiMdjK/qWEg3wSUY
/d68coj5WcQPpeKVVpfCl3PD6Ai396jZOc7mtNlQIWmB66r99aZwBMfjUxlFHkeh
ekM8bUfQAc14qgfuGmG2/2n4dm+kiEqFTqR1lNcbDrf+s122nfOv4Izrokz98k6U
P+aU+JUJ9hwk2Bq4+nkXG6odP2cd9n8jtFCmeFJqABcea253R3fWDRPAIdZfjWuS
j3sN0UECgYEA9kKgo1N4YaX9m8wB430hy2WdL6frN4nUdw02aoQTLyYIFKUivdIQ
Mtx0SRW8hnL5ySoLUdGKBdc86/xUegSdX0iwbrlMdPrNMzyocHgIKXmfmsHOFWnm
jtyg3uzCNgOFBUV4lqRJGaQbgO4inhNz19pNM+WjE+7s6QXnxnKN2YkCgYEA8Oq3
BIi0N/umAeuBk519ipQGvhOdfQric1aPi8+ovFQ28W8qe6W+9H+7i9lVJgIc6x5/
9yJCypjB+JsbAfnDSnzWL1l3AABEclgUvoqZJDGNeJYtCjuz2WassCNOUgsMC6dV
SCkonBzeCz8WLsb4y/iYPbp8JKarTqNLWEkYGm8CgYAwFfCCE+F0x9HOozZXMm7v
5Yac8KAIdzxqhsTyZZnNYhK/3UL8Z9FL7Sozvy/R3Q+TTUdqkYzu+QlnVx0zukT0
fyAcbshUK0j4UUbet0F4v8v/jwprugMQMFqlTPvbSjKmRdt3JtszS40nTtipn0jG
hFUA5j1Cviu6kLGiWWoDaQKBgCYUg59E2G+s6D2PcyjZEPnxketDgHY+XTLr8L6h
sUMrcI/TCX0H4toUwplFXg8m8Fk9te5jTPlnEenw4mD6kKLafqR3WLb4U9lbENRZ
ZgFxj7IK0s22SCRJ9WvV+NBBDMNezL0ePFwIuBRBAYmdS8A56B6BtpO4gIVqjENF
MnkjAoGAdoX3XOhU6aP9kxvxf7ggW8klrrAuqOZJdiSJjvxvlvoiAJbniH8Elz+n
ijPM5hgPBNYNEfDi5bc+Zx0rVq9viq01/KrtdNLkKsGuMCz1S25NKNhL2uDNkU2v
VuttnnhDIov9DrgfKkwjdujea0tmjsh+Df6i4dVSOyodEBhQK4c=
-----END RSA PRIVATE KEY-----
`)

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//block, _ := pem.Decode([]byte(config.Config.Sign.PriSign))
	block, _ := pem.Decode([]byte(pri))
	if block == nil {
		return nil, nil
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
