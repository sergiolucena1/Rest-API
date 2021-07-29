package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/models"
	"strconv"
)

//Primeiro endpoint
func ShowProduct(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id) // convertendo pra inteiro
	if err != nil{
		c.JSON(400,gin.H{
			"error": "ID tem que ser inteiro",
		})

		return
	}
	db := database.GetDatabase()

	var product models.Product
	err = db.First(&product, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não consigo encontrar o produto:" + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}

//Segundo endpoint
func CreateProduct(c *gin.Context){
	db := database.GetDatabase()

	var product models.Product

	err:= c.ShouldBindJSON(&product)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create product: " + err.Error(),
		})
	}
	c.JSON(200, product)
}

func ShowProducts(c *gin.Context){

	db := database.GetDatabase()

	var products []models.Product
	err := db.Find(&products).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list products: " + err.Error(),
		})
		return
	}
	c.JSON(200, products)

}

func UpdateProduct(c *gin.Context){
	db := database.GetDatabase()

	var product models.Product

	err:= c.ShouldBindJSON(&product)
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update product: " + err.Error(),
		})
	}
	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id) // convertendo pra inteiro
	if err != nil{
		c.JSON(400,gin.H{
			"error": "ID tem que ser inteiro",
		})

		return
	}
	db := database.GetDatabase()

	err = db.Delete(&models.Product{},newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete product: " + err.Error(),
		})
		return
	}
	c.Status(204)
}