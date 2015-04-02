package main

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/AiminGubian/models"
	_ "github.com/dyzdyz010/AiminGubian/routers"
	. "github.com/qiniu/api/conf"
)

func main() {

	// Qiniu
	ACCESS_KEY = models.Appconf.String("qiniu::ak")
	SECRET_KEY = models.Appconf.String("qiniu::sk")

	beego.Run()
}
