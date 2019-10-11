package models

import (
	"bytes"
	"coin.merchant/conn"
	"coin.merchant/constant"
	"common/clog"
	"common/models"
	"common/parse"
	"fmt"
)

type MerchantCoinDetail struct {
	Id            int64              `json:"id"`
	Fid           int64              `json:"fid"`
	MerchantId    int64              `json:"merchantId" db:"merchant_id"`
	StoreId       int64              `json:"storeId" db:"store_id"`
	Uid           int64              `json:"uid"`
	Uphone        int64              `json:"uPhone" db:"u_phone"`
	Type          int                `json:"type"`
	Coin          string             `json:"coin"`
	PayAmount     float64            `json:"payAmount" db:"pay_amount"`
	ReceiveAmount float64            `json:"receiveAmount" db:"receive_amount"`
	Income        int                `json:"income"`
	Ctime         parse.JsonDateTime `json:"ctime"`
}

type MerchantCoinDetailQueryResult struct {
	Fid           int64              `json:"fid"`
	StoreId       int64              `json:"storeId" db:"store_id"`
	UserPhone     string             `json:"userPhone" db:"user_phone"`
	Coin          string             `json:"coin"`
	PayAmount     float64            `json:"payAmount" db:"pay_amount"`
	ReceiveAmount float64            `json:"receiveAmount" db:"receive_amount"`
	Ctime         parse.JsonDateTime `json:"ctime"`
}

/**
积分支付订单列表
*/
type AdminMerchantCoinDetailQueryResult struct {
	Id            int64              `json:"id"`
	OrderNo       string             `json:"orderNo" db:"order_no"`
	PayTime       parse.JsonDateTime `json:"payTime" db:"pay_time"`
	MerchantName  string             `json:"merchantName" db:"merchant_name"`
	StoreId       int64              `json:"storeId" db:"store_id"`
	Coin          string             `json:"coin"`
	Money         float64            `json:"money"`
	UnitPrice     float64            `json:"unitPrice" db:"unit_price"`
	PayAmount     float64            `json:"payAmount" db:"pay_amount"`
	ReceiveAmount float64            `json:"receiveAmount" db:"receive_amount"`
	MFee          float64            `json:"mFee" db:"m_fee"`
	UFee          float64            `json:"uFee" db:"u_fee"`
	UserName      string             `json:"userName" db:"user_name"`
	UserPhone     string             `json:"userPhone" db:"user_phone"`
	Type          int8               `json:"type"`
}

type AdminMerchantCoinDetailQueryParams struct {
	BeginTime    string `json:"beginTime"`
	EndTime      string `json:"endTime"`
	MerchantName string `json:"merchantName"`
	Coin         string `json:"coin"`
	UserPhone    string `json:"userPhone"`
	Type         string `json:"type"`
}

type AdminCoinPayTotalQueryParams struct {
	Coin         string `json:"coin"`
	PayTimeMonth string `json:"payTimeMonth"`
	FormatDate   string `json:"formatDate"` // 0按日查询 1按月查询
}

type AdminCoinPayStatQueryResult struct {
	PayTime       parse.JsonDateTime `json:"payTime" db:"pay_time"`
	Coin          string             `json:"coin"`
	MerchantCount int64              `json:"merchantCount" db:"merchant_count"`
	StoreCount    int64              `json:"storeCount" db:"store_count"`
	PayCount      int64              `json:"payCount" db:"pay_count"`
	PayAmount     float64            `json:"payAmount" db:"pay_amount"`
	Fee           float64            `json:"fee" db:"fee"`
	UserCount     int64              `json:"userCount" db:"user_count"`
}

/**
 *  查询指定商户  所有的积分收入明细
 */
func (result MerchantCoinDetailQueryResult) AllIncomes(mid int64, page *models.Page) (coinDetailQueryResult []*MerchantCoinDetailQueryResult, err error) {
	err = conn.DB.Select(&coinDetailQueryResult, "select fid,store_id,user_phone,coin,pay_amount,receive_amount,ctime from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0 order by id desc limit ?,? ", mid, page.Start, page.Size)
	if err != nil {
		clog.Errorf("商户查询全部收入失败,merchant_id=%+v,err=%+v", mid, err)
		return
	}
	return coinDetailQueryResult, nil
}

func (result MerchantCoinDetailQueryResult) Incomes(mid int64, storeId string, page *models.Page) (coinDetailQueryResult []*MerchantCoinDetailQueryResult, err error) {
	if storeId == "" {
		err = conn.DB.Select(&coinDetailQueryResult, "select fid,store_id,user_phone,coin,pay_amount,receive_amount,ctime from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0 order by id desc limit ?,? ", mid, page.Start, page.Size)
		if err != nil {
			clog.Errorf("商户查询全部收入失败,merchant_id=%+v,err=%+v", mid, err)
			return
		}
		return coinDetailQueryResult, nil
	}
	err = conn.DB.Select(&coinDetailQueryResult, "select fid,store_id,user_phone,coin,pay_amount,receive_amount,ctime from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0 and store_id=? order by id desc limit ?,? ", mid, storeId, page.Start, page.Size)
	if err != nil {
		clog.Errorf("商户指定商铺收入,merchant_id=%+v,storeId=%+v,err=%+v", mid, storeId, err)
		return nil, err
	}
	return coinDetailQueryResult, nil
}

/**
 * 查询数据总数
 */
func (result MerchantCoinDetailQueryResult) Total(mid int64, storeId string) (total int64, err error) {
	if storeId == "" {
		err = conn.DB.Get(&total, "select count(id) from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0", mid)
		if err != nil {
			clog.Errorf("分页查询商户收入总条数异常,merchant_id=%+v,err=%+v", mid, err)
			return 0, err
		}
		return
	}
	err = conn.DB.Get(&total, "select count(id) from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0 and store_id=? ", mid, storeId)
	if err != nil {
		clog.Errorf("分页查询商户指定商铺收入总条数异常,merchant_id=%+v,storeId=%+v,err=%+v", mid, storeId, err)
		return 0, err
	}
	return
}
func (detail MerchantCoinDetail) Total(mid int64) (total float64, err error) {
	err = conn.DB.Get(&total, "select sum(receive_amount) from t_merchant_coin_detail where merchant_id=? and income=0 and `type`=0", mid)
	if err != nil {
		clog.Errorf("分页查询商户指定商铺收入总条数异常,merchant_id=%+v,err=%+v", mid, err)
		return 0, err
	}
	return
}

func (params *AdminMerchantCoinDetailQueryParams) MerchantCoinDetailsTotal() (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant_coin_detail mcd inner join t_merchant m on mcd.merchant_id = m.id where 1=1 ")
	if len(params.BeginTime) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.pay_time >='%s'", params.BeginTime))
	}
	if len(params.EndTime) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.pay_time <='%s'", params.EndTime))
	}
	if len(params.MerchantName) > 0 {
		buffer.WriteString(fmt.Sprintf(" and m.merchant_name ='%s'", params.MerchantName))
	}
	if len(params.Coin) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.coin ='%s'", params.Coin))
	}
	if len(params.UserPhone) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.user_phone ='%s'", params.UserPhone))
	}
	if len(params.Type) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.type ='%s'", params.Type))
	}
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询积分订单支付失败:err=%+v", err)
		return 0, constant.Merchant_coin_detail_error
	}
	return
}

func (params *AdminMerchantCoinDetailQueryParams) MerchantCoinDetails(page *models.Page) (merchantCoinDetails []*AdminMerchantCoinDetailQueryResult, err error) {
	var buffer bytes.Buffer
	var sqlAppend = []interface{}{}
	buffer.WriteString("select mcd.id,mcd.order_no,mcd.pay_time,m.merchant_name,mcd.store_id,mcd.coin,mcd.money,mcd.unit_price,mcd.pay_amount,mcd.m_fee,mcd.u_fee,mcd.user_phone,mcd.user_name,mcd.type from t_merchant_coin_detail mcd inner join t_merchant m on mcd.merchant_id = m.id where 1=1 ")
	if len(params.BeginTime) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.pay_time >= ?"))
		sqlAppend = append(sqlAppend, params.BeginTime)
	}
	if len(params.EndTime) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.pay_time <= ?"))
		sqlAppend = append(sqlAppend, params.EndTime)
	}
	if len(params.MerchantName) > 0 {
		buffer.WriteString(fmt.Sprintf(" and m.merchant_name = ?"))
		sqlAppend = append(sqlAppend, params.MerchantName)
	}
	if len(params.Coin) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.coin = ?"))
		sqlAppend = append(sqlAppend, params.Coin)
	}
	if len(params.UserPhone) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.user_phone = ?"))
		sqlAppend = append(sqlAppend, params.UserPhone)
	}
	if len(params.Type) > 0 {
		buffer.WriteString(fmt.Sprintf(" and mcd.type = ?"))
		sqlAppend = append(sqlAppend, params.Type)
	}
	buffer.WriteString(" order by mcd.id desc ")
	if page != nil {
		buffer.WriteString(" limit ?,? ")
		sqlAppend = append(sqlAppend, page.Start, page.Size)
	}
	merchantCoinDetails = []*AdminMerchantCoinDetailQueryResult{}
	err = conn.DB.Select(&merchantCoinDetails, buffer.String(), sqlAppend...)
	if err != nil {
		clog.Errorf("查询积分订单支付失败:params=%+v,err=%+v", *params, err)
	}
	return
}

func (params *AdminCoinPayTotalQueryParams) CoinPayTotal() (total int64, err error) {
	var sql bytes.Buffer
	sql.WriteString("SELECT COUNT(0) FROM (SELECT count(0) FROM t_merchant_coin_detail WHERE 1 = 1")
	if len(params.Coin) > 0 {
		sql.WriteString(fmt.Sprintf(" and coin ='%s'", params.Coin))
	}
	if len(params.PayTimeMonth) > 0 {
		sql.WriteString(" AND DATE_FORMAT(pay_time, '%Y-%m') = ")
		sql.WriteString(fmt.Sprintf("'%s'", params.PayTimeMonth))
	}
	if params.FormatDate == "0" {
		sql.WriteString(" GROUP BY coin,DATE_FORMAT(pay_time, '%Y-%m-%d')) temp")
	}
	if params.FormatDate == "1" {
		sql.WriteString(" GROUP BY coin,DATE_FORMAT(pay_time, '%Y-%m')) temp")
	}
	err = conn.DB.Get(&total, sql.String())
	if err != nil {
		clog.Errorf("查询积分支付统计出错:params=%+v,err=%+v", *params, err)
		return
	}
	return
}

func (params *AdminCoinPayTotalQueryParams) CoinPayStat(page *models.Page) (coinPayStat []*AdminCoinPayStatQueryResult, err error) {
	var sql bytes.Buffer
	sql.WriteString("SELECT pay_time,coin,count(DISTINCT merchant_id) merchant_count,count( DISTINCT store_id) AS store_count,count(uid) pay_count,sum(pay_amount) AS pay_amount,sum(m_fee+u_fee) AS fee,count(DISTINCT uid) AS user_count FROM t_merchant_coin_detail WHERE 1=1 ")
	if len(params.Coin) > 0 {
		sql.WriteString(fmt.Sprintf(" and coin ='%s'", params.Coin))
	}
	if len(params.PayTimeMonth) > 0 {
		sql.WriteString(" AND DATE_FORMAT(pay_time, '%Y-%m') = ")
		sql.WriteString(fmt.Sprintf("'%s'", params.PayTimeMonth))
	}
	if params.FormatDate == "0" {
		sql.WriteString(" GROUP BY coin,DATE_FORMAT(pay_time, '%Y-%m-%d')")
	}
	if params.FormatDate == "1" {
		sql.WriteString(" GROUP BY coin,DATE_FORMAT(pay_time, '%Y-%m')")
	}
	err = conn.DB.Select(&coinPayStat, sql.String())
	if err != nil {
		clog.Errorf("查询积分支付统计出错:params=%+v,err=%+v", *params, err)
		return
	}
	return
}
