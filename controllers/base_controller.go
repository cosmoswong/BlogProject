package controllers

import (
    "github.com/astaxie/beego"
    log "github.com/inconshreveable/log15"
)

type BaseController struct {
    beego.Controller
    IsLogin   bool
    LoginUser interface{}
}

func (this *BaseController) Prepare()  {
    loginUser := this.GetSession("loginUser")
    log.Info("current session user:",loginUser)
    if loginUser !=nil {
       this.IsLogin =true
       this.LoginUser=loginUser
    }else {
        this.IsLogin =false
    }
    this.Data["IsLogin"] = this.IsLogin
    this.Data["loginUser"] = this.LoginUser
}
