package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/shipments-service/api"
	"github.com/udayangaac/shipments-service/api/controllers"
	"github.com/udayangaac/shipments-service/config"
	"github.com/udayangaac/shipments-service/repo"
	"github.com/udayangaac/shipments-service/repo/entity"
	"github.com/udayangaac/shipments-service/yamlmgr"
	_ "gorm.io/driver/mysql"
)

func main() {

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	yamlMgr := yamlmgr.NewYamlManager()

	// Read configurations.
	config.Configurations{
		new(config.ServerConfig),
		new(config.DatabaseConfig),
	}.Init(yamlMgr)

	db, err := getDatabase(config.DatabaseConf)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&entity.Shipment{},
		&entity.User{},
		&entity.Consignee{},
		&entity.Shipper{},
		&entity.Invoice{},
	)

	userCtrl := controllers.UserController{
		UserRepo: repo.NewUserRepo(db),
	}

	shipmentCtrl := controllers.ShipmentController{
		LineChan: make(chan []string),
		ShipmentRepo: repo.NewShipmentRepo(db),
	}

	shipmentCtrl.InitWritingPool()
	go api.GetEngine(userCtrl, shipmentCtrl).Run(fmt.Sprintf(":%v", config.ServerConf.Port))

	<-osSignal
}

func getDatabase(dbConf config.DatabaseConfig) (db *gorm.DB, err error) {
	connectionString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbConf.UserName,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	)
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return
	}
	db.SetLogger(&customLogger{})
	db.LogMode(true)
	return
}

type customLogger struct{}

func (c *customLogger) Print(v ...interface{}) {
	log.Trace(v)
}
