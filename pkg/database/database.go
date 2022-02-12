package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	// DB_HOST = "payroll-postgres" // run in docker
	DB_HOST     = "172.20.0.2" // run in docker
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASS     = "admin"
	DB_DATABASE = "payroll"
)

// var DB = &gorm.DB{}

func InitDB() (*gorm.DB, error) {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	// DB_USER := os.Getenv("DB_USER")
	// DB_PASS := os.Getenv("DB_PASS")
	// DB_DATABASE := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASS, DB_DATABASE, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("Connected to database")

	// db.AutoMigrate(&payroll_panel.Payroll{})

	// // Repository Payroll
	// payrollRepository := payroll_panel.NewRepository(db)
	// payrollService := payroll_panel.NewService(payrollRepository)

	// payrollRequest := payroll_panel.CreatePayrollRequest{
	// 	Id_Payment:     "1001",
	// 	Id_Employee:    "1001",
	// 	Full_Name:      "Karyawan A",
	// 	Job_Title:      "Full Stack Developer",
	// 	Payment_Period: "Monthly",
	// 	Payment_Date:   "31 December 2021",
	// 	Payment_Status: "Paid",
	// 	Basic_Salary:   "1000000",
	// 	Bpjs:           "100000",
	// 	Tax:            0.05,
	// 	Total_Salary:   "850000",
	// }

	// payrollService.Create(payrollRequest)

	// db.AutoMigrate(&approval_panel.Approval{})

	// // Repository Approval
	// approvalRepository := approval_panel.NewRepository(db)
	// approvalService := approval_panel.NewService(approvalRepository)

	// approvalRequest := approval_panel.CreateApprovalRequest{
	// 	Id_Request:      "1001",
	// 	Id_Employee:     "1001",
	// 	Full_Name:       "Karyawan A",
	// 	Leave_Type:      "Sick",
	// 	Job_Title:       "Full Stack Developer",
	// 	Division:        "Developer",
	// 	Description:     "Sick Leave",
	// 	Address:         "Jakarta",
	// 	Start_Date:      "2021-12-31",
	// 	End_Date:        "2022-01-01",
	// 	Approval_Status: "Pending",
	// }
	// approvalService.Create(approvalRequest)

	return db, nil
}
