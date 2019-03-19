package main


/***********
国际化处理
golang.org/x/text 这个包，看它提供了什么用来格式化和本地化的工具
Go 中大多数消息（message）要么用 fmt 要么通过 template 包处理。
golang.org/x/text 包含多层子包，提供了很多的工具和函数，并且用 fmt 风格的 API 来格式化 字符串

1.硬编码方式
2.自动加载消息
一直以来，大多数本地化的框架都会把每个语言的译文分别存于文件里，这些文件会被动态加载。你可以把 这些文件交给翻译的人，在他们搞定后，你再把译文合并到你的应用中。
为了协助这个过程，Go 作者们开发了一个命令行小工具叫 gotext
安装：go get -u golang.org/x/text/cmd/gotext




中文简写的一些扩展知识

zh-CHS 是单纯的简体中文。
zh-CHT 是单纯的繁体中文。

zh-Hans和zh-CHS相同相对应。
zh-Hant和zh-CHT相同相对应。

以上时zh-CHS/zh-Hans 和 zh-CHT/zh-Hant的关系。

然后是
zh-CN 简体中文，中华人民共和国
zh-HK 繁体中文，香港特别行政区
zh-MO 繁体中文，澳门特别行政区
zh-SG 繁体中文，新加坡-
zh-SG 简体中文，新加坡
zh-TW 繁体中文，台湾



参照资料
https://www.colabug.com/3411106.html
******/

import (
	"golang.org/x/text/message"
	"golang.org/x/text/language"
	"fmt"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/currency"
	"golang.org/x/text/message/catalog"
	"io/ioutil"
	"encoding/json"
)
func init() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")

	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")

	message.SetString(language.Chinese, "HOW_ARE_U", "%s 你好吗?")
	message.SetString(language.AmericanEnglish, "HOW_ARE_U", "%s How are you?")


	//根据参数处理不同的返回语句
	message.Set(language.English, "APP_COUNT",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))


	// 以上代码和以下代码都是硬编码方式
	for _, e := range msaArry {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}

}

func main() {
	// 中文版
	p := message.NewPrinter(language.Chinese)
	p.Printf("%s went to %s.", "彼得", "英格兰")
	fmt.Println()
	p.Printf("%s has been stolen.", "宝石")
	fmt.Println()
	p.Printf("HOW_ARE_U", "竹子")
	fmt.Println()


	// 英文版本
	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")
	fmt.Println()
	p.Printf("HOW_ARE_U", "bamboo")
	fmt.Println()



	fmt.Println("placehold中的条件判断-------------------")
	// 条件判断
	p.Printf("APP_COUNT", 1)
	fmt.Println()
	p.Printf("APP_COUNT", 2)
	p.Println()


	fmt.Println("货币单位-------------------")
	// 货币单位
	p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))//符号  美元货币格式化
	fmt.Println()
	p.Printf("%d", currency.NarrowSymbol(currency.JPY.Amount(1.6)))//窄符号  日元货币格式化
	fmt.Println()
	p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(12.255)))//国际符号代码 欧元格式化
	fmt.Println()



	fmt.Println("国家语言简写格式-------------------")
	//调用硬编码中的消息 国际化
	p = message.NewPrinter(language.English)
	p.Printf("HELLO_WORLD","bamboo")
	p.Println()
	p.Printf("TASK_REM", 2)
	p.Println()



	//语言类型构建
	zh, _ := language.ParseBase("zh") // 语言
	CN, _ := language.ParseRegion("CN") // 地区
	zhLngTag, _ := language.Compose(zh, CN)
	fmt.Println(zhLngTag) // 打印 zh-CN
	fmt.Println(language.Chinese)// 打印中文缩写
	fmt.Println(language.SimplifiedChinese)// 打印中文缩写
	fmt.Println(language.TraditionalChinese)// 打印中文缩写
	fmt.Println(language.AmericanEnglish)// 打印中文缩写


	fmt.Println("从配置文件中读取配置json解析并加载到系统中-------------------")
	InitConfig("go-base/locales/el/messages_zh.json")//加载配置信息json
	InitConfig("go-base/locales/el/messages_en.json")//加载配置信息json
	p = message.NewPrinter(language.SimplifiedChinese)//设置语言类型
	p.Printf("HELLO_1", "Peter")
	fmt.Println()

	p.Printf("VISITOR", "Peter","用户管理系统接口")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish)//设置语言类型
	p.Printf("VISITOR", "Peter","UER MANAGE SYSTEM API")
	fmt.Println()


	msg := Message{"en", "HELLO_WORLD", "%s Hello World"}
	p.Printf("CONFIG", msg) //传一个对象值进去
	fmt.Println()


}


type langMsg struct {
	tag, key string
	msg      interface{}
}


//手工 硬编码方式
var msaArry = [...]langMsg{
	{"en", "HELLO_WORLD", "%s Hello World"},
	{"zh", "HELLO_WORLD", "%s 你好世界"},
	{"en", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"zh", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "剩余一项任务！",
		"=2", "剩余两项任务！",
		"other", "剩余 [1]d 项任务！",
	)},
}



//


// 后面的jsoN字符串隐射在生成json字符串时需要,默认
type Message struct {
	Id string `json:"id"`
	Message string `json:"message,omitempty"`
	Translation   string `json:"translation,omitempty"`
}

type I18n struct {
	Language string `json:"language"`
	Messages []Message `json:"messages"`
}


// ioutil读写文件，依赖 io/ioutil 主要侧重文件和临时文件的读取和写入，对文件夹的操作较少
func  ReadI18nJson(file string) string {
	b, err := ioutil.ReadFile(file)
	Check(err)
	str := string(b)
	return str

}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func InitConfig(jsonPath string) {

	var i18n I18n
	str := ReadI18nJson(jsonPath)
	json.Unmarshal([]byte(str), &i18n)
	fmt.Println(i18n.Language)

	msaArry := i18n.Messages
	tag := language.MustParse(i18n.Language)
	// 以上代码和以下代码都是硬编码方式
	for _, e := range msaArry {

		fmt.Println(e.Id+"\t"+e.Translation)
		message.SetString(tag, e.Id, e.Translation)

	}
}