package service

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"common/clog"
	"github.com/garyburd/redigo/redis"
	"rdgo"
	"strconv"
)

const (
	YMDHMS = "20060102150405"
)

type MerchantQrPayService struct {
}

/**
 * 支付
 */
func (s *MerchantQrPayService) Pay(param *models.QrPayServiceParam, bo *models.QrPayServiceBo) (result *models.QrCodePayResult, err error) {
	//生成订单号
	//orderNo, err := GeneratorOrderNo()
	if err != nil {
		clog.Errorf("支付接口生成订单号失败,err=%+v", err)
		return nil, constant.Query_store_error
	}
	//写入 商户交易记录表 用户发起扣款   商户还未收款

	//todo 发起 支付方预扣款

	//支付成功 更新状态 商户交易记录表 用户扣款成功  商户还未收款

	//支付失败 更新状态 商户交易记录表 用户扣款失败  商户还未收款   订单已经结束

	//发送mq 通知coin 通知同一笔订单已经取消

	//封装结果
	result = &models.QrCodePayResult{}
	result.MerchantId = bo.MerchantId
	result.MerchantName = bo.MerchantName
	result.Coin = bo.Coin
	return result, nil
}

/**
 * 生成订单号
 */
func GeneratorOrderNo() (orderNo string, err error) {
	//
	rd := rdgo.Pool.Get()
	defer rd.Close()
	orderNo, err = redis.String(rd.Do("spop", constant.REDIS_ORDER_NO_POOL))
	if err != nil {
		clog.Errorf("redis获取订单号 失败,err=%+v", err)
		return "", err
	}
	if orderNo == "" || len(orderNo) == 0 {
		clog.Errorf("redis获取订单号 失败,单订池空了,err=%+v", err)
		//新增订单
		FillPool()
	}
	return constant.MERCHANT_ORDER_HEAD, nil
}

/**
 * 新增10W个订单
 */
func FillPool() (err error) {
	rd := rdgo.Pool.Get()
	defer rd.Close()
	count, err := redis.Int64(rd.Do("scard", constant.REDIS_ORDER_NO_POOL))
	if err != nil {
		clog.Errorf("redis获取订单池失败,,err=%+v", err)
		return err
	}
	// 订单池还有单订号
	if count > 50000 {
		return nil
	}
	max, err := redis.Int64(rd.Do("GET", constant.REDIS_ORDER_NO_POOL_MAX))
	if err != nil {
		clog.Errorf("redis获取订单池最大订单号失败,err=%+v", err)
		return err
	}
	size := 100000
	orderNos := [100000]string{}
	for i := 0; i < size; i++ {
		itoa := strconv.Itoa(i)
		i2, err := strconv.ParseInt(itoa, 10, 64)
		if err != nil {
			return err
		}
		orderNos[i] = strconv.FormatInt(max+i2, 10)
	}

	return
}
