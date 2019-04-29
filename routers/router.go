package routers

import (
	"go_where/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "*:Get")
    beego.Router("/getData", &controllers.GetDataController{}, "get:Get;post:PostFunc")
    beego.Router("/getList", &controllers.GetListController{})
    beego.Router("/getComment", &controllers.GetCommentController{})
}
