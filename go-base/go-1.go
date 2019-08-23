package main

import (
	"fmt"
	//"reflect"
	"math"
)


/***

引用类型
引用类型和原始的基本类型恰恰相反，它的修改可以影响到任何引用到它的变量。在Go语言中，引用类型有切片、map、接口、函数类型以及chan。

引用类型之所以可以引用，是因为我们创建引用类型的变量，其实是一个标头值，标头值里包含一个指针，指向底层的数据结构，
当我们在函数中传递引用类型时，其实传递的是这个标头值的副本，它所指向的底层结构并没有被复制传递，这也是引用类型传递高效的原因


 */

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")


	var width, height int = 100, 50 // var 声明多个变量,并按次序赋值
	fmt.Println("width is", width, "height is", height)

	name, age := "naveen", 29 // 简短声明,:= 左边的变量必须都给定了右边对应的值，否则报错

	fmt.Println("my name is", name, "age is", age)

	/**
	常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	 */
	const a, b, c = 1, false, "str" //多重赋值
	fmt.Println(a, b, c)

	//类型转换
	inti := 55      //int
	floatj := 67.8    //float64
	total := inti + int(floatj) //j is converted to int GO是强类型，必须进行强转，本身没有自动类型转换功能，否则会报错
	fmt.Println(total)


	var di = 5.9/8
	fmt.Printf("a's type %T value %v \n",di, di)


	fmt.Println("数组和切片类型-------------------------------------------")
	//1.数组var identifier [size]type 注意它是定长size,去掉size则是切片的定义规则
	//2.数组之间赋值是值拷贝，不会修改原始数组中的值，这和切片不同，切片直接赋的是原始数组的头指针
	//3.数组的大小是类型的一部分。因此 [5]int = [25]int 是不同类型不能进行赋值。数组不能调整大小，因为 切片slices 的存在能解决这个问题。
	var n [10]int /* n 是一个长度为 10 的数组 */

	/* 为数组 n 初始化元素 */
	for i := 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}
	fmt.Println("n is ", n)


	aold := [...]string{"USA", "China", "India", "Germany", "France"}
	anew := aold // a copy of a is assigned to b
	anew[0] = "Singapore"
	fmt.Println("aold is ", aold)
	fmt.Println("anew is ", anew)

	//二维数组
	var ar = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	/* 输出二维数组元素 */
	for i,r :=range ar {
		for j,v :=range r {
			fmt.Printf("a[%d][%d] = %d\n", i,j, v )
		}
	}

	//切片(Slice) var identifier []type
	//类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大1被,相比数组初始化时指定了[size]而切片则没有
	//1.这是我们使用range去求一个slice的和使用数组跟这个很类似,_表示忽略返回值,如果换成其他字符这里表示返回index
	//2.数组是定长不能修改长度，切片则可以，切片实际是数组的指针引用，因此数组是值传递，而切片是地址传递
	//3.切片可以指定长度和容量make([]type, len, capacity)
	//4.切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。让我们假设我们有一个非常大的数组，我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。
	//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
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
	fmt.Println("切片的修改：可以修改原始数组的值")
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	fmt.Println("切片的容量是从创建切片的底层数组索引开始到数组结尾元素长度，长度则是从索引起始位置到结束索引位置的长度")
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))//容量是7-1=6,长度是引用结束索引值-起始索引值：3-1=2

	fmt.Println("使用 append 可以将新元素追加到切片上,新切片容量会翻了一番")
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)

	cars1 := []string{}
	cars1 = append(cars1, "Toyota")


	fmt.Println("MAP--------------------------------------------")
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

	fmt.Println(countryCapitalMap ["France"])

	/*删除元素*/ delete(countryCapitalMap, "France")


	/*使用键便利输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}




	//map嵌套遍历:原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	var aa interface{}
	aa = map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": []map[string]interface{}{{"name": "1", "group": "2"}, {"name": "3", "group": "4"}},
	}

	bb := aa.(map[string]interface{})
	for _, v := range bb["c"].([]map[string]interface{}) {
		for k1, v1 := range v {
			fmt.Println(k1, "   ", v1)
		}
	}

/*	var cc interface{}
	cc = map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": [] Books{{"Name", "Sex", "Age",6495407,32.12}},
	}
	//fmt.Println( cc["c"])
	dd := cc.(map[string]interface{})
	for _, v0 := range dd["c"].([]map[string]interface{}) {
		//for _, v0 := range cc["c"]  {
		for k2, v2 := range v0 {
		fmt.Println( k2, "   ", v2)

		}
	}*/





	fmt.Println("switch--------------------------------------------")


		num := 75
		switch { // 表达式被省略了 因此默认它为 true，true 值会和每一个 case 的求值结果进行匹配
		case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough //fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中,相当于 c,java中不使用break的作用
		case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
		case num < 200:
			fmt.Printf("%d is lesser than 200 \n", num)

		}










	fmt.Println("函数--------------------------------------------")
	//函数
	str_1, str_2 := swap("Mahesh", "Kumar")
	fmt.Println(str_1, str_2)
	////我们继续以 rectProps 函数为例，该函数计算的是面积和周长。假使我们只需要计算面积，而并不关心周长的计算结果，该怎么调用这个函数呢？这时，空白符 _ 就上场了。
	x, _ := swap("Mahesh", "Kumar") // 返回值周长被丢弃
	fmt.Printf("x is  %s \n", x)
	/*
	1.在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字(即可以在包外访问，像java的 public方法的作用)。其它包只能访问被导出的函数和变量
	2.init 函数
		所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。init 函数的形式如下：

		func init() {
		}
		init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。

		包的初始化顺序如下：
		首先初始化包级别（Package Level）的变量（即包中全局变量）
		紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用
	3.使用空白标识符（Blank Identifier）_
		有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量 import (  _ "geometry/rectangle" )//只进行初始化
		var _ = rectangle.Area 这一行屏蔽了错误
	*/


	fmt.Println("结构体--------------------------------------------")
	//结构体
	var Book1 Books        /* 声明 Book1 为 Books 类型 */

	Book1 = Books{"Name", "Sex", "Age",6495407, 32.12,2}
	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	Book1.price = 32.12

	/* 打印 Book1 信息 */
	fmt.Println( Book1)
	fmt.Printf( "Book 1 title : %s\n", Book1.title)

	fmt.Println("1.匿名字段:当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。")
	p := Person{"Naveen", 50,Address{"shenzhen","guangzhou"},Company{}}
	fmt.Println(p)
	fmt.Println("City:",p.address.city)

	fmt.Println("2.提升字段（Promoted Fields）:如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体字段一样使用")
	fmt.Println("name:",p.name)
	fmt.Println("Cname:",p.cname)

	fmt.Println("3.结构体相等性（Structs Equality）:结构体是值类型(可比较的类型,map类型就不可以)。如果它的每一个字段都是可比较的，则该结构体也是可比较，否则不可比较报错")
	c1 := Company{"shenzhen","1353535435"}
	c2:= Company{"shenzhen","1353535435"}
	if c1 == c2 {
		fmt.Println("Company 1 and Company 2 are equal")
	} else {
		fmt.Println("Company 1 and Company 2 are not equal")
	}

	fmt.Println("接口和方法调用--------------------------------------------")
	fmt.Println("1.结构体只要实现接口中的方法就可以了")
	//接口和方法调用
	Book1.print()
	fmt.Println("接口的案例")
	p1 :=Pen{"penc1", 0, 2.12, 32,0.2}
	p2 :=Pen{"penc2", 0, 2.12, 32,0.2}
	orderArry := []Order{Book1, p1, p2}
	totalExpense(orderArry)//总价格

	//注意非指针方法和指针方法的区别
	fmt.Println("针方法和指针方法的区别--------------------------------------------")
	fmt.Println("Go语言里有两种类型的接收者：值接收者和指针接收者,使用值类型接收者定义的方法，在调用的时候，使用的其实是值接收者的一个副本，所以对该值的任何操作，不会影响原来的类型变量")
	p1.modify1()
	fmt.Println(p1.category) //没有修改成功
	p1.modify2()
	fmt.Println(p1.category) //修改为：李四
	(&p1).modify2()
	fmt.Println(p1.category) //修改为：李四


	fmt.Println("函数作为参数--------------------------------------------")
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9)) //使用函数

	//函数作为参数实例2
	// 传递带一个参数函数作为参数
	funcInvokeOne(funcOne,"bamboo")
	funcInvokeMany(funcMany,"bamboo",1,3)


	fmt.Println("闭包--------------------------------------------")
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	fmt.Println("可变参数--------------------------------------------")
	print("1","2","3")


	fmt.Println("接口作为参数,调用时才确定是哪个类型的方法--------------------------------------------")
	//需要一个animal接口作为参数
	invoke(Book1)
	invoke(p1)


	fmt.Println("多个返回值--------------------------------------------")
	str,v :=getMore()
	fmt.Println(str,v)

}

/**
//函数 可以返回多个值func 函数名称(参数名称  参数类型...)  返回数据类型列表-不返回值可以省略
GO函数的定义方式
func functionname(parametername type) returntype {
    // 函数体（具体实现的功能）
}
_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。

在go中一般是值传递，除非你显示的使用指针


 */
func swap(x, y string) (string, string) {
	return y, x
}



type Address struct {
	city, state string
}
type Company struct {
	cname, tel string
}
type Person struct {
	name string//字段
	int//匿名字段
	address Address//嵌套结构
	Company//嵌套匿名接结构 能够字段提升
}


// 接口
type Order interface {
	print() ;//不带返回值
	totalPrice() float32;//总价//带返回值
}

//结构体
type Books struct {
	title string //标题
	author string //作者
	subject string //主题
	book_id int //ID
	price float32 //价格
	numb  int//数量
}


//实现接口方法
func (books Books) print() {
	 fmt.Printf("title:%s,author:%s\n",books.title,books.title)
}

//书的价格统计
func (books Books) totalPrice() float32{
	return  books.price*float32(books.numb)
}

//结构体
type Pen struct {
	category string //品牌
	pen_id int //ID
	price float32 //每只笔价格
	numb  int//只
	paf   float32 //额外税
}

func (pen Pen) print() {
	fmt.Printf("category:%s,price:%f\n",pen.category,pen.price)
}
//笔的价格统计
func (pen Pen) totalPrice() float32{
	return  pen.price*float32(pen.numb)+pen.paf
}


func (pen Pen) modify1() {
	pen.category = "李四"
}

func (pen *Pen) modify2(){
	pen.category = "李四"
}

//统计所有订单的总额
func totalExpense(s []Order) {
	expense := float32(0)
	for _, v := range s {
		expense = expense + v.totalPrice()
	}
	fmt.Println("Total Expense Per Month $%f", expense)
}


//匿名函数，可作为闭包：匿名函数的优越性在于可以直接使用函数内的变量，不必申明，外部无法使用其内部变量
func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

// 可变参数
func print (a ...interface{}){
	for _,v:=range a{
		fmt.Print(v)
	}
	fmt.Println()
}



// 接口作为参数 有泛型的感觉，只有调用的时候才知道具体是哪个类型
func invoke(a Order){
	a.print()
}

// 一个
func funcOne(s  interface{}) {
	fmt.Println("调用函数的字符串是:%s",s)
}
// 带一个参数的函数
func funcInvokeOne(handle interface{}, args interface{}) {
	handle.(func(interface{}))(args)
}

// 多个参数
func funcMany(args... interface{}) {
	fmt.Println("多个参数", args)
}
// 带多个参数的函数
func funcInvokeMany(handle  interface{}, args... interface{}) {
	handle.(func(...interface{}))(args)
}


// 多个返回值

func getMore() (string , interface{}){

	p1 :=Pen{"penc1", 0, 2.12, 32,0.2}
	return "str",p1
}