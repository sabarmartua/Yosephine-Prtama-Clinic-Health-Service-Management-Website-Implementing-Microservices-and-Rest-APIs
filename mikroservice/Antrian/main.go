package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Antrian/conn"
	"github.com/sabarmartua/Antrian/controller"
	"github.com/sabarmartua/Antrian/repository"
	"github.com/sabarmartua/Antrian/service"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = conn.SetupDatabaseConnection()
	antrianRepository    repository.AntrianRepository    = repository.NewAntrianRepository(db)
	antrianService       service.AntrianService          = service.NewAntrianService(antrianRepository)
	antrianController    controller.AntrianController    = controller.NewAntrianController(antrianService)
)

func main() {
	
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	antrianRoutes := r.Group("/api/antrian")
	{
		antrianRoutes.POST("/create", antrianController.Create)
		antrianRoutes.GET("/all", antrianController.GetAll)
		antrianRoutes.GET("/user/:userID", antrianController.GetByUserID)
		antrianRoutes.DELETE("/delete/:id", antrianController.Delete)
	}

	r.Run("192.168.154.117:8091")
}
