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
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, "My name is", u.Name)
}



// 反射
func Info(o interface{}) {
	t := reflect.TypeOf(o)         //反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	fmt.Println("Type:", t.Name()) //调用t.Name方法来获取这个类型的名称

	if k := t.Kind(); k != reflect.Struct { //通过kind方法判断传入的类型是否是我们需要反射的类型
		fmt.Println("xx")
		return
	}


	v := reflect.ValueOf(o) //打印出所包含的字段
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := t.Field(i)               //通过这个i作为它的索引，从0开始来取得它的字段
		val := v.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)
	}


	for i := 0; i < t.NumMethod(); i++ { //这里同样通过t.NumMethod来获取它拥有的方法的数量，来决定循环的次数
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}


//根据反射 修改实例的属性值
func Set(o interface{} ,pro string,value string) {
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
	// 基本反射方法
	u := User{1, "Jack", 23}
	//Info(u)


	//2 通过反射修改struct中的内容:修改属性Name的值为bamboo
	// reflect.ValueOf(X)只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。
	//Set(&u ,"Name","bamboo")
	fmt.Println(u)


	// 通过发射进行方法的调用 动态调用方法
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Name"))//根据属性名称获取值
	fmt.Println(v.FieldByName("Age"))
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("JOE")}
	mv.Call(args)

}