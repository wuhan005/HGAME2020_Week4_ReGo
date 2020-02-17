package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model 			`json:"-"`
	Mail 		string 	`json:"mail"`
	Name 		string 	`json:"name"`
	Password 	string 	`json:"password"`
}

type Token struct {
	gorm.Model
	Uid 	uint
	Str 	string
}