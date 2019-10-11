package models

import (
	"bytes"
	"coin.merchant/conn"
	"coin.merchant/constant"
	"common/clog"
	"common/parse"
	"fmt"
	"time"
)

/**
 * 商户积分实体对象
 */
type MerchantCoin struct {
	Id          int64              `json:"id"`
	MerchantId  int64              `json:"merchantId"`
	Uid         int64              `json:"uid"`
	Coin        string             `json:"coin"`
	Amount      float64            `json:"amount"`
	LockAmount  float64            `json:"lockAmount"`
	TotalAmount float64            `json:"totalAmount"`
	WalletAddr  string             `json:"walletAddr"`
	Utime       parse.JsonDateTime `json:"utime"`
	Ctime       parse.JsonDateTime `json:"ctime"`
}

/**
商户积分钱包地址-新增
*/
func (c *MerchantCoin) Add() (id int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("insert into t_merchant_coin  (merchant_id ,uid ,coin ,amount ,lock_amount ,total_amount ,wallet_addr ,utime ,ctime)")
	buffer.WriteString("values(?,?,?,?,?,?,?,?,?)")
	result, err := conn.DB.Exec(buffer.String(), c.MerchantId, c.Uid, c.Coin, c.Amount, c.LockAmount, c.TotalAmount, c.WalletAddr, time.Now(), time.Now())
	if err != nil {
		clog.Errorf("新增商户积分钱包地址信息失败,rank=%+v,err=%+v", c, err)
		return
	}
	id, err = result.LastInsertId()
	clog.Infof("新增商户信息成功, merchantId=%d", c.Id)
	return
}

/**
查询商户积分地址是否存在
*/
func (c *MerchantCoin) CheckExist(merchantId int64, coin string) (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant_coin where 1=1 ")
	if merchantId > 0 {
		buffer.WriteString(fmt.Sprintf(" and merchant_id ='%d'", merchantId))
	}
	if len(coin) > 0 {
		buffer.WriteString(fmt.Sprintf(" and coin ='%s'", coin))
	}
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询商户积分钱包地址信息失败:err=%+v", err)
		return 0, constant.Query_merchant_coin_error
	}
	return
}
