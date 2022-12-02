package main

import (
	"bee-api-demo/models"
	_ "bee-api-demo/routers"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterAll()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	fmt.Printf("hello")
}
