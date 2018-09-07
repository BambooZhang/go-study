package main

import (
	"fmt"
)


func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	/**
	常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	 */
	const a, b, c = 1, false, "str" //多重赋值
	fmt.Println(a, b, c)

	//数组var identifier [size]type 注意它是定长size
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i,j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

	//二维数组
	var ar = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	/* 输出二维数组元素 */
	for i,r :=range ar {
		for j,v :=range r {
			fmt.Printf("a[%d][%d] = %d\n", i,j, v )
		}
	}

	//切片(Slice) var identifier []type
	//类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大,相比数组初始化时指定了[size]而切片则没有
	//1.这是我们使用range去求一个slice的和使用数组跟这个很类似,_表示忽略返回值,如果换成其他字符这里表示返回index
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	//golang之下划线(_)之语义说明三种用法（忽略返回值，用在变量(特别是接口断言）,用在import package只做初始化init ）:参考资料https://blog.csdn.net/qq_21816375/article/details/77971697
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} //定义一个非定长数组并负责初始化
	for i,v :=range balance{ //遍历输出
		fmt.Printf("Element[%d] = %f\n", i, v )
	}
	//以上都是在定义的时候初始化了，如果不给默认值则需要使用slice1 := make([]type, len, capacity)创建,len() 和 cap() 函数分别对应这里的len和capacity
	//len() 和 cap() 函数
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。





	/*2.创建集合MAP
	 range：关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值
	*/
	var countryCapitalMap map[string]string //定义集合
	countryCapitalMap = make(map[string]string)//要使用它必须初始化否则默认nail map会报错

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"


	/*删除元素*/ delete(countryCapitalMap, "France")


	/*使用键便利输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}


	//函数
	str_1, str_2 := swap("Mahesh", "Kumar")
	fmt.Println(str_1, str_2)



	//结构体
	var Book1 Books        /* 声明 Book1 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	Book1.price = 32.12

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
	fmt.Printf( "Book 1 book_id : %f\n", Book1.price)

	//接口和方法调用
	fmt.Println(Book1.toString())
	Book1.print()

}


//函数 可以返回多个值func 函数名称(参数列表) （返回数据类型列表-不返回值可以省略）
func swap(x, y string) (string, string) {
	return y, x
}


// 接口
type BookInf interface {
	toString() string;//带返回值
	print() ;//不带返回值
}

//结构体
type Books struct {
	title string //标题
	author string //作者
	subject string //主题
	book_id int //ID
	price float32 //价格
}

//实现接口方法
func (books Books) toString() string{
	return "title:"+books.title+",author:"+books.author
}

func (books Books) print() {
	 fmt.Printf("title:%s,author:%s",books.title,books.title)
}

