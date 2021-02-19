package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/royal-dargon/go1/Library-test/model"
)

type Userinfo struct {
	Id          string `json:"user_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Userpicture string `json:"user_picture"`
	motto       string
}

func Register(name string, password string) string {
	count := 0
	if err := model.DB.Table("users").Count(&count).Error; err != nil {
		log.Println("registerError" + err.Error())
		return ""
	}
	temp := 0 + count
	id := strconv.Itoa(temp)
	user := Userinfo{Id: id, UserName: name, Password: password}
	if err := model.DB.Table("users").Create(&user).Error; err != nil {
		log.Println("registError" + err.Error())
		return ""
	}
	return id
}

func CreateUser(c *gin.Context) {
	var user Userinfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误."})
		return
	}
	user_id := Register(user.UserName, user.Password)
	c.JSON(200, gin.H{
		"user_id": user_id,
	})
}

func main() {
	DB := model.Initdb()
	defer DB.Close()
	router := gin.Default()
	router.POST("/api/v1/user", CreateUser)
	router.Run()
}
