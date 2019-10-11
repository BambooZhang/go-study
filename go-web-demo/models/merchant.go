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
商户实体，对应数据库:zwy_merchant表t_merchant
*/
type Merchant struct {
	Id            int64              `json:"id"`
	Uid           int64              `json:"uid"`
	MerchantName  string             `json:"merchantName"`
	HeadImage     string             `json:"headImage"`
	Address       string             `json:"address"`
	Phone         string             `json:"phone"`
	Coin          string             `json:"coin"`
	CorporateName string             `json:"corporateName"`
	BeginTime     parse.JsonDateTime `json:"beginTime"`
	EndTime       parse.JsonDateTime `json:"endTime"`
	ValidityMonth int64              `json:"validityMonth"`
	Introduce     string             `json:"introduce"`
	Status        int8               `json:"status"`
	Operator      string             `json:"operator"`
	Utime         parse.JsonDateTime `json:"utime"`
	Ctime         parse.JsonDateTime `json:"ctime"`
}

type MerchantQueryResult struct {
	Id            int64              `json:"id"`
	Uid           int64              `json:"uid"`
	MerchantName  string             `json:"merchantName" db:"merchant_name"`
	HeadImage     string             `json:"headImage" db:"head_image"`
	Address       string             `json:"address"`
	Phone         string             `json:"phone"`
	Coin          string             `json:"coin"`
	CorporateName string             `json:"corporateName" db:"corporate_name"`
	BeginTime     parse.JsonDateTime `json:"beginTime" db:"begin_time"`
	EndTime       parse.JsonDateTime `json:"endTime" db:"end_time"`
	ValidityMonth int64              `json:"validityMonth" db:"validity_month"`
	Introduce     string             `json:"introduce"`
	Status        int8               `json:"status"`
	Operator      string             `json:"operator"`
	Utime         parse.JsonDateTime `json:"utime"`
	Ctime         parse.JsonDateTime `json:"ctime"`
}

type MerchantUpdateStatusUpdateParams struct {
	Id       int64  `json:"id"`
	Status   int8   `json:"status"`
	Operator string `json:"operator"`
}

/**
新增商户信息
*/
func (c *Merchant) Add() (id int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("insert into t_merchant  (uid,merchant_name,head_image,address,phone,coin,corporate_name,begin_time,end_time,validity_month,introduce,status,operator,utime,ctime)")
	buffer.WriteString("values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	result, err := conn.DB.Exec(buffer.String(), c.Uid, c.MerchantName, c.HeadImage, c.Address, c.Phone, c.Coin, c.CorporateName, time.Time(c.BeginTime), time.Time(c.EndTime), c.ValidityMonth, c.Introduce, c.Status, c.Operator, time.Now(), time.Now())
	if err != nil {
		clog.Errorf("新增商户信息失败,rank=%+v,err=%+v", c, err)
		return
	}
	id, err = result.LastInsertId()
	clog.Infof("新增商户信息成功, merchantId=%d", c.Id)
	return
}

/**
查询用户名是否存在
*/
func (result *Merchant) CheckMerchantName(merchantName string) (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant where 1=1 ")
	if len(merchantName) > 0 {
		buffer.WriteString(fmt.Sprintf(" and merchant_name ='%s'", merchantName))
	}
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询merchant信息失败:err=%+v", err)
		return 0, constant.Query_merchant_error
	}
	return
}

/**
查询列表总数
*/
func (result *Merchant) MerchantsTotal() (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_merchant where 1=1 ")
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询merchant信息失败:err=%+v", err)
		return 0, constant.Query_merchant_error
	}
	return
}

/**
查询列表
*/
func (c *Merchant) Merchants(page *models.Page) (results []*MerchantQueryResult, err error) {
	results = []*MerchantQueryResult{}
	err = conn.DB.Select(&results, "select id,uid,address,corporate_name,introduce,merchant_name,head_image,begin_time,end_time,validity_month,phone,coin,`status`,operator,utime,ctime from t_merchant where 1=1 order by id asc limit ?,?", page.Start, page.Size)
	if err != nil {
		clog.Errorf("查询merchant信息失败:err=%+v", err)
		return nil, constant.Query_merchant_error
	}
	return
}

/**
修改状态
*/
func (m *MerchantUpdateStatusUpdateParams) UpdateStatus() (err error) {
	sql := "update t_merchant set status=?,operator=?,utime=? where id=?"
	_, err = conn.DB.Exec(sql, m.Status, m.Operator, time.Now(), m.Id)
	if err != nil {
		clog.Errorf("更新merchant.status失败:err=%+v", err)
	}
	return err
}

/**
 *  判断手机号 是否是商户的
 */
func (c *Merchant) CheckPhoneValid(phone string) (total int, err error) {
	err = conn.DB.Get(&total, "select count(1) from t_merchant where phone=? and status=1 ", phone)
	if err != nil {
		clog.Errorf("验证手机号是否是商户失败:err=%+v", err)
		return 0, err
	}
	return
}

/**
 * 通过商户id 查询商户信息
 */
func (c *Merchant) Query(mid int64) (merchant *MerchantQueryResult, err error) {
	merchant = &MerchantQueryResult{}
	err = conn.DB.Get(merchant, "select id,merchant_name,head_image,address,phone,corporate_name,begin_time,end_time,introduce from t_merchant where id=? and status=1 ", mid)
	if err != nil {
		clog.Errorf("查询商户信息异常 :mid=%+v,err=%+v", mid, err)
		return nil, err
	}
	if merchant == nil {
		clog.Errorf("查询商户信息异常 未能查询到商户信息 :mid=%+v,err=%+v", mid, err)
		return nil, constant.Query_merchant_error
	}
	return merchant, nil
}

func (c *Merchant) Modify() (bool, error) {
	result, err := conn.DB.Exec("update t_merchant set introduce=?,head_image=?,utime=? where id=? ", c.Introduce, c.HeadImage, time.Now(), c.Id)
	if err != nil {
		clog.Errorf("修改商户信息异常  :merchantId=%+v,err=%+v", c.Id, err)
		return false, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		clog.Errorf("修改商户信息异常:store=%+v,err=%+v", c.Id, err)
		return false, err
	}
	if count == 0 {
		clog.Infof("修改商户信息异常:store=%+v,err=%+v", c.Id, err)
		return false, constant.Modify_merchant_error
	}
	return true, nil
}

/**
 *  通过手机号码查询 商户信息
 */
func (c *Merchant) PhoneQuery(phone string) (user *User, err error) {
	user = &User{}
	err = conn.DB.Get(user, "select id as  mid,uid,phone from t_merchant where phone=? and status=1 ", phone)
	if err != nil {
		clog.Infof("通过手机号码查询 商户id失败:phone=%+v,err=%+v", phone, err)
		return nil, err
	}
	return user, nil

}

/**
 * 获取商户下所有的二维码
 */
func (c *Merchant) QrCodes(mid int64) (results []*QrCodeManageResult, err error) {
	results = []*QrCodeManageResult{}
	err = conn.DB.Select(&results, "SELECT t.id,t.store_id,ts.store_name,t.qr_code,coin,t.qr_code_url FROM t_store_qr_code t LEFT JOIN t_store ts on t.store_id =ts.id WHERE t.merchant_id=? ", mid)
	if err != nil {
		clog.Infof("查询商户二维码信息失败 :mid=%+v,err=%+v", mid, err)
		return nil, err
	}
	return results, nil
}
