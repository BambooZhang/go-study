package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/***
io的使用
 */

func main() {


	//ioutilFile()
	osFile()

}

// ioutil读写文件，依赖 io/ioutil 主要侧重文件和临时文件的读取和写入，对文件夹的操作较少
func ioutilFile()  {
	b, err := ioutil.ReadFile("README.md")
	check(err)
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)

	// write file:if this file not exit  will create
	//d1 := []byte("hello\ngo\n")
	err1 := ioutil.WriteFile("tst.txt", b, 0644) //write byte
	check(err1)
	err2 := ioutil.WriteFile("tst.txt", []byte(str+"\nzjcjava@163.com"), 0644) //write string
	check(err2)
}


//使用os进行读写文件，功能强大，包括文件和文件夹的读写权限等操作
func osFile()  {
	//创建一个文件夹：
	os.Mkdir("test_java", 0777)
	//创建多级文件夹：
	os.MkdirAll("test_go/go1/go2", 0777)

	//删除文件夹
	err := os.Remove("test_java")
	check(err)
	os.RemoveAll("test_go")


	fl, err := os.Open("README.md")
	check(err)
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}







func check(e error) {
	if e != nil {
		panic(e)
	}
}