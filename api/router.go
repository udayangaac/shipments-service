package api

import (
	"github.com/gin-gonic/gin"
	"github.com/udayangaac/shipments-service/api/controllers"
)

func GetEngine(userCtrl controllers.UserController, shipementCtrl controllers.ShipmentController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/user", userCtrl.CreateUser)
		v1.GET("/user", userCtrl.LoginUser)
		v1.GET("/shipments", shipementCtrl.Upload)
	}
	return r
}
