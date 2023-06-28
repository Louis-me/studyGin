package main

import (
	"example.com/myGin/database"
	"example.com/myGin/models"
	"example.com/myGin/routers"
)

func main() {
	database.InitDb()
	database.Db.AutoMigrate(&models.User{})
	defer database.Close()
	router := routers.InitRouter() //指定路由
	router.Run(":8000")            //在8000端口上运行

}
