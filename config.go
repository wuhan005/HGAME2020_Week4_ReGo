package main

import "github.com/BurntSushi/toml"

type DB struct {
	User 		string	`toml:"user"`
	Password 	string	`toml:"password"`
	Addr 		string	`toml:"addr"`
	Name 		string	`toml:"name"`
}

type Server struct {
	Port 	string  `toml:"port"`
	Flag 	string 	`toml:"flag"`
}

type Config struct {
	DB		DB		`toml:"db"`
	Server 	Server 	`toml:"server"`
}

func (s *Service) InitConfig(){
	var c *Config
	_, err := toml.DecodeFile("./conf/config.toml", &c)
	if err != nil{
		panic(err)
	}
	s.Conf = c
}