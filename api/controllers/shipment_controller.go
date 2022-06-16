package controllers

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/repo"
	"github.com/udayangaac/shipments-service/repo/entity"
)

const numOfRoutine = 100

type ShipmentController struct {
	ShipmentRepo repo.ShipmentRepo
	LineChan     chan string
}

func (c *ShipmentController) Upload(ctx *gin.Context) {
	csvPartFile, _, openErr := ctx.Request.FormFile("file")
	if openErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   openErr.Error(),
		})
	}
	csvLines, readErr := csv.NewReader(csvPartFile).ReadAll()
	if readErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   readErr.Error(),
		})
	}
	for _, line := range csvLines {
		fmt.Println(line)
	}
}

func mapCsvLineToShipmen(line string) (entity.Shipment, error) {
	const (
		TRACKING_ID int = 0
		ORIGIN_PORT
		DST_PORT
		NUM_PKGS
		CONSIGNEE_NUM
		CONSIGNEE_NAME
		CONSIGNEE_ADDR
		CONSIGNEE_CONTACT_NO
		SHIPPER_NUM
		SHIPPER_NAME
		SHIPPER_ADDR
		SHIPPER_CONTACT_NO
		INVOICE_NUM
		INVOICE_AMT
		INVOICE_CURRENCY_CODE
	)
	// values := strings.Split(line, ",")
	// TODO: CSV Should not be comma separated. because address consist of commas.
	return entity.Shipment{}, nil
}

func (c *ShipmentController) InitWritingPool() {
	for i := 0; i < numOfRoutine; i++ {
		go func(lc chan string, sr repo.ShipmentRepo) {
			for {
				line := <-lc
				shipment, err := mapCsvLineToShipmen(line)
				if err != nil {
					log.Error(err)
					continue
				}
				err = sr.Save(context.Background(), shipment)
				if err != nil {
					log.Error(err)
				}
			}

		}(c.LineChan, c.ShipmentRepo)
	}
}
