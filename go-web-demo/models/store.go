package models

import (
	"bytes"
	"coin.merchant/conn"
	"coin.merchant/constant"
	"common/clog"
	"common/models"
	"common/parse"
	"time"
)

/**
 * 商铺 实体
 */
type Store struct {
	Id         int64              `json:"id"`
	MerchantId int64              `json:"merchantId"`
	StoreName  string             `json:"storeName"`
	Name       string             `json:"name"`
	Phone      string             `json:"phone"`
	Address    string             `json:"address"`
	Status     int8               `json:"status"`
	Delete     int8               `json:"delete"`
	Operator   string             `json:"operator"`
	Utime      parse.JsonDateTime `json:"utime"`
	Ctime      parse.JsonDateTime `json:"ctime"`
}

/**
 * 商铺 列表缩略信息 实体
 */
type StoreQueryResult struct {
	Id        int64  `json:"id"`
	StoreName string `json:"storeName" db:"store_name"`
	Address   string `json:"address"`
	status    int    `json:"status"`
}

/**
 * 商铺 详细信息 实体
 */
type StoreQueryDetailResult struct {
	Id           int64  `json:"id"`
	StoreName    string `json:"storeName" db:"store_name"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	MerchantId   int64  `json:"merchantId" db:"merchant_id"`
	MerchantName string `json:"merchantName" db:"merchant_name"`
	QrCode       string `json:"qrCode" db:"qr_code"`
}

type StoreQueryDetailResultA struct {
	StoreName string `json:"storeName" db:"store_name"`
}

type StoreAddParams struct {
	MerchantId int64  `json:"merchantId"`
	StoreName  string `json:"storeName"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

type StoreModifyParams struct {
	Id         int64  `json:"id"`
	MerchantId int64  `json:"merchantId"`
	StoreName  string `json:"storeName"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

/**
 *新增 商铺
 */
func (c *StoreAddParams) Add() (id int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("insert into t_store  (merchant_id,store_name,name,phone,address,status,`delete`,utime,ctime)")
	buffer.WriteString("values(?,?,?,?,?,?,?,?,?)")
	result, err := conn.DB.Exec(buffer.String(), c.MerchantId, c.StoreName, c.Name, c.Phone, c.Address, 0, 0, time.Now(), time.Now())
	if err != nil {
		clog.Errorf("新增商铺失败,rank=%+v,err=%+v", c, err)
		return
	}
	id, err = result.LastInsertId()
	clog.Infof("新增商铺成功, merchantId=%d", id)
	return
}

func (c *StoreModifyParams) Modify() (err error) {
	var buffer bytes.Buffer
	buffer.WriteString("update  t_store  set store_name=?,name=?,phone=?,address=?,utime=? ")
	buffer.WriteString("where merchant_id=? and id=?")
	result, err := conn.DB.Exec(buffer.String(), c.StoreName, c.Name, c.Phone, c.Address, time.Now(), c.MerchantId, c.Id)
	if err != nil {
		clog.Errorf("商铺信息修改失败,rank=%+v,err=%+v", c, err)
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		clog.Errorf("商铺信息修改:rank=%+v,err=%+v", *c, err)
		return err
	}
	if count == 0 {
		clog.Infof("商铺信息修改:rank=%+v,err=%+v", *c, err)
		return constant.Modify_store_error
	}
	return
}

func (store *Store) Stores(mid int64) (results []*StoreQueryResult, err error) {
	results = []*StoreQueryResult{}
	err = conn.DB.Select(&results, "select id,store_name ,address,status from t_store where merchant_id=? and `delete`=0 and status=0", mid)
	if err != nil {
		clog.Errorf("查询商户商铺列表失败,merchantId=%+v,err=%+v", mid, err)
		return nil, constant.Query_store_list_error
	}
	return
}

func (store *Store) DeleteStore() (err error) {
	result, err := conn.DB.Exec("update  t_store  set `delete`=1 where merchant_id=? and  id =?", store.MerchantId, store.Id)
	if err != nil {
		clog.Errorf("商铺信息删除失败:store=%+v,err=%+v", store, err)
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		clog.Errorf("商铺信息删除失败:store=%+v,err=%+v", store, err)
		return err
	}
	if count == 0 {
		clog.Infof("商铺信息删除失败:store=%+v,err=%+v", store, err)
		return constant.Modify_store_error
	}
	return
}

func (store *Store) Query(merchantId int64, sid int64) (result *StoreQueryDetailResult, err error) {
	result = &StoreQueryDetailResult{}
	err = conn.DB.Get(result, "select s.id,s.store_name,s.`name`,s.phone,s.address,m.id AS merchant_id,m.merchant_name from t_store s LEFT JOIN t_merchant m ON s.merchant_id = m.id where s.merchant_id=? and s.id=? and `delete`=0 ", merchantId, sid)
	if err != nil {
		clog.Errorf("查询商户商铺列表失败,storeId=%+v,err=%+v", store.Id, err)
		return
	}
	if result == nil {
		clog.Errorf("查询单个商铺详情失败:store=%+v,err=%+v", &sid, err)
		return nil, constant.Query_store_error
	}
	return
}

func (store *Store) DisableStore() (err error) {
	sql := "update t_store set `status` = ?,operator = ?,utime = ? where id = ? and merchant_id = ?"
	_, err = conn.DB.Exec(sql, store.Status, store.Operator, time.Now(), store.Id, store.MerchantId)
	if err != nil {
		clog.Errorf("商户店铺管理-禁用失败:err=%+v", err)
	}
	return err
}

/**
商户店铺管理-列表总计
*/
func (result *Store) StoresTotal() (total int64, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("select count(1) from t_store s inner join t_merchant m on s.merchant_id = m.id inner join t_store_qr_code sqc on s.id = sqc.store_id where 1=1 and s.`delete` = 0")
	err = conn.DB.Get(&total, buffer.String())
	if err != nil {
		clog.Errorf("查询商户店铺信息失败:err=%+v", err)
		return 0, constant.Query_store_error
	}
	return
}

var adminStoresSQL = "select s.id,s.merchant_id,s.store_name,s.address,s.`name`,s.phone,m.merchant_name,sqc.qr_code from t_store s inner join t_merchant m on s.merchant_id = m.id inner join t_store_qr_code sqc on s.id = sqc.store_id"

func (c *Store) AdminStores(page *models.Page) (results []*StoreQueryDetailResult, err error) {
	results = []*StoreQueryDetailResult{}
	err = conn.DB.Select(&results, adminStoresSQL+" where 1=1 and s.`delete` = 0 order by s.id asc limit ?,?", page.Start, page.Size)
	if err != nil {
		clog.Errorf("查询商户店铺信息失败:err=%+v", err)
		return nil, constant.Query_store_error
	}
	return
}
