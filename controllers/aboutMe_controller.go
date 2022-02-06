package controllers

type AboutMeController struct {
    BaseController
}

func (this *AboutMeController) Get()  {
    this.Data["wechat"] = "微信：**"
    this.Data["qq"] = "QQ：**"
    this.Data["tel"] = "Tel：**"
    this.TplName = "aboutMe.html"
}

