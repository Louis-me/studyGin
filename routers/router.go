package routers

import (
	"example.com/myGin/apis"
	"example.com/myGin/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	// 这样使用，并不能进行token验证
	// v1 := router.Group("/v1")
	// v1.Use(jwt.JWTAuth())
	router.POST("/login", apis.UserLogin)
	// 加载 HTML 模板文件
	router.LoadHTMLGlob("template/user/*")
	router.GET("/GetUserList", jwt.JWTAuth(), apis.GetUserList)
	router.POST("/AddNewUser", jwt.JWTAuth(), apis.AddNewUser)
	router.POST("/DelUser", jwt.JWTAuth(), apis.DelUser)
	router.POST("/EditUser", jwt.JWTAuth(), apis.EditUser)
	router.GET("/UserGet/:id", jwt.JWTAuth(), apis.GetUser)

	return router

}
