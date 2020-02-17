package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

type form struct {
	Code 	string 	`json:"code"`
}

func (s *Service) GetFlag(c *gin.Context) (int, interface{}){
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool){
		return s.makeErrJSON(403, 40300,  "你不是 admin 哦~")
	}
	var input form
	err := c.ShouldBindJSON(&input)
	if err != nil{
		return s.makeErrJSON(403, 40300, "Wrong code!")
	}

	totp := gotp.NewDefaultTOTP("X5JMTFGT4FVJ34GV")
	if input.Code == totp.Now(){
		return s.makeSuccessJSON(s.Conf.Server.Flag)
	}
	return s.makeErrJSON(403, 40300, "Wrong code!")
}
