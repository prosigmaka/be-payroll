package router

import (
	// "fmt"
	"be-payroll/api/middleware"
	"be-payroll/internal/approval_panel"
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

	// Approval Panel
	approvalRepo := approval_panel.NewRepository(db)
	approvalService := approval_panel.NewService(approvalRepo)
	approvalHandler := approval_panel.NewApprovalHandler(approvalService)

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

			// Payroll Panel Routes API
			admin.GET("/payroll-panel", middleware.CORSMiddleware(), payrollHandler.GetAllList)
			admin.GET("/payroll-panel/detail/:id", middleware.CORSMiddleware(), payrollHandler.GetListById)
			admin.POST("/payroll-panel/post", middleware.CORSMiddleware(), payrollHandler.Create)
			admin.PUT("/payroll-panel/update/:id", middleware.CORSMiddleware(), payrollHandler.Update)
			admin.DELETE("/payroll-panel/remove/:id", middleware.CORSMiddleware(), payrollHandler.Delete)
			// admin.DELETE("/payroll-panel/", middleware.CORSMiddleware(), payrollHandler.DeleteAllList)
			admin.GET("/payroll-panel/query", middleware.CORSMiddleware(), payrollHandler.QueryList)

			// Implementing Go Routine for Payroll Panel
			admin.POST("/payroll-panel/post/list", middleware.CORSMiddleware(), payrollHandler.CreateList)

			// Approval Routes API
			admin.GET("/approval", middleware.CORSMiddleware(), approvalHandler.GetAllList)
			admin.GET("/approval/detail/:id", middleware.CORSMiddleware(), approvalHandler.GetListById)
			admin.POST("/approval/post", middleware.CORSMiddleware(), approvalHandler.Create)
			admin.PUT("/approval/update/:id", middleware.CORSMiddleware(), approvalHandler.Update)
			admin.DELETE("/approval/remove/:id", middleware.CORSMiddleware(), approvalHandler.Delete)
			// admin.DELETE("/approval/", middleware.CORSMiddleware(), approvalHandler.DeleteAllList)
			admin.GET("/approval/query", middleware.CORSMiddleware(), approvalHandler.Query)

			// Implementing Go Routine for Approval
			admin.POST("/approval/post/list", middleware.CORSMiddleware(), approvalHandler.CreateList)
		}
	}

	return router
}
