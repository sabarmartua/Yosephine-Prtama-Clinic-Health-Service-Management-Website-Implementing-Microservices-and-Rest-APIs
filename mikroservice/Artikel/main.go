package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Artikel/conn"
	"github.com/sabarmartua/Artikel/controller"
	"github.com/sabarmartua/Artikel/repository"
	"github.com/sabarmartua/Artikel/service"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                     = conn.SetupDatabaseConnection()
	artikelRepository repository.ArtikelRepository = repository.NewArtikelRepository(db)
	artikelService    service.ArtikelService       = service.NewArtikelService(artikelRepository)
	artikelController controller.ArtikelController = controller.NewArtikelController(artikelService)
)

func main() {
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	artikelRoutes := r.Group("/api/artikel")
	{
		artikelRoutes.GET("/all", artikelController.All)
		artikelRoutes.GET("/:id", artikelController.FindByID)
		artikelRoutes.POST("/create", artikelController.Insert)
		artikelRoutes.PUT("/update/:id", artikelController.Update)
		artikelRoutes.DELETE("/delete/:id", artikelController.Delete)
		artikelRoutes.GET("/:id/related", artikelController.RelatedByCategory)
	}

	r.Run("192.168.154.117:8084")
}
