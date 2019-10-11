package controllers

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"rdgo"
)

type merchantQrPayController struct {
	Ctx iris.Context
}

func MerchantQrPayMvc(mvc *mvc.Application) {
	mvc.Handle(new(merchantQrPayController))
}

func (m *merchantQrPayController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/qr/pay/{uid:int64}/info/{key:string}", "Info") //  预支付(扫码)
	b.Handle("POST", "/qr/pay", "Confirm")                           //  确定支付
}

/**
/merchant/qr/pay/prepayment
预支付(扫码)
*/
func (m *merchantQrPayController) Info(uid int64, key string) {
	// 校验
	if len(key) != 32 {
		clog.Errorf("二维码信息查询异常,二维码key,长度错误,uid=%+v", uid)
		m.Ctx.JSON(code.ReturnError(code.Params_input_error))
		return
	}
	service := &service.MerchantQrPayCheckService{}
	result, err := service.Info(key, uid)
	if err != nil {
		clog.Errorf("二维码信息查询异常,uid=%+v,err=%+v", uid, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("info", result))
}

/**
/merchant/qr/pay/confirm
确定支付
*/
func (m *merchantQrPayController) Confirm() {
	params := &models.QrPayServiceParam{}
	if err := m.Ctx.ReadJSON(&params); err != nil {
		clog.Errorf("扫码支付错误,参数解析错误")
		m.Ctx.JSON(code.ReturnError(constant.Pay_param_error))
		return
	}
	data, err := json.Marshal(params)
	if err != nil {
		clog.Errorf("扫码支付错误,参数解析错误")
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	bytes, err := json.Marshal(data)
	clog.Infof("扫码支付,参数信息params=%+v", string(bytes))
	checkService := &service.MerchantQrPayCheckService{}
	//checkCode唯一校验
	rd := rdgo.Pool.Get()
	defer rd.Do("del", params.ScanCode)
	defer rd.Close()
	status, err := redis.Int64(rd.Do("setnx", params.ScanCode, "lock"))
	if err != nil || status == 0 {
		clog.Errorf("发起了重复支付,uid=%+v", params.Uid)
		m.Ctx.JSON(code.ReturnError(constant.Pay_duplicate_payment_error))
		return
	}
	//校验 scanCode和 ts
	scan, bo, err := checkService.CheckScanCode(params.ScanCode)
	if !scan || err != nil {
		clog.Errorf("scanCode校验失败,失败原因=%+v", err)
		m.Ctx.JSON(code.ReturnError(constant.Pay_scan_code_error))
		return
	}
	// 校验 Sign
	sign, err := checkService.CheckSign(params)
	if !sign || err != nil {
		clog.Errorf("sign校验失败,失败原因=%+v", err)
		m.Ctx.JSON(code.ReturnError(constant.Pay_sign_error))
		return
	}
	// 校验 交易密码是否正确
	err = checkService.CheckPwd(params)
	if err != nil {
		clog.Errorf("密码校验失败,失败原因=%+v", err)
		m.Ctx.JSON(code.ReturnError(constant.Pay_pwd_error))
		return
	}
	payService := &service.MerchantQrPayService{}
	// 校验 发送支付请求 TODO
	result, err := payService.Pay(params, bo)
	if err != nil {
		clog.Errorf("支付失败 err=%+v", err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	//返回支付结果
	m.Ctx.JSON(code.DefaultSuccess().Add("result", result))
}
