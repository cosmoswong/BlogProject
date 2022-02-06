package main

import (
	_ "BlogProject/routers"
	"BlogProject/utils"
	"github.com/astaxie/beego"

)

func main() {
	utils.InitMysql()
	beego.Run()
}

