package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (s *Service) InitDataBase(){
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			s.Conf.DB.User,
			s.Conf.DB.Password,
			s.Conf.DB.Addr,
			s.Conf.DB.Name,
		))

	if err != nil{
		panic(err)
	}

	s.DB = db
	s.DB.AutoMigrate(&User{}, &Token{})
}