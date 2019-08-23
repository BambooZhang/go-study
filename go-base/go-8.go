package main

import (
	"fmt"
	"reflect"
)

/***
反射reflection

反射，就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），
在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型才有反射一说。

反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
反射会将匿名字段作为独立字段(匿名字段本质)
想要利用反射修改对象状态，前提是interface.data是settable，即pointer-interface
通过反射可以“动态”调用方法

Golang reflect慢主要有两个原因

涉及到内存分配以及后续的GC；

reflect实现里面有大量的枚举，也就是for循环，比如类型之类的。

 */

type User struct {
	Id   int
	Name string  "user's name" //这引号里面的就是tag
	Age  int
	Dress string `species:"gopher" color:"blue"` //注意和上面的区别
}

func (u User) Hello(name string)  string{
	fmt.Println("Hello", name, "My name is", u.Name)
	return u.Name
}







// 反射
func Info(o interface{}) {
	t := reflect.TypeOf(o)         //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t.Name()) //调用t.Name方法来获取这个类型的名称

	// 遍历出出所有类型
	if k := t.Kind(); k != reflect.Struct { //通过kind方法判断传入的类型是否是我们需要反射的类型
		fmt.Println("xx")
		return
	}


	//复合使用 TypeOf 和ValueOf  打印出所包含的字段对应的值和tag
	v := reflect.ValueOf(o)
	fmt.Println("该对象的所有属性值是:", v) // ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
	fmt.Println("Fields:")
	for i := 0; i < v.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		fmt.Printf("\n%6s:%v \t kin=%v \t tag=%s ", v.Type().Field(i).Name, v.Field(i).Interface() ,v.Type().Field(i).Type.Kind(), v.Type().Field(i).Tag)
		fmt.Printf(" \t tag中color值=%v  ",  v.Type().Field(i).Tag.Get("color"))

	}


	for i := 0; i < t.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}


//根据反射 修改实例的属性值
func Set(o interface{} ,pro string,value string) {

	t := reflect.TypeOf(o)         //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t.Name()) //调用t.Name方法来获取这个类型的名称

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("类型不符合或者该类型不可修改")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName(pro)
	if !f.IsValid() {
		fmt.Println("xiugaishibai")
	}
	if f.Kind() == reflect.String {
		f.SetString(value)
	}

}

func main() {





	fmt.Println("通过反射属性,值和获取struct的tag-----------------------------------------------")
	// 基本反射方法
	u := User{1, "Jack", 23,"蓝色校服"}
	Info(u)



	fmt.Println("通过反射修改struct中的内容-----------------------------------------------")
	//2 通过反射修改struct中的内容:修改属性Name的值为bamboo
	// reflect.ValueOf(X)只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。
	Set(&u ,"Name","bamboo")
	fmt.Println(u)


	fmt.Println("反射动态调用方法-----------------------------------------------")
	// 通过发射进行方法的调用 动态调用方法
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Name"))//根据属性名称获取值
	fmt.Println(v.FieldByName("Age"))
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("JOE")}
	rs :=mv.Call(args)
	str :=rs[0].Interface().(string) // 获取调用方法的返回值
	fmt.Println("返回值是：",str)




	fmt.Println("指针反射的区别-----------------------------------------------")
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n",typeOfCat.Name(), typeOfCat.Kind())
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())



	fmt.Println("数组-----------------------------------------------")
	int_ary := [4]int{1, 2, 3, 4}
	float32_ary := [4]float32{1.1, 2.2, 3.3, 4.4}
	float64_ary := [4]float64{1.1, 2.2, 3.3, 4.4}

	dump_interface_array(int_ary);
	dump_interface_array(float32_ary);
	dump_interface_array(float64_ary);


}





func dump_interface_array(args interface{}) {
	val := reflect.ValueOf(args)
	fmt.Println(val.Kind())
	if val.Kind() == reflect.Array {
		fmt.Println("len = ", val.Len())
		for i:=0; i<val.Len(); i++ {
			e := val.Index(i)
			switch e.Kind() {
			case reflect.Int:
				fmt.Printf("%v, ", e.Int())
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				fmt.Printf("%v, ", e.Float())
			default:
				panic(fmt.Sprintf("invalid Kind: %v", e.Kind()))
			}
		}
		fmt.Println()
	}
}
