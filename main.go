package main

import (
	"authhttp/router"

	"github.com/astaxie/beego"
)

func main() {

	router.Init()

	beego.Run()

}
