package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Conf	*Config
	DB		*gorm.DB
	Router	*gin.Engine
}

func (s *Service) Init(){
	s.InitConfig()
	s.InitDataBase()
	s.InitRouter()
}