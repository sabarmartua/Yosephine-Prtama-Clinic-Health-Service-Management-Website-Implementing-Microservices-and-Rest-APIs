package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/Kategori-Artikel/conn"
	"github.com/sabarmartua/Kategori-Artikel/controller"
	"github.com/sabarmartua/Kategori-Artikel/repository"
	"github.com/sabarmartua/Kategori-Artikel/service"
	"gorm.io/gorm"
)

var (
	db                          *gorm.DB                              = conn.SetupDatabaseConnection()
	kategoriArtikelRepository   repository.KategoriArtikelRepository = repository.NewKategoriArtikelRepository(db)
	KategoriArtikelService      service.KategoriArtikelService       = service.NewKategoriArtikelService(kategoriArtikelRepository)
	kategoriArtikelController   controller.KategoriArtikelController = controller.NewKategoriArtikelController(KategoriArtikelService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	kategoriArtikelRoutes := r.Group("/api/kategori-artikel")
	{
		kategoriArtikelRoutes.GET("/", kategoriArtikelController.All)
		kategoriArtikelRoutes.POST("/", kategoriArtikelController.Insert)
		kategoriArtikelRoutes.GET("/:id", kategoriArtikelController.FindByID)
		kategoriArtikelRoutes.PUT("/:id", kategoriArtikelController.Update)
		kategoriArtikelRoutes.DELETE("/:id", kategoriArtikelController.Delete)
	}
	r.Run("192.168.154.117:8083")
}
