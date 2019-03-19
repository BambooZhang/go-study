package main



/***********
JSON解析和对象话
JSON(Javascript Object Notation)是一种轻量级的数据交换语言，以文字为基础，具有自我描述性且易于让人阅读
go自带的有解析工具encoding/json


interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：
bool代表JSON booleans
float64代表JSON numbers
string代表JSON strings
nil 代表JSON null



生成JSON字符串注意事项：

针对JSON的输出，我们在定义struct tag的时候需要注意几点：
字段的tag是“-”，那么这个字段不会输出到JSON
tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中的serverName
tag中如果带有“omitempty”选项，那么如果该字段值为空，就不会输出到JSON串中
如果字段类型是bool,string,int,int64等，而tag中带有“,string”选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串


Marshal函数只有在转换成功的时候才会返回数据，在转换的过程中我们需要注意几点：
JSON对象只支持string作为key,所以要编码一个map,那么必须是map[string]T这种类型（T是Go语言中任意的类型）
Channel,complex和function是不能被编码成JSON的
嵌套的数据时不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null


参照资料
据说 json-iteator 是目前golang中对json格式数据处理最快的包(比官方json包快6倍)，好像是滴滴团队开源的
https://blog.csdn.net/luckytanggu/article/details/79795357
******/
import (
	"fmt"
	"os"
	"strings"

	"bou.ke/monkey"
)

func main() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}