package conf

import (
	"ProManageSystem/DB"
	"ProManageSystem/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	// 链接mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)
	DB.MysqlInit(dsn)

	//开启数据迁移
	model.Migration()
}
