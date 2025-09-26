package controller

import (
	"go-crud/model"
	"go-crud/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Data struct {
	DB  *gorm.DB
	log *logrus.Logger
}

func NewData(db *gorm.DB, log *logrus.Logger) Data {
	return Data{
		DB:  db,
		log: log,
	}
}

func (d *Data) GetData(c *gin.Context) {
	data, err := service.GetData(d.DB, c)
	if err != nil {
		d.log.WithField("module", "service").Error("Failed to get data")
		c.JSON(500, err)
		return
	}

	d.log.WithField("module", "service").Info("Success to get data")
	responseJSON := struct {
		Message string       `json:"message"`
		Type    string       `json:"type"`
		Data    []model.Data `json:"data"`
	}{
		Message: "Success Get Data",
		Type:    "data",
		Data:    data,
	}
	c.JSON(200, responseJSON)
}

func (d *Data) CreateData(c *gin.Context) {
	var parameter model.Data

	err := c.ShouldBindJSON(&parameter)
	if err != nil {
		d.log.WithField("module", "service").Error("Failed to create data")
		c.JSON(400, err)
		return
	}

	d.log.WithField("module", "service").Info("Success to create data")
	err = service.CreateData(d.DB, c, parameter)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (d *Data) UpdateData(c *gin.Context) {
	var parameter model.Data
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	parameter.ID = idInt

	err := c.ShouldBindJSON(&parameter)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = service.UpdateData(d.DB, c, parameter)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func (d *Data) DeleteData(c *gin.Context) {
	id := c.Param("id")
	service.DeleteData(d.DB, c, id)
	c.JSON(200, gin.H{"status": "success"})
}
