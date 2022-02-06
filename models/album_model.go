package models

import (
    "BlogProject/utils"
    "github.com/astaxie/beego/logs"
)

type Album struct {
    Id         int
    Filepath   string
    Filename   string
    Status     int
    Createtime int64
}

//-------插入图片---------------
func InsertAlbum(album Album) (int64, error) {
    return utils.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",
        album.Filepath, album.Filename, album.Status, album.Createtime)
}

func QueryAllAlbum() ([]Album,error)  {

    sql:="select id,filepath,filename,status,createtime from album"

    rows, err := utils.QueryDB(sql)
    if err != nil{
        logs.Error(err)
    }
    var id int
    var filepath string
    var filename string
    var status int
    var createtime int64

    var albums []Album
    for rows.Next() {
        rows.Scan(&id, &filepath, &filename,&status,&createtime)
        album := Album{id, filepath, filename, status, createtime}
        albums=append(albums, album)
    }

    return albums,nil
}
