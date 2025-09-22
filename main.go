package main

import (
	"go-crud/controller"
	"go-crud/db"
	"go-crud/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db := db.Connect()

	db.AutoMigrate(&model.Data{})

	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	data := controller.NewData(db, log)

	router.GET("/data", data.GetData)
	router.POST("/data", data.CreateData)
	router.PUT("/data/:id", data.UpdateData)
	router.DELETE("/data/:id", data.DeleteData)
	log.Println("Server running at http://localhost:8080")
	router.Run()
}
