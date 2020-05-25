package router

import (
	"authhttp/controller"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/authz"
	"github.com/casbin/casbin"

	"github.com/astaxie/beego/plugins/cors"
)

func Init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{}, // TODO
		AllowHeaders:     []string{"Content-Type", "Transaction-Id"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*",
		beego.BeforeRouter,
		authz.NewAuthorizer(casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")))

	beego.Router("/hello", &controller.HelloController{}, "get:Hello")
	beego.Router("/admin", &controller.HelloController{}, "get:Hello")

}
