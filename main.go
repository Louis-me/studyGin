package main

import (
	"example.com/myGin/database"
	"example.com/myGin/routers"
)

func main() {
	database.InitDb()
	defer database.Close()
	router := routers.InitRouter() //指定路由
	router.Run(":8000")            //在8000端口上运行

}
