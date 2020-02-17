package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Service) InitRouter(){
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders: 	  []string{"*"},
	}))

	r.POST("/register", func(c *gin.Context) {
		c.JSON(s.Register(c))
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(s.Login(c))
	})

	authorized := r.Group("/")
	authorized.Use(s.AuthRequired())
	{
		authorized.PUT("/profile", func(c *gin.Context) {
			c.JSON(s.UpdateProfile(c))
		})

		authorized.GET("/profile", func(c *gin.Context){
			c.JSON(s.GetProfile(c))
		})
		
		authorized.POST("/flag", func(c *gin.Context) {
			c.JSON(s.GetFlag(c))
		})
	}

	s.Router = r
	panic(s.Router.Run(s.Conf.Server.Port))
}

func (s *Service) AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(s.makeErrJSON(403, 40300, "请登录"))
			c.Abort()
			return
		}

		var tokenData Token
		s.DB.Where(&Token{Str: token}).Find(&tokenData)
		if tokenData.ID == 0 {
			c.JSON(s.makeErrJSON(401, 40100, "请登录"))
			c.Abort()
			return
		}

		var user User
		s.DB.Model(&User{}).Where(&User{Model: gorm.Model{ID: tokenData.Uid}}).Find(&user)
		c.Set("isAdmin", user.Name == "admin")
		c.Set("uid", tokenData.Uid)
		c.Next()
	}
}