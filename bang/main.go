package main

import (
	"Bang/model"
	"Bang/router"
	"log"

	"github.com/gin-gonic/gin"
)

// @title 2021-bang
// @version 1.0
// @description This is a project made by CCNU
// @Base /Bang/v1
// @Schemas http

func main() {
	r := gin.Default()
	model.DB = model.Initdb()
	defer model.DB.Close()
	router.Router(r)
	if err := r.Run(":10086"); err != nil {
		log.Fatal(err.Error())
	}
}
