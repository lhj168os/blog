package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/?:id",&controllers.Controller{})
	beego.AutoRouter(&controllers.Controller{})
}
