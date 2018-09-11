package main

import (
	"github.com/Luxurioust/excelize"
	"fmt"
	"os"
	"strconv"
	"reflect"
)

/***
https://raw.githubusercontent.com/Luxurioust/excelize/master/excelize.png
Excelize 是 Golang 编写的一个用来操作 Office Excel 文档类库，基于微软的 Office OpenXML 标准。可以使用它来读取、写入 XLSX 文件。
相比较其他的开源类库，Excelize 支持写入原本带有图片(表)的文档，还支持向 Excel 中插入图片，并且在保存后不会丢失图表样式。

项目主页 github.com/Luxurioust/excelize
安装
go get github.com/Luxurioust/excelize
 */

//结构体
type Person struct {
	Name string //姓名
	Sex string //性别
	Age int //年龄
	Bothday string //出生年月
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


}

func main() {

		titleList := []string{"姓名", "性别", "年龄", "出生年月"}
		proList := []string{"Name", "Sex", "Age", "Bothday"}
		list := []Person{
			{"小明","男", 20,"2018-09-1"},
			{"小妮","女", 30,"2018-09-1"},
			{"小爽","女", 25,"2018-09-1"},
		}


	/*	for i, head := range titleList {
			fmt.Printf(string('A'+i)+"1" )
			fmt.Printf("a[%d]=[%s]\n", i,head )

		}

		for i, data := range list {
			fmt.Printf(string('B'+i)+strconv.Itoa(i+2) )
			//fmt.Printf("\na[%d]=[%s]\n", i,data )
			for _, pro := range proList {
				v := reflect.ValueOf(data)
				fmt.Println(v.FieldByName(pro))//根据属性名称获取值
			}

		}*/
	writeExcel("./Workbook",titleList ,proList,list);
	readExcel("./Workbook.xlsx")
}

/**
写excel
filePath:文件路径
titleList：列名：[姓名，性别，年龄，出生日期]
proList：类的属性[Name,Sex,Aage,Borthday]
Person：数据集合[person]
 */
func  writeExcel(filePath string,titleList []string,proList []string,list []Person)  {


	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet( "Sheet1")
	// Set value of a cell.
	for i, head := range titleList {
		xlsx.SetCellValue("Sheet1",string('A'+i)+"1", head )
	}


	//xlsx.SetCellValue("Sheet1", "A1", "姓名")
	//xlsx.SetCellValue("Sheet1", "B1", "年龄")
	for i, data := range list {
		for j, pro := range proList {
			v := reflect.ValueOf(data)
			cel :=v.FieldByName(pro)
			//fmt.Println(cel)//根据属性名称获取值
			xlsx.SetCellValue("Sheet1", string('A'+j)+strconv.Itoa(i+2), cel)
		}
	}


	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs(filePath+".xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("创建excel完成")
}


/**
读取excel数据
 */
func readExcel(filepath string) {
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx. GetCellValue("Sheet1", "B2")//直接指定读取
	fmt.Println(cell)

	sheetMap := xlsx.GetSheetMap()
	fmt.Println(sheetMap)
	for k, v := range sheetMap {
		fmt.Printf("%d=%s\n", k, v)
		//v := xlsx.GetSheetName(1)//不遍历也可以直接指定  默认从1开始
		rows := xlsx.GetRows(v)
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}

	}


}