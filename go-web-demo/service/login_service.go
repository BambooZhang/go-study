package service

import (
	"coin.merchant/config"
	"coin.merchant/constant"
	"coin.merchant/models"
	"common/clog"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type LoginService struct {
}

func (service LoginService) SendCode(phone string) (result bool, err error) {
	merchant := &models.Merchant{}
	//验证手机号码是不是商户号
	merchantTwo, err := merchant.CheckPhoneValid(phone)
	if err != nil {
		return false, err
	}
	if merchantTwo == 0 {
		clog.Errorf("验证手机号是否是商户失败,该手机号不是商户", err)
		return false, constant.Merchant_phone_not_existent
	}
	//发送短信验证码
	client := &http.Client{}
	sprintf := strconv.Itoa((rand.Intn(900000) + 100000))
	//生成要访问的url
	url := config.Config.Servers.Notice_url + "notice/sms/send?phone=" + phone + "&code=" + sprintf
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	response, _ := client.Do(reqest)
	//处理返回结果
	body, err := ioutil.ReadAll(response.Body)
	if response == nil {
		return false, constant.Sms_code_error
	}
	aaa := &models.Response{}
	err = json.Unmarshal([]byte(string(body)), &aaa)
	if aaa.Code != 0 {
		return false, constant.Sms_code_error
	}
	return true, nil
}

func (service LoginService) Login(login *models.Login) (status bool, user *models.User, err error) {
	//生成要访问的url
	url := config.Config.Servers.Notice_url + "notice/sms/verify?phone=" + login.Phone + "&code=" + login.Code
	//提交请求
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, nil)
	response, _ := client.Do(reqest)
	//处理返回结果
	body, err := ioutil.ReadAll(response.Body)
	if response == nil {
		return false, nil, constant.Sms_code_error
	}
	aaa := &models.Response{}
	err = json.Unmarshal([]byte(string(body)), &aaa)
	if aaa.Code != 0 {
		return false, nil, constant.Sms_code_error
	}
	//获取商户信息
	merchant := &models.Merchant{}
	user, err = merchant.PhoneQuery(login.Phone)
	if err != nil {
		clog.Errorf("登陆失败,通过手机号码查询 商户信息为空 err=%+v,phone=%+v", err, login.Phone)
		return false, nil, err
	}
	return true, user, nil
}
