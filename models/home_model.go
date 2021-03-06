package models

import (
    "fmt"
    "github.com/astaxie/beego"
    "strconv"
)

type HomeFooterPageCode struct {
    HasPre bool
    HasNext bool
    ShowPage string
    PreLink string
    NextLink string
}


//-----------翻页-----------
//page是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
    pageCode := HomeFooterPageCode{}
    //查询出总的条数
    num := GetArticleRowsNum()
    //从配置文件中读取每页显示的条数
    pageRow, _ := beego.AppConfig.Int("articleListPageNum")
    //计算出总页数
    allPageNum := (num-1)/pageRow + 1

    pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

    //当前页数小于等于1，那么上一页的按钮不能点击
    if page <= 1 {
        pageCode.HasPre = false
    } else {
        pageCode.HasPre = true
    }

    //当前页数大于等于总页数，那么下一页的按钮不能点击
    if page >= allPageNum {
        pageCode.HasNext = false
    } else {
        pageCode.HasNext = true
    }

    pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
    pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
    return pageCode

}