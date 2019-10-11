package constant

import "common/code"

var (
	Query_merchant_error             = &code.MyError{"6001", "查询商户信息失败"}
	Modify_store_error               = &code.MyError{"6002", "修改商铺信息失败"}
	Query_store_list_error           = &code.MyError{"6003", "查询商铺列表失败"}
	Merchant_name_exist_error        = &code.MyError{"6004", "商户名不可重复"}
	Query_merchant_fee_config_error  = &code.MyError{"6005", "查询商户费率配置信息失败"}
	Query_store_error                = &code.MyError{"6006", "查询商铺信息失败"}
	Merchant_login_phone_error       = &code.MyError{"6007", "商户登陆手机号码错误"}
	Merchant_phone_not_existent      = &code.MyError{"6008", "商户手机号码不存在"}
	Merchant_stat_error              = &code.MyError{"6009", "商户收入信息统计异常"}
	Sms_code_error                   = &code.MyError{"6010", "发送短信验证码失败"}
	Params_error                     = &code.MyError{"6011", "参数传递错误"}
	Modify_merchant_error            = &code.MyError{"6012", "修改商户信息失败"}
	Query_store_qr_code_error        = &code.MyError{"6013", "无效的二维码"}
	Login_error                      = &code.MyError{"6014", "登陆信息有误"}
	Merchant_service_not_begin_error = &code.MyError{"6015", "商户服务时间未到,支付失败"}
	Merchant_service_expire_error    = &code.MyError{"6016", "商户服务时间已到期,支付失败"}
	Merchant_service_not_start_error = &code.MyError{"6017", "商户服务未开启,支付失败"}
	Merchant_fee_config_error        = &code.MyError{"6018", "商家手续费配置异常"}
	Merchant_get_price_error         = &code.MyError{"6019", "获取实时价格出现异常"}
	Merchant_fee_not_exist_error     = &code.MyError{"6020", "商户不存在"}
	Merchant_coin_detail_error       = &code.MyError{"6021", "查询为空"}
	Merchant_coin_exist_error        = &code.MyError{"6022", "商户积分地址不可重复添加"}
	Query_merchant_coin_error        = &code.MyError{"6023", "查询商户积分钱包地址信息失败"}
	Merchant_coin_fee_exist_error    = &code.MyError{"6024", "商户积分费率不可重复添加"}
	Params_illegal_error             = &code.MyError{"6025", "非法操作"}
)

var (
	Pay_param_error                = &code.MyError{"6401", "支付参数传递错误"}
	Pay_scan_code_error            = &code.MyError{"6402", "scanCode校验失败"}
	Pay_sign_error                 = &code.MyError{"6403", "sign校验失败"}
	Pay_pwd_error                  = &code.MyError{"6404", "支付密码错误"}
	Pay_balance_insufficient_error = &code.MyError{"6405", "余额不足"}
	Pay_unknown_error              = &code.MyError{"6406", "未知异常"}
	Pay_account_error              = &code.MyError{"6407", "账户异常"}
	Pay_scan_code_invalid          = &code.MyError{"6408", "scanCode已经过期"}
	Pay_order_no_error             = &code.MyError{"6409", "生成支付订单号失败"}
	Pay_duplicate_payment_error    = &code.MyError{"6410", "发起了重复支付"}
)
