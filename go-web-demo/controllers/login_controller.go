package controllers

import (
	"coin.merchant/constant"
	"coin.merchant/models"
	"coin.merchant/service"
	"common/clog"
	"common/code"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

type loginController struct {
	Ctx iris.Context
}

func LoginMvc(mvc *mvc.Application) {
	mvc.Handle(new(loginController))
}

func (m *loginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/code/{phone:string}", "SendCode")
	b.Handle("POST", "/login", "Login")
}

/**
 * 发送短信验证码
 */
func (m *loginController) SendCode(phone string) {
	//验证手机号码正确性
	if len(phone) != 11 {
		clog.Errorf("商户登陆发送验证码失败,手机号长度不正确:phone=%+v", phone)
		m.Ctx.JSON(code.ReturnError(constant.Merchant_login_phone_error))
		return
	}
	int64, err := strconv.ParseInt(phone, 10, 64)
	if err != nil {
		clog.Errorf("商户登陆发送验证码失败,手机号纯在字符:phone=%+v", int64)
		m.Ctx.JSON(code.ReturnError(constant.Merchant_login_phone_error))
		return
	}

	service := &service.LoginService{}
	status, err := service.SendCode(phone)
	if err != nil {
		clog.Errorf("商户登陆 发送验证码失败:phone=%+v,err=%+v", phone, err)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	if !status {
		clog.Errorf("商户登陆 发送验证码失败:phone=%+v", phone)
		m.Ctx.JSON(code.ReturnError(err))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess())
}

/**
 * 验证登陆
 */
func (m *loginController) Login() {
	login := &models.Login{}
	if err := m.Ctx.ReadJSON(&login); err != nil {
		clog.Infof("修改商铺失败:store=%+v,err=%+v", &login, err)
		return
	}
	service := &service.LoginService{}
	result, user, err := service.Login(login)
	if err != nil && !result {
		clog.Errorf("商户登陆失败 :phone=%+v", login.Phone)
		m.Ctx.JSON(code.ReturnError(constant.Login_error))
		return
	}
	m.Ctx.JSON(code.DefaultSuccess().Add("user", user))
}
