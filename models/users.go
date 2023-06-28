package models

import (
	"fmt"

	"example.com/myGin/database"
)

type User struct {
	Name     string `form:"name",json:"name",bingding:"required",gorm:"unique;not null"`
	Password string `form:"password",json:"password",bingding:"required"gorm:"NOT NULL"`
	Id       int    `form:"id",gorm:"PRIMARY_KEY"`
}

// 登录
func (u *User) Login() (user1 User, err error) {
	obj := database.Db.Where("name=? and password=?", u.Name, u.Password).First(&user1)
	if err = obj.Error; err != nil {
		fmt.Printf("这是登陆错误  %v 和 %T", err, err)
		return
	}
	fmt.Println(user1)
	return

}

// 获取用户列表
func (u *User) UserList() (users []User, err error) {
	if err = database.Db.Find(&users).Error; err != nil {
		return
	}
	return

}

// 新增用户
func (u *User) UserAdd() (id int, err error) {
	result := database.Db.Create(&u)
	if result.Error != nil {
		err = result.Error
		fmt.Println("create is wrong,error:", err)
		return
	}
	fmt.Println("创建数据成功")
	id = u.Id
	return
}

// 删除用户
func (u *User) DeleteUser() (flag bool) {
	flag = true
	var users []User
	database.Db.Where("id=?", u.Id).Find(&users)
	if len(users) == 0 {
		fmt.Println("user is not exists")
		flag = false
		return
	}
	result := database.Db.Unscoped().Delete(&u)
	// 下面的语句有问题无论删除是否成功都成功
	if result.Error != nil {
		fmt.Println("del failed", result.Error)
		flag = false
		return
	}
	return
}

// 修改用户
func (u *User) UpdateUser() (flag bool) {
	flag = true
	result := database.Db.Where("Id=?", u.Id).Updates(&u)
	if result.Error != nil {
		flag = false
		err := result.Error
		fmt.Println("update failed", err)
		return
	}
	fmt.Println("update success")
	return
}
