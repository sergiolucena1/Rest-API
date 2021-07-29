package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sergiolucena1/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/:id", controller.ShowProduct) //  mostrar produtos por id
			products.GET("/", controller.ShowProducts) // mostrar todos os produtos
			products.POST("/", controller.CreateProduct) //  criar produto
			products.PUT("/", controller.UpdateProduct) //  atualizar produto
			products.DELETE("/:id", controller.DeleteProduct) //  deletar produto
		}
	}
	return router
}
