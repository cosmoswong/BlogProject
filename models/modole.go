package models

import (
    "database/sql"
    "github.com/astaxie/beego/logs"
    _ "github.com/go-sql-driver/mysql"
)

func init()  {
    db, err := sql.Open("mysql", "root:Admin@123@tcp(127.0.0.1:3306)/goBlog")
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    logs.Info("连接数据库成功")
}
