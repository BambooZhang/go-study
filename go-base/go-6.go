package main

import "fmt"
import "html/template"
import (
	"os"
	//"net/http"
	"strings"
)

/***
模板的使用
 */

type Person struct {
	Id      int
	Name    string
	Country string
	Age  int
}

func main() {

	//1.字符串模板
	// stringPase()

	//2.file模板
	filePase()
}

func stringPase()  {
	liumiaocn := Person{Id: 1001, Name: "liumiaocn", Country: "China"}

	fmt.Println("liumiaocn = ", liumiaocn)

	tmpl := template.New("tmpl1")
	tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	tmpl.Execute(os.Stdout, liumiaocn)
}

func filePase()  {

	funcMaps := template.FuncMap{"Func": strFirstToUpper}   //把定义的函数实例

	tmpl , err :=  template.New("test.html").Funcs(funcMaps).ParseFiles("test.html")//注册要使用的函数
	if err != nil {
		//...
	}
	//tmpl  :=  template.Must(template.ParseFiles("test.html"))//注册要使用的函数
	list := []Person{
		{1001,"Learn Go", "china",20},
		{1002,"Read Go Web Examples", "hongKong",25},
		{1003,"Create a web app in Go", "huBei",18},
	}


	tmpl.Execute(os.Stdout, struct{ Lists []Person }{list})
	/*fmt.Println("http server have stared :  http://localhost:8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Lists []Person }{list})
	})
	http.ListenAndServe(":8081", nil)*/



}


/**
 * 字符串转为驼峰 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func strToCam(str string) string {
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
 * 字符串转为驼峰 首字母大写
 */
func strFirstToUpper(str string) string {
	temp := strToCam(str)
	var upperStr string
	vv := []rune(temp)
	upperStr=string(vv[0]-32)+temp[1 : len(temp)]
	return upperStr
}