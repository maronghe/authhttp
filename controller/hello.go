package controller

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (hello *HelloController) Hello() {

	hello.Ctx.Output.SetStatus(200)
	hello.Ctx.Output.Body([]byte("Hello"))
	hello.ServeJSON()
}
