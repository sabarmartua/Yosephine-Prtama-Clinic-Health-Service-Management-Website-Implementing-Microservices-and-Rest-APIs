package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sabarmartua/FAQ/conn"
	"github.com/sabarmartua/FAQ/controller"
	"github.com/sabarmartua/FAQ/repository"
	"github.com/sabarmartua/FAQ/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = conn.SetupDatabaseConnection()
	faqRepository     repository.FAQRepository    = repository.NewFAQRepository(db)
	faqService        service.FAQService          = service.NewFAQService(faqRepository)
	faqController     controller.FAQController    = controller.NewFAQController(faqService)
)

// membuat variable db dengan nilai setup database connection
func main() {
	defer conn.CloseDatabaseConnection(db)
	r := gin.Default()

	faqRoutes := r.Group("/api/faq")
	{
		faqRoutes.GET("/all", faqController.All)
		faqRoutes.POST("/create", faqController.Insert)
		faqRoutes.GET("/:id", faqController.FindByID)
		faqRoutes.PUT("/update/:id", faqController.Update)
		faqRoutes.DELETE("/delete/:id", faqController.Delete)
	}
	r.Run("192.168.154.117:8086")
}
