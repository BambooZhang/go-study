package repository

import (
	"coin.merchant/constant"
	"common/clog"
	"context"
	"price/rpc/client"
)

func PriceFromRpc(coin string) (price float64, err error) {
	result := &rpc.RpcPriceDaySeviceResult{}
	err = rpc.RpcPriceDaySevice().Call(context.Background(), "SimplePriceByCoin", coin, result)
	if err != nil || result.Today <= 0 {
		clog.Errorf("获取积分价格失败,coin=%s,err=%+v", coin, err)
		err = constant.Merchant_get_price_error
		return
	}
	price = result.Today
	return
}
