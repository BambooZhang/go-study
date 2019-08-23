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
	"encoding/json"
	"fmt"
)


// 后面的jsoN字符串隐射在生成json字符串时需要,默认
type Server struct {
	ServerName string//默认会和字段名称相同
	ServerIP   string `json:"serverIp"` // 注意这里的IP和Ip的大小写我做了不同的隐射处理
	Inet    string `json:"inet,omitempty"`//omitempty:如果ServerIP为空，则不输出到JSON中
	Meme    string `json:"-"`//“ - ”，则该字段总是被省略不会转成JSON字段和值
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1","Meme":"Babmoo_002SD"},
            {"serverName":"Beijing_VPN","serverIP":"127.0.0.2","Meme":"Babmoo_002SD"}]}`

	json.Unmarshal([]byte(str), &s) // json字符串解析成对象
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)



	fmt.Println("-------------------interface{}可以用来存储任意数据类型的对象")

	// interface{}可以用来存储任意数据类型的对象
	// 通过下面的示例可以看到，通过interface{}与type assert的配合，我们就可以解析未知结构的JSON函数了。
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents": [ "Gomez", "Moticia" ]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}

	m := f.(map[string]interface{})

	for k, v := range m {

		fmt.Print(fmt.Println(m[k]))
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)

		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("-------------------MAP成JSON-----------")
	json1, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json err: ", err)
	}

	fmt.Println(string(json1))

	fmt.Println("-------------------生成JSON-----------")
	//生成JSON 若我们要输出JSON数据串，可通过Marshal函数来处理
	var serList Serverslice
	serList.Servers = append(serList.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "117.0.0.1",Inet:"192.168.0.1",Meme:"Babmoo_002SD"})
	serList.Servers = append(serList.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2",Meme:"Babmoo_002SA"})

	json, err := json.Marshal(serList)
	if err != nil {
		fmt.Println("json err: ", err)
	}

	fmt.Println(string(json))

}