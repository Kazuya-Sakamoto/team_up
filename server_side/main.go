package main

import (
	_ "app/server_side/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

