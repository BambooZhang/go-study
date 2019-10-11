package routers

import (
	"coin.merchant/controllers"
	"coin.merchant/controllers/admin"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouters(app *iris.Application) {
	// admin
	mvc.Configure(app.Party("/admin/merchant"), admin.AminMerchantMvc)
	mvc.Configure(app.Party("/admin/merchant"), admin.AminMerchantFeeMvc)
	mvc.Configure(app.Party("/admin/merchant"), admin.AminMerchantStoreMvc)
	mvc.Configure(app.Party("/admin/merchant"), admin.AminMerchantCoinDetailMvc)
	mvc.Configure(app.Party("/admin/merchant"), admin.AminMerchantCoinMvc)

	//对用户接口
	mvc.Configure(app.Party("/merchant/store"), controllers.StoreMvc)
	mvc.Configure(app.Party("/merchant"), controllers.LoginMvc)
	mvc.Configure(app.Party("/merchant/coin"), controllers.CoinMvc)
	mvc.Configure(app.Party("/merchant/stat"), controllers.StatMvc)
	mvc.Configure(app.Party("/merchant"), controllers.MerchantMvc)
	mvc.Configure(app.Party("/merchant"), controllers.MerchantQrPayMvc)
}
