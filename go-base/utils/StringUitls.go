package utils


import (
	"fmt"
	"strings"
)

/**
*字符串工具类
* autor:zjcjava@163.com
* time:2018-9-10
 */


func main() {
	thistr := "abc_def_bamboo"

	fmt.Println(StrToCam(thistr))
	fmt.Println(StrToFirstUpper(thistr))

}

/**
 * 字符串转为驼峰 abc_def_bamboo -> abcDefBamboo
 */
func StrToCam(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	//fmt.Print(temp[0] + upperStr)
	return temp[0] + upperStr
}

/**
 * 字符串转为驼峰 首字母大写abc_def_bamboo -> AbcDefBamboo
 */
func StrToFirstUpper(str string) string {
	temp := StrToCam(str)
	var upperStr string
	vv := []rune(temp)
	upperStr=string(vv[0]-32)+temp[1 : len(temp)]
	return upperStr
}