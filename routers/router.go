package routers

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/AiminGubian/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/sections/:id", &controllers.MainController{}, "get:Section")
	beego.Router("/entries/:id", &controllers.MainController{}, "get:Entry")
	beego.Router("/host", &controllers.MainController{}, "get:HostIntro")
	beego.Router("/gallery", &controllers.MainController{}, "get:Gallery")
	beego.Router("/express", &controllers.MainController{}, "get:ExpressIntro")
	beego.Router("/express/charts", &controllers.MainController{}, "get:Charts")
	beego.Router("/express/timeline", &controllers.MainController{}, "get:TimeLine")
	beego.Router("/express/propaganda", &controllers.MainController{}, "get:Propaganda")

	beego.Router("/admin", &controllers.AdminController{}, "get:Index")
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login;post:PostLogin")
	beego.Router("/admin/sections/:id", &controllers.AdminController{}, "get:Section")
	beego.Router("/admin/entries/:id", &controllers.AdminController{}, "get:Entry;post:PostEntry")
	beego.Router("/admin/entries/delete", &controllers.AdminController{}, "get:DeleteEntry")
	beego.Router("/admin/broadcast", &controllers.AdminController{}, "post:PostBroadcast")
}
