package main

import (
	"Bang/model"
	"Bang/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model.DB = model.Initdb()
	defer model.DB.Close()
	router.Router(r)
	if err := r.Run(":10086"); err != nil {
		log.Fatal(err.Error())
	}
}
