package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"ubs-api/models"

	"github.com/gin-gonic/gin"
)

func FindAllUBS(c *gin.Context) {
	var UBS []models.UBS

	models.DB.Find(&UBS)

	c.JSON(http.StatusOK, gin.H{"data": UBS})

}

func FindUBSByCityName(c *gin.Context) {

	var UBS []models.UBS
	var City models.City

	cityName := strings.ToUpper(c.Param("city"))

	log.Println("city: " + cityName)

	if err := models.DB.Where("city_name = ?", cityName).First(&City).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"404": "Dont exist an UBS in this city."})
		return
	}

	ibgeCode := strconv.FormatUint(uint64(City.IBGECode), 10)

	log.Println("code: " + ibgeCode)

	if err := models.DB.Where("ibge_code = ?", ibgeCode).Find(&UBS).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"500": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ubs": UBS})
}

func FindUBSByUF(c *gin.Context) {

	var UBS []models.UBS
	var City models.City

	uf := strings.ToUpper(c.Param("uf"))
	log.Println("uf: " + uf)

	if err := models.DB.Where("uf = ?", uf).First(&City).Error ; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"404": "Dont exist an UBS in this state."})
		return
	}

	log.Println("code :" + City.UFCode)
	if err := models.DB.Where("uf_code = ?", City.UFCode).Find(&UBS).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"500": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ubs": UBS})
}
