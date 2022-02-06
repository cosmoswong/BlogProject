## BlogProject

技术栈：go1.17.6，Beego 1.12.2

### 运行效果：

注册：

![image-20220206161540526](static/img/image-20220206161540526.png)

登录：

![image-20220206161715059](static/img/image-20220206161715059.png)

查看详情：

![image-20220206161820698](static/img/image-20220206161820698.png)

写博客：

![image-20220206161901305](static/img/image-20220206161901305.png)

相册：

![image-20220206161928580](static/img/image-20220206161928580.png)

标签：

![image-20220206161950957](static/img/image-20220206161950957.png)

关于我：

![image-20220206162105105](static/img/image-20220206162105105.png)



### API：

项目用到的路由：

```go
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
```



