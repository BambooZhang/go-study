package models

import (
	"coin.merchant/conn"
	"common/clog"
)

type MerchantStat struct {
	TotalAmount float64 `json:"totalAmount" db:"totalAmount"`
	TotalBuy    int64   `json:"totalBuy" db:"totalBuy"`
	TotalStore  int64   `json:"totalStore" db:"totalStore"`
	TotalUser   int64   `json:"totalUser" db:"totalUser"`
}

func (c *MerchantStat) MerchantStat(mid int64) (stat *MerchantStat, err error) {
	stat = &MerchantStat{}
	err = conn.DB.Get(stat, "select SUM(receive_amount) AS totalAmount,count(id) AS totalBuy,"+
		"count(distinct store_id) AS totalStore,count(distinct uid) AS totalUser "+
		"from t_merchant_coin_detail where merchant_id=? and  income=0 and  `type`=0 group by merchant_id ", mid)
	if err != nil {
		clog.Errorf("统计商户收入信息异常,mid=%+v,err=%+v", mid, err)
		return stat, err
	}
	return stat, nil
}
