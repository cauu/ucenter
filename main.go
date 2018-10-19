package main

import (
	"fmt"
	"os"
	"ucenter/library/types"
	_ "ucenter/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

func main() {
	host := os.Getenv("MYSQL_HOST")

	if host == "" {
		initOrm("127.0.0.1:3306")
	} else {
		initOrm(host)
	}

	beego.Run()
}

func initOrm(host string) {
	user := "root"
	pwd := "123456"

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/UCENTER?charset=utf8&loc=UTC", user, pwd, host)

	types.InitIDGen("123")

	orm.RegisterDriver("mysql", orm.DRMySQL)

 	orm.RegisterDataBase("default", "mysql", connStr)
	// orm.RegisterDataBase("default", "mysql", "root@tcp(127.0.0.1:3306)/UCENTER?charset=utf8&loc=UTC")

	orm.RunCommand()
	orm.Debug = true
	orm.DebugLog = orm.NewLog(os.Stdout)

	orm.RunSyncdb("default", false, false)
}
