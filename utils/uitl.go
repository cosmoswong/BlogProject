package utils

import (
    "crypto/md5"
    "database/sql"
    "fmt"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    "log"
    "time"
)
var db *sql.DB

func InitMysql()  {
    fmt.Println("InitMysql....")
    driverName := beego.AppConfig.String("driverName")

    //注册数据库驱动
    orm.RegisterDriver(driverName, orm.DRMySQL)

    //数据库连接
    user := beego.AppConfig.String("mysqluser")
    pwd := beego.AppConfig.String("mysqlpwd")
    host := beego.AppConfig.String("host")
    port := beego.AppConfig.String("port")
    dbname := beego.AppConfig.String("dbname")

    //dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
    dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

    db1, err := sql.Open(driverName, dbConn)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        db = db1
        //CreateTableWithUser()
    }

}

//创建用户表
func CreateTableWithUser() {
    sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

    ModifyDB(sql)
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
    log.Println("sql: %s,args%:",sql,args)
    result, err := db.Exec(sql, args...)
    if err != nil {
        log.Println(err)
        return 0, err
    }
    count, err := result.RowsAffected()
    if err != nil {
        log.Println(err)
        return 0, err
    }
    return count, nil
}


//查询
func QueryRowDB(sql string) *sql.Row{
    return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
    return db.Query(sql)
}


func MD5(str string) string {
    return fmt.Sprintf("%x",md5.Sum([]byte(str)))
}

func SwitchTimeStampToData(param int64) string  {
   return time.Unix(param,0).Format("2006-01-02")
}
