package conn

import (
	"common/clog"
	"common/server"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB *sqlx.DB

func init() {
	clog.Infof("start connect db")
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	database := server.Config.Database
	//"root:root@tcp(120.78.205.38:3306)/zwy_match?charset=utf8
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", database.User, database.Password, database.Host, database.Port, database.Name, database.Charset)
	//db, err := sql.Open("mysql", source);
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		clog.Fatalf("connect db fail,err=%+v", err)
	}

	if err = db.Ping(); err != nil {
		clog.Fatalf("connect db fail,err=%+v", err)
	}
	db.SetConnMaxLifetime(10 * time.Second)
	clog.Infof("start connect db success")
	DB = db
}
