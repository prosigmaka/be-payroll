package router

import (
	// "fmt"
	"be-payroll/api/middleware"
	"be-payroll/internal/payroll_panel"

	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "github.com/gorilla/mux"
)

func Route(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Payroll Panel
	payrollRepo := payroll_panel.NewRepository(db)
	payrollService := payroll_panel.NewService(payrollRepo)
	payrollHandler := payroll_panel.NewPayrollHandler(payrollService)

	v1 := router.Group("/v1")
	{
		admin := v1.Group("/api/admin")
		{
			// admin.GET("/payroll-panel", payrollHandler.GetAllList)
			// admin.GET("/payroll-panel/:id", payrollHandler.GetListById)
			// admin.POST("/payroll-panel", payrollHandler.CreateList)
			// admin.PUT("/payroll-panel/:id", payrollHandler.UpdateList)
			// admin.DELETE("/payroll-panel/:id", payrollHandler.DeleteList)
			// // admin.DELETE("/payroll-panel/", payrollHandler.DeleteAllList)
			// admin.GET("/payroll-panel/query", payrollHandler.QueryList)

			admin.GET("/payroll-panel", middleware.CORSMiddleware(), payrollHandler.GetAllList)
			admin.GET("/payroll-panel/detail/:id", middleware.CORSMiddleware(), payrollHandler.GetListById)
			admin.POST("/payroll-panel/post", middleware.CORSMiddleware(), payrollHandler.CreateList)
			admin.PUT("/payroll-panel/update/:id", middleware.CORSMiddleware(), payrollHandler.UpdateList)
			admin.DELETE("/payroll-panel/remove/:id", middleware.CORSMiddleware(), payrollHandler.DeleteList)
			// admin.DELETE("/payroll-panel/", middleware.CORSMiddleware(), payrollHandler.DeleteAllList)
			admin.GET("/payroll-panel/query", middleware.CORSMiddleware(), payrollHandler.QueryList)
		}
	}

	return router
}
