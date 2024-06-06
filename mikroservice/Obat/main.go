package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sabarmartua/Obat/conn"
    "github.com/sabarmartua/Obat/controller"
    "github.com/sabarmartua/Obat/repository"
    "github.com/sabarmartua/Obat/service"
    "gorm.io/gorm"
)

var (
    db              *gorm.DB                  = conn.SetupDatabaseConnection()
    obatRepository  repository.ObatRepository = repository.NewObatRepository(db)
    obatService     service.ObatService       = service.NewObatService(obatRepository)
    obatController controller.ObatController = controller.NewObatController(obatService)
)

func main() {
    defer conn.CloseDatabaseConnection(db)
    r := gin.Default()

    obatRoutes := r.Group("/api/obat")
    {
        obatRoutes.GET("/all", obatController.All)
        obatRoutes.POST("/create", obatController.Insert)
        obatRoutes.GET("/:id", obatController.FindByID)
        obatRoutes.PUT("/update/:id", obatController.Update)
        obatRoutes.DELETE("/delete/:id", obatController.Delete)
    }

    r.Run("192.168.154.117:8090")
}
