package controllers

import (
    "BlogProject/models"
    "strconv"
    "time"
)

type ArticleController struct {
    BaseController
}

func (this *ArticleController) Get()  {
    if this.IsLogin {
        id, _ := this.GetInt("id")
        if id != 0 {
            art := models.QueryArticleWithId(id)

            this.Data["Title"] = art.Title
            this.Data["Tags"] = art.Tags
            this.Data["Short"] = art.Short
            this.Data["Content"] = art.Content
            this.Data["Id"] = art.Id
        }
        this.TplName = "write_article.html"
    }else {
        this.TplName = "login.html"
    }

}

func (this *ArticleController) Post()  {
    title := this.GetString("title")
    tags := this.GetString("tags")
    short := this.GetString("short")
    content := this.GetString("content")
    id := this.GetString("id")
    atoi, _ := strconv.Atoi(id)
    article := models.Article{Id:atoi,Title: title, Tags: tags, Short: short, Content: content,Createtime: time.Now().Unix(),Author:this.LoginUser.(string)}

    var err error
    if atoi == 0 {
        _, err = models.InsertArticle(article)
    }else {
        _, err = models.UpdateArticle(article)
    }

    if err!= nil{
      this.Data["json"]= map[string]interface{}{"code": 0, "message": "error"}
    }else {
        this.Data["json"]= map[string]interface{}{"code": 1, "message": "ok"}
    }
    this.ServeJSON()
}

func (this *ArticleController) ShowArticleDetail()  {
    idStr:=this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    art := models.QueryArticleWithId(id)
    this.Data["Title"] = art.Title
    this.Data["Content"] = art.Content
    this.TplName="show_article.html"
}

func (this *ArticleController)  DeleteArticleById()  {
    id,_ :=this.GetInt("id")
    models.DeleteArticleById(id)
    this.Redirect("/", 302)
}