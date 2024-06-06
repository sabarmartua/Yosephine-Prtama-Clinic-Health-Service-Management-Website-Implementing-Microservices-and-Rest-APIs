package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Ulasan/conn"
	"github.com/sabarmartua/Ulasan/controller"
	"github.com/sabarmartua/Ulasan/repository"
	"github.com/sabarmartua/Ulasan/service"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = conn.SetupDatabaseConnection()
	ulasanRepository     repository.UlasanRepository     = repository.NewUlasanRepository(db)
	ulasanService        service.UlasanService           = service.NewUlasanService(ulasanRepository)
	ulasanController     controller.UlasanController     = controller.NewUlasanController(ulasanService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	ulasanRoutes := r.Group("/api/ulasan")
	{
		ulasanRoutes.POST("/create", ulasanController.Create)
		ulasanRoutes.PUT("/update/:id", ulasanController.Update)
		ulasanRoutes.DELETE("/delete/:id", ulasanController.Delete)
		ulasanRoutes.GET("/all", ulasanController.GetAll)
	}

	r.Run("192.168.154.117:8087")
}
