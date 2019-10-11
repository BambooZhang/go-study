package utils

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

/**
 * 生成二维码内容key
 */
func BuildQrCodeKey(id int64) (key string) {
	//随机
	head := GetRandomString(8)
	store := GetStoreString(id)
	center := GetRandomString(8)
	timeStr := GetNowTimeString()
	tail := GetRandomString(2)
	return head + store + center + timeStr + tail
}

/**
 * 取当前时间字符串 例如:2019022715
 */
func GetNowTimeString() (timeStr string) {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006010215")
	return formatTimeStr
}

/**
 *  获取店铺id 后四位 不足四位补全四位
 */
func GetStoreString(id int64) (store string) {
	str := strconv.FormatInt(id, 10)
	i := len(str)
	var buf bytes.Buffer
	if i < 4 {
		//补全到四位
		for y := 0; y < 4-i; y++ {
			buf.WriteString("0")
		}
		buf.WriteString(str)
		return buf.String()
	}
	if i > 4 {
		//截取后四位
		return str[i-4:]
	}
	return str
}

/**
 * 生成随机串
 */
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
