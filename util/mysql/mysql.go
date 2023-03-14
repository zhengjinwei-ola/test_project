package mysql

import (
	"database/sql"
	"strings"
	"test_project/app/pb"
	"test_project/util/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/util/gconv"
)

var Db *sql.DB

func InitDB() *sql.DB {
	mysqlConfig := config.ConfigServer.Get("mysql")
	mysqlConfigS := &pb.MysqlConfig{}
	err := gconv.Struct(mysqlConfig, mysqlConfigS)
	if err != nil {
		panic(err)
	}
	//Golang数据连接："用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8"
	path := strings.Join([]string{mysqlConfigS.GetUserName(), ":", mysqlConfigS.GetPassword(), "@tcp(", mysqlConfigS.GetIp(), ":", mysqlConfigS.GetPort(), ")/", mysqlConfigS.GetDatabase(), "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	if err != nil {
		//如果打开数据库错误，直接panic
		panic(err)
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	//验证连接
	if err := db.Ping(); err != nil {
		panic(err)
	}
	//将数据库连接的指针引用返回
	Db = db
	return db
}
