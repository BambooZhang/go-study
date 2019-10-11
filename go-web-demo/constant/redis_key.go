package constant

import "fmt"

const (
	BLOCK_HEIGHT     = "block:height:confirms" //map
	USER_SHARE_SUM   = "user:share:sum:%d"     //用户分享总数,只存放1天的有效期
	LOGIN_PHONE_CODE = "login:phone:code:%s"   //登陆短信验证码

	MERCHANT_QR_PAY_CODE = "m:q:p:c:%s" // 当前扫码的支付凭证

	REDIS_ORDER_NO_POOL     = "order:no:pool"     //订单池
	REDIS_ORDER_NO_POOL_MAX = "order:no:pool:max" //订单池最大值
)

const (
	LOCK_USER_COIN_DETAIL  = "l:u:c:d:%d" //user_coin_detail的所，fid
	LOCK_USER_COINS_WALLET = "l:u:c:w:%d" //user_coin_detail的所，fid
)

func ParseKey(path string, params ...interface{}) string {
	if params == nil || len(params) == 0 {
		return path
	}
	return fmt.Sprintf(path, params...)
}
