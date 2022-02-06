package routers

import (
	"BlogProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
    beego.Router("/exit/:name",&controllers.ExitController{})
    beego.Router("/article/add",&controllers.ArticleController{})
	beego.Router("/article/update",&controllers.ArticleController{})
	beego.Router("/article/:id",&controllers.ArticleController{},"GET:ShowArticleDetail")
	beego.Router("/article/delete",&controllers.ArticleController{},"GET:DeleteArticleById")
	//标签
	beego.Router("/tags", &controllers.TagsController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})
    //上传
	beego.Router("/upload", &controllers.AlbumController{})
	beego.Router("/aboutme", &controllers.AboutMeController{})


}
