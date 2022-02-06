package models

import (
    "BlogProject/utils"
    "fmt"
    "github.com/astaxie/beego"
    log "github.com/inconshreveable/log15"
)

type Article struct {
    Id         int
    Title      string
    Tags       string
    Short      string
    Content    string
    Author     string
    Createtime int64
    //Status int //Status=0为正常，1为删除，2为冻结
}

func InsertArticle(article Article) (int64,error)  {
   return utils.ModifyDB("insert into article(title,author,tags,short,content,createtime) values(?,?,?,?,?,?)",
       article.Title,article.Author,article.Tags,article.Short,article.Content,article.Createtime)
}

func UpdateArticle(article Article)(int64,error)  {
   return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",
       article.Title,article.Tags,article.Short,article.Content,article.Id)
}

func DeleteArticleById(id int)  {
    utils.ModifyDB("delete from article where id=?", id)
}



//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
    //从配置文件中获取每页的文章数量
    num, _ := beego.AppConfig.Int("articleListPageNum")
    page--
    fmt.Println("---------->page", page)
    return QueryArticleWithPage(page, num)
}
func QueryArticleWithPage(page int,num int) ([]Article, error)  {
    sql := fmt.Sprintf("limit %d,%d", page*num, num)

    return QueryArticleWithCondition(sql)
}

func QueryArticleWithId(id int) Article {

    sql := fmt.Sprintf("where id= %d", id)
    articles, _ := QueryArticleWithCondition(sql)

       return articles[0]

}


//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
    rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
    if err != nil {
        log.Info("error:",err)
    }
    var paramList []string
    for rows.Next() {
        arg := ""
        rows.Scan(&arg)
        paramList = append(paramList, arg)
    }
    return paramList
}

func QueryArticleWithTag(tag string) []Article {
    sql := " where tags like '%&" + tag + "&%'"
    sql += " or tags like '%&" + tag + "'"
    sql += " or tags like '" + tag + "&%'"
    sql += " or tags like '" + tag + "'"
    articles, _ := QueryArticleWithCondition(sql)
    return articles
}



func QueryArticleWithCondition(sql string) ([]Article, error) {

    sql = "select id,title,tags,short,content,author,createtime from article " + sql
    log.Info("QueryArticleWithCondition:%s",sql)

    rows, err := utils.QueryDB(sql)

    if err != nil {
        return nil,err
    }

    var articles []Article
    for rows.Next() {
        id:=0
        tite:=""
        tags:=""
        short:=""
        content:=""
        author:=""
        var createtime int64
        createtime = 0
        rows.Scan(&id,&tite,&tags,&short,&content,&author,&createtime)
        artilce:= Article{id,tite,tags,short,content,author,createtime}
        articles=append(articles,artilce)
    }

    return articles,nil
}





//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
    if artcileRowsNum == 0 {
        artcileRowsNum = QueryArticleRowNum()
    }
    return artcileRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
    row := utils.QueryRowDB("select count(id) from article")
    num := 0
    row.Scan(&num)
    return num
}
