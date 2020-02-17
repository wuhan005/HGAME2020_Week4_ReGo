package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func (s *Service) Register(c *gin.Context)(int, interface{}) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Name == "" || user.Mail == "" || user.Password == ""{
		return s.makeErrJSON(403,40300, "注册失败！")
	}

	if !strings.Contains(user.Mail, "@") {
		return s.makeErrJSON(403,40300, "邮箱格式不正确")
	}
	if len(user.Name) < 5 {
		return s.makeErrJSON(403,40300, "用户名长度需大于 5 位！")
	}
	if len(user.Password) < 8 {
		return s.makeErrJSON(403,40300, "密码长度需大于 8 位！")
	}

	var u1 User
	s.DB.Model(&User{}).Where(&User{Name: user.Name}).Find(&u1)
	if u1.Name != ""{
		return s.makeErrJSON(403,40301, "用户名重复，换个名字吧~")
	}

	var u2 User
	s.DB.Model(&User{}).Where(&User{Mail: user.Mail}).Find(&u2)
	if u2.Mail != ""{
		return s.makeErrJSON(403,40301, "电子邮箱重复")
	}

	tx := s.DB.Begin()
	if tx.Create(&user).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "注册失败！")
	}
	tx.Commit()
	return s.makeSuccessJSON("注册成功")
}

func (s *Service) Login(c *gin.Context) (int, interface{}){
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Mail == "" || user.Password == ""{
		return s.makeErrJSON(403,40300, "登录失败！")
	}

	var u User
	s.DB.Model(&User{}).Where(&User{
		Mail:     user.Mail,
		Password: user.Password,
	}).Find(&u)
	if u.Name != ""{
		token := s.NewToken(u.ID)
		return s.makeSuccessJSON(token)
	}
	return s.makeErrJSON(403,40301, "登录失败！")
}

func (s *Service) UpdateProfile(c *gin.Context) (int, interface{}){
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Password == ""{
		return s.makeErrJSON(403,40300, "修改个人信息失败！")
	}

	if len(user.Password) < 8 {
		return s.makeErrJSON(403,40300, "密码长度需大于 8 位！")
	}

	uid, _ := c.Get("uid")
	tx := s.DB.Begin()
	if s.DB.Model(&User{}).Where(&User{Model: gorm.Model{ID: uid.(uint)}}).Update(&user).RowsAffected != 1 {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, "修改个人信息失败！")
	}
	tx.Commit()
	return s.makeSuccessJSON("修改成功")
}

func (s *Service) GetProfile(c *gin.Context) (int ,interface{}){
	uid, _ := c.Get("uid")
	var user User
	s.DB.Model(&User{}).Where(&User{Model: gorm.Model{ID: uid.(uint)}}).Find(&user)
	user.Password = ""
	return s.makeSuccessJSON(user)
}