package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "admin"
	DB_NAME = "payroll"
)

// var DB = &gorm.DB{}

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)

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

	return db, nil
}
