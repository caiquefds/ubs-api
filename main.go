package main

import (
	"ubs-api/controllers"
	"ubs-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
  
	models.ConnectDatabase()
  
	r.GET("/UBS", controllers.FindAllUBS) 
	r.GET("/UBS/city/:city",controllers.FindUBSByCityName)
	r.GET("/UBS/uf/:uf",controllers.FindUBSByUF)

  
	r.Run()
  }
  
