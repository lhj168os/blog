package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/blog.log"}`)
	beego.Run()
}

