package main


import (
	"os"
	"text/template"
	"strconv"
)

type x struct {
	A姓名, B级别, C性别 string
}

const M = `{{range $k,$v := .}}{{$k|Func|print}}{{$v.A姓名}}   // "|"作用相当于管道,用来传值
{{end}}`

func main() {
	var di = []x{{"曦晨", "1", "男"}, {"晨曦", "2", "女"}}
	Func := template.FuncMap{"Func": ce}   //把定义的函数实例
	t := template.New("")
	t.Funcs(Func)    //注册要使用的函数
	t.Parse(M)
	t.Execute(os.Stdout, di)
}

func ce(i int) string {
	return "姓名：" +strconv.Itoa(i)

}