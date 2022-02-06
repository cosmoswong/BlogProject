package controllers

import (
    "github.com/astaxie/beego"
    log "github.com/inconshreveable/log15"
)

type ExitController struct {
    beego.Controller
}

func (this ExitController) Get()  {
    loginUser := this.GetSession("loginUser")
    log.Info("will remove sessionUser:",loginUser)
    this.DelSession("loginUser")
    this.Redirect("/",302)
    
}
