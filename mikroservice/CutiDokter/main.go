package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/CutiDokter/conn"
	"github.com/sabarmartua/CutiDokter/controller"
	"github.com/sabarmartua/CutiDokter/repository"
	"github.com/sabarmartua/CutiDokter/service"
	"gorm.io/gorm"
)

var (
	db                   *gorm.DB                        = conn.SetupDatabaseConnection()
	cutiDokterRepository repository.CutiDokterRepository = repository.NewCutiDokterRepository(db)
	CutiDokterService    service.CutiDokterService       = service.NewCutiDokterService(cutiDokterRepository)
	CutiDokterController controller.CutiDokterController = controller.NewCutiDokterController(CutiDokterService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	CutiDokterRoutes := r.Group("/api/cutidokter")
	{
		CutiDokterRoutes.GET("/all", CutiDokterController.All)
		CutiDokterRoutes.POST("/create", CutiDokterController.Insert)
		CutiDokterRoutes.GET("/:id", CutiDokterController.FindByID)
		CutiDokterRoutes.PUT("/update/:id", CutiDokterController.Update)
		CutiDokterRoutes.DELETE("/delete/:id", CutiDokterController.Delete)
	}
	r.Run("192.168.154.117:8085")
}
