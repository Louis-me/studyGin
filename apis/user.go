package apis

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/myGin/middleware/jwt"
	"example.com/myGin/models"
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt/v5" // go get -u github.com/golang-jwt/jwt/v5
)

// 存放登录后的token和用户数据
type LoginResult struct {
	User  interface{}
	Token string
}

func UserLogin(c *gin.Context) {
	var user models.User
	// c.ShouldBindJSON
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		msg, err := user.Login() // 调用model层的方法
		if err != nil {
			if err.Error() == "record not found" {
				c.JSON(http.StatusOK, gin.H{
					"msg":  "用户不存在",
					"user": nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg":  "登陆错误",
					"user": nil,
				})

			}

		} else {
			GengerateToken(c, msg) //创建toke
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg":  "登陆成功",
			// 	"user": msg,
			// })
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}

}

// 生成token
func GengerateToken(c *gin.Context, user models.User) {
	j := &jwt.JWT{
		[]byte("newtrekWang"), // 秘钥
	}
	claims := jwt.CustomClaims{
		user.Id,
		user.Name,
		user.Password,
		jwtgo.RegisteredClaims{
			ExpiresAt: jwtgo.NewNumericDate(time.Now().Add(time.Hour * 24)), //定义过期时间为24小时
			Issuer:    "newtrekWang",                                        //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return

}

// 获取用户列表
func GetUserList(c *gin.Context) {
	var user models.User
	users, err := user.UserList()
	if err != nil {
		fmt.Println("get user list error:", err)
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": users,
	// })
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"users": users,
	})
}

func AddNewUser(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		id, err := user.UserAdd() // 调用model层的方法
		fmt.Println("新增的用户id:", id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "创建用户失败",
				"code": -1,
				"err":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "创建登录用户成功",
				"code": 0,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}

}

func DelUser(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		flag := user.DeleteUser() // 调用model层的方法
		if !flag {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除用户失败",
				"code": -1,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除用户成功",
				"code": 0,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}

}

func EditUser(c *gin.Context) {
	var user models.User
	if c.Bind(&user) == nil { //把客户端格式传过来的数据绑定到结构体user中去
		flag := user.UpdateUser() // 调用model层的方法
		if !flag {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "更新用户失败",
				"code": -1,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "更新用户成功",
				"code": 0,
			})
		}
	} else {
		c.JSON(400, gin.H{"JSON=== status": "binding JSON error!"})
	}

}
