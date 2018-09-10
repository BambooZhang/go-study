package main


import (

	"fmt"
	"./cal"
	"./utils"

)
/*
包和可见性的使用

包是结构化代码的一种方式：每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。

如同其它一些编程语言中的类库或命名空间的概念，每个 Go 文件都属于且仅属于一个包。一个包可以由许多以 .go 为扩展名的源文件组成，因此文件名和包名一般来说都是不相同的。

你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

当你导入多个包时，导入的顺序会按照字母排序。

如果包名不是以 . 或 / 开头，如 "fmt" 或者 "container/list"，则 Go 会在全局文件进行查找；如果包名以 ./ 开头，则 Go 会在相对目录中查找；如果包名以 / 开头（在 Windows 下也可以这样使用），则会在系统的绝对路径中查找。

导入包即等同于包含了这个包的所有的代码对象。

除了符号 _，包中所有代码对象的标识符必须是唯一的，以避免名称冲突。但是相同的标识符可以在不同的包中使用，因为可以使用包名来区分它们。

包通过下面这个被编译器强制执行的规则来决定是否将自身的代码对象暴露给外部文件：

** 可见性规则 **

当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的 代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

（大写字母可以使用任何 Unicode 编码的字符，比如希腊文，不仅仅是 ASCII 码中的大写字母）。

因此，在导入一个外部包后，能够且只能够访问该包中导出的对象。
)*/


func main() {


	thisdate := "2018-09-10 14:55:06"
	fmt.Println(utils.Formate(thisdate,utils.DATE_TIME))
	fmt.Println(utils.Formate(thisdate,utils.DATE))
	fmt.Println(utils.Formate(thisdate,utils.DATE))
	fmt.Println(utils.Formate(thisdate,utils.TIME))
	fmt.Println(utils.Formate(thisdate,utils.SHORT_TIME))
	fmt.Println(utils.Formate(thisdate,utils.NEW_DATE_TIME))
	fmt.Println(utils.Formate(thisdate,utils.NEW_TIME))
	fmt.Println(utils.NewDateStamp()) //当前时间戳
	fmt.Println(utils.NewDatetime()) //当前时间


	//包和可见性
	result := cal.Add(1,2)
	fmt.Println(result)

}

/*
func utils.utils.Formate(datastr string,layout string) string {
	timeformatdate, _ := time.Parse(datetime, datastr)
	convdate := timeformatdate.Format(layout)
	return convdate
}*/
