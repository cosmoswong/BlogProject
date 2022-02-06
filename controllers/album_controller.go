package controllers

import (
    "BlogProject/models"
    "errors"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "io"
    "os"
    "path/filepath"
    "time"
)

type AlbumController struct {
    BaseController
}

func (this *AlbumController) Get()  {
    albums, err := models.QueryAllAlbum()
    if err !=nil{
        logs.Error(err)
    }
    this.Data["Album"] = albums
    this.TplName="album.html"
}

func (this *AlbumController) Post()  {
    fileData, fileHeader, err := this.GetFile("upload")
    if err != nil {
        this.responseErr(err)
        return
    }
    fmt.Println("name:", fileHeader.Filename, fileHeader.Size)
   //文件大小限制
    fileSize,_:=beego.AppConfig.Int("fileSize")
    if fileHeader.Size > int64(fileSize*1024*1024) {
         this.responseErr(errors.New(fmt.Sprintf("文件大小过大!控制在%dM内",fileHeader)))
    }

   //文件后缀名限制，文件格式限制
    filename := fileHeader.Filename
    fileExt := filepath.Ext(filename)


    if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".gif" && fileExt != ".jpeg" {
       this.responseErr(fmt.Errorf("不支持的格式类型，只能是.jpg、.png、.gif、.jpeg"))
        return
    }

    //保存文件
    imagePath:=beego.AppConfig.String("imagePath")
    destPath:=fmt.Sprintf("%s/images/%s",imagePath,time.Now().Format("2006-01-02"))

    err = os.MkdirAll(destPath, os.ModePerm)
    if err != nil {
        logs.Error(err)
        this.responseErr(err)
        return
    }
    fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)
    filePathStr:=filepath.Join(destPath,fileName)

    destFile, err := os.Create(filePathStr)
    if err != nil {
        this.responseErr(err)
        return
    }

    _, err = io.Copy(destFile, fileData)
    if err != nil {
        this.responseErr(err)
        return
    }

    //文件信息写入到数据库
    models.InsertAlbum(models.Album{Filepath: filePathStr,Filename: fileName,Status: 0,Createtime: time.Now().Unix()})

    this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
    this.ServeJSON()

}
func (this *AlbumController) responseErr(err error) {
    this.Data["json"] = map[string]interface{}{"code": 0, "message": err.Error()}
    this.ServeJSON()
}


