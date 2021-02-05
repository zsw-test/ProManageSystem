package DB

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//mysql连接单例
var Mysqldb *gorm.DB

//mysql初始化
func MysqlInit(connstr string) {
	db, err := gorm.Open("mysql", connstr)
	if connstr == "" || err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	Mysqldb = db
	db.LogMode(true)

}
