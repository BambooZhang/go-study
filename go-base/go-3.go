package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

/**
CREATE TABLE `user_info` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
)ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

 */




func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	// 插入数据
	stmt, err := db.Prepare("INSERT user_info SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("test", " 研发部门", "2017-12-09")
	checkErr(err)
	id, err := res.LastInsertId()//获取插入后的主键
	checkErr(err)
	fmt.Println(id)


	//查询数据，指定字段名，返回sql.Rows结果集
	// 查询数据
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("%d\t%s\t%s\t%s\n",uid,username,department,created)
	}


	// 更新一条数据
	result,err := db.Exec("UPDATE user_info  SET username=? WHERE username=?", "test1","test")
	checkErr(err)
	num,err := result.RowsAffected()
	checkErr(err)
	fmt.Printf("修改的数据条数%d\n",num)



	// 删除数据,预编译sql
	stmt, err = db.Prepare("delete from user_info where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id) //执行sql
	checkErr(err)
	num1,err := res.RowsAffected()//影响的行数
	checkErr(err)
	fmt.Printf("删除的数据条数%d\n",num1)
	db.Close()







}



func checkErr(err error) {
	if err != nil {
	panic(err)
	}
}