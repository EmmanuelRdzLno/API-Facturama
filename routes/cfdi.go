package routes

import (
	"github.com/gin-gonic/gin"
	"facturama-api/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/cfdi", controllers.CreateCfdi)
		api.POST("/sandbox/cfdi", controllers.CreateTasteCfdi)
		api.GET("/cfdi", controllers.GetCfdis)
		api.GET("/cfdi/:id/download", controllers.DownloadFiles)
	}
}
