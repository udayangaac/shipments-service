package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/udayangaac/shipments-service/repo"
)

type ShipmentController struct {
	ShipmentRepo repo.ShipmentRepo
}

func (c *ShipmentController) Upload(ctx *gin.Context) {

}
