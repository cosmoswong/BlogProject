package controllers

import (
    "BlogProject/models"
)

type HomeController struct {
    BaseController
}

func (this *HomeController) Get()  {
    tag:=this.GetString("tag")

    page, _ := this.GetInt("page")
    if page <= 0 {
        page = 1
    }
    var artList []models.Article


    if len(tag) >0 {
        artList= models.QueryArticleWithTag(tag)
    }else {
        artList, _ = models.FindArticleWithPage(page)
        this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
    }

    this.Data["HasFooter"] = true

    this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
    this.TplName="home.html"

}
