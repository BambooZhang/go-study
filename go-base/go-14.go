package main



/***********
 反射找到对应的方法
******/
import (
	"net/http"
	"fmt"
	"strings"
	"reflect"
)


// index handler
func Indexhandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world")
}




func main() {
	http.HandleFunc("/",Indexhandler) //添加一个路径和对应的handler
	http.Handle("/handle/", http.HandlerFunc(say)) //改行可注释，对子路径handle对应的方法
	http.ListenAndServe("127.0.0.1:9090",nil)
}



// 定义一个处理类，分别实现各个路径对应的处理方法
type Handlers struct {
}

// rest方法
func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("res"))
}

// bamboo方法
func (h *Handlers) BambooAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("bamboo")
	w.Write([]byte("bamboo"))
}

//根据/handle路径匹配其后的所有子路径，利用反射找到对应的方法进行处理
// 子路径对应的方法名规则：子路径Action
func say(w http.ResponseWriter, req *http.Request) {
	pathInfo := strings.Trim(req.URL.Path, "/")
	fmt.Println("pathInfo:", pathInfo)

	// 拆分路径
	parts := strings.Split(pathInfo, "/")
	fmt.Println("parts:", parts)

	var action= "ResAction"
	fmt.Println(strings.Join(parts, "|"))
	if len(parts) > 1 {
		fmt.Println("子路径对应的处理方法名称")
		action = strings.Title(parts[1]) + "Action" //拼接处方法名
	}
	fmt.Println("action:", action)

	// 反射找到对应的方法
	handle := &Handlers{}
	methodValue := reflect.ValueOf(handle)
	method := methodValue. MethodByName(action) //根据方法名找到该方法
	/*m, ok := methodValue.Type.MethodByName(funcName)
	if !ok {
		c.addErr(fmt.Errorf("MVC: function '%s' doesn't exist inside the '%s' controller",
			funcName, c.fullName))
		return nil
	}*/
	//if(method.Ca){
	//	fmt.Println("404:", action)
	//}

	fmt.Println("200:", action)
	r := reflect.ValueOf(req)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r}) //调用对应的方法
}