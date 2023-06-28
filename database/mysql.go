package database

//安装gorm
//go get gorm.io/gorm
// 安装mysql驱动
//go get gorm.io/driver/mysql
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql error:", err)
	}
	Db = database

}
func Close() {
	sqlDB, _ := Db.DB()
	sqlDB.Close()
}
