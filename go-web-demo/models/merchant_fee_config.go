package models

import (
	"bytes"
	"coin.merchant/conn"
	"coin.merchant/constant"
	"common/clog"
	"common/models"
	"common/parse"
	"fmt"
	"time"
)

/**
商户费率配置实体，对应数据库:zwy_merchant表t_merchant_fee_config
*/
type MerchantFeeConfig struct {
	Id         int64              `json:"id"`
	MerchantId int64              `json:"merchantId"`
	Coin       string             `json:"coin"`
	FeeType    int8               `json:"feeType"`
	FeeCoin    string             `json:"feeCoin"`
	FeeRate    float64            `json:"feeRate"`
	Utime      parse.JsonDateTime `json:"utime"`
	Ctime      parse.JsonDateTime `json:"ctime"`
}

type MerchantFeesQueryResult struct {
	Id           int64              `json:"id"`
	MerchantId   int64              `json:"merchantId" db:"merchant_id"`
	MerchantName string             `json:"merchantName" db:"merchant_name"`
	Coin         string             `json:"coin"`
	FeeType      int8               `json:"feeType" db:"fee_type"`
	FeeCoin      string             `json:"feeCoin" db:"fee_coin"`
	FeeRate      float64            `json:"feeRate" db:"fee_rate"`
	Utime        parse.JsonDateTime `json:"utime"`
	Ctime        parse.JsonDateTime `json:"ctime"`
}

type MerchantFeeUpdateUpdateParams struct {
	Id      int64   `json:"id"`
	FeeType int8    `json:"feeType"`
	FeeRate float64 `json:"feeRate"`
	FeeCoin string  `json:"feeCoin"`
}

/**
新增商户积分费率配置
*/
func (m *MerchantFeeConfig) Add() (id int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("insert into t_merchant_fee_config (merchant_id ,coin ,fee_type ,fee_coin, fee_rate ,utime ,ctime)")
	buffer.WriteString("values(?,?,?,?,?,?,?)")
	result, err := conn.DB.Exec(buffer.String(), m.MerchantId, m.Coin, m.FeeType, m.FeeCoin, m.FeeRate, time.Now(), time.Now())
	if err != nil {
		clog.Errorf("新增商户积分费率配置成功,rank=%+v,err=%+v", m, err)
		return
	}
	id, err = result.LastInsertId()
	clog.Infof("新增商户积分费率配置成功, id=%d", m.Id)
	return
}

/**
查询商户积分地址是否存在
*/
func (c *MerchantFeeConfig) CheckExist(merchantId int64, coin string) (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant_fee_config where 1=1 ")
	if merchantId > 0 {
		buffer.WriteString(fmt.Sprintf(" and merchant_id ='%d'", merchantId))
	}
	if len(coin) > 0 {
		buffer.WriteString(fmt.Sprintf(" and coin ='%s'", coin))
	}
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询商户积分费率配置信息失败:err=%+v", err)
		return 0, constant.Query_merchant_fee_config_error
	}
	return
}

/**
查询列表总数
*/
func (result *MerchantFeeConfig) MerchantFeesTotal() (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant_fee_config where 1=1 ")
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询merchantFee信息失败:err=%+v", err)
		return 0, constant.Query_merchant_fee_config_error
	}
	return
}

/**
查询列表
*/
func (c *MerchantFeeConfig) MerchantFees(page *models.Page) (results []*MerchantFeesQueryResult, err error) {
	results = []*MerchantFeesQueryResult{}
	err = conn.DB.Select(&results, "select mfc.id ,mfc.merchant_id ,mfc.coin ,mfc.fee_type ,mfc.fee_coin,mfc.fee_rate ,mfc.utime ,mfc.ctime,m.merchant_name from t_merchant_fee_config mfc LEFT JOIN t_merchant m ON mfc.merchant_id = m.id where 1=1 order by mfc.id asc limit ?,?", page.Start, page.Size)
	if err != nil {
		clog.Errorf("查询merchantFee信息失败:err=%+v", err)
		return nil, constant.Query_merchant_fee_config_error
	}
	return
}

/**
费率积分配置-修改
*/
func (m *MerchantFeeUpdateUpdateParams) UpdateMerchantFee() (err error) {
	sql := "update t_merchant_fee_config set fee_rate =?,fee_type = ?,fee_coin = ?,utime =? where id = ?"
	_, err = conn.DB.Exec(sql, m.FeeRate, m.FeeType, m.FeeCoin, time.Now(), m.Id)
	if err != nil {
		clog.Errorf("更新merchantFee失败:err=%+v", err)
	}
	return err
}

/**
根据id获取积分费率配置详情
*/
func (m *MerchantFeeConfig) MerchantFeeId(id int64) (result *MerchantFeesQueryResult, err error) {
	result = &MerchantFeesQueryResult{}
	err = conn.DB.Get(result, "select mfc.id ,mfc.merchant_id ,mfc.coin ,mfc.fee_type ,mfc.fee_coin, mfc.fee_rate ,mfc.utime ,mfc.ctime,m.merchant_name from t_merchant_fee_config mfc LEFT JOIN t_merchant m ON mfc.merchant_id = m.id where mfc.id = ? ", id)
	if err != nil {
		clog.Errorf("根据Id查询merchantFee信息失败:id=%d,err=%+v", id, err)
		return nil, constant.Query_merchant_fee_config_error
	}
	return
}

/**
根据merchantId和积分类型获取积分费率配置详情
*/
func (m *MerchantFeeConfig) MerchantFeeQuery(merchantId int64, coin string) (result *MerchantFeesQueryResult, err error) {
	result = &MerchantFeesQueryResult{}
	err = conn.DB.Get(result, "select id,merchant_id,coin,fee_type,fee_rate from t_merchant_fee_config where merchant_id = ? and coin = ? ", merchantId, coin)
	if err != nil {
		clog.Errorf("根据Id查询merchantFee信息失败:merchantId=%d,err=%+v", merchantId, err)
		return nil, constant.Merchant_fee_not_exist_error
	}
	return
}

/**
 * 获取商户的积分 配置信息
 */
func (m *MerchantFeeConfig) CoinConfig(mid int64) (coins []*string, err error) {

	err = conn.DB.Select(&coins, "select coin from t_merchant_fee_config where  merchant_id =?", mid)
	if err != nil {
		clog.Errorf("查询商户 支持的所有积分:merchantId=%d,err=%+v", mid, err)
		return nil, err
	}
	return

}
