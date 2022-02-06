package controllers

import (
    "BlogProject/models"
    "BlogProject/utils"
    "fmt"
    "github.com/astaxie/beego"
    log "github.com/inconshreveable/log15"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get()  {
    this.TplName = "login.html"
}
func (this *LoginController) Post()  {
    //获取表单信息
    username := this.GetString("username")
    password := this.GetString("password")

    fmt.Println(username, password)
    log.Info(username, password)

    param := models.QueryUserWithParam(username, utils.MD5(password))
    if param == 0{
        this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名不存在!"}
    }else {
        this.SetSession("loginUser",username)
        this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功!"}
    }
    this.ServeJSON()
}