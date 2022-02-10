package main

import (
	_ "Campus-Server/routers"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	err := orm.RegisterDataBase("default", "postgres", sqlconn)
	if err != nil {
		panic(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
