package payroll_panel

import (
	"encoding/json"
)

type CreatePayrollRequest struct {
	Id_Payment     json.Number `json:"id_payment" binding:"required,number"`
	Id_Employee    json.Number `json:"id_employee" binding:"required,number"`
	Full_Name      string      `json:"full_name" binding:"required"`
	Job_Title      string      `json:"job_title" binding:"required"`
	Payment_Period string      `json:"payment_period" binding:"required"`
	Payment_Date   string      `json:"payment_date" binding:"required"`
	Payment_Status string      `json:"payment_status" binding:"required"`
	Basic_Salary   json.Number `json:"basic_salary" binding:"required,number"`
	Bpjs           json.Number `json:"bpjs" binding:"required,number"`
	Tax            float64     `json:"tax" binding:"required"`
	Total_Salary   json.Number `json:"total_salary" binding:"required,number"`
}

type UpdatePayrollRequest struct {
	Id_Payment     json.Number `json:"id_payment" binding:"required,number"`
	Id_Employee    json.Number `json:"id_employee" binding:"required,number"`
	Full_Name      string      `json:"full_name" binding:"required"`
	Job_Title      string      `json:"job_title" binding:"required"`
	Payment_Period string      `json:"payment_period" binding:"required"`
	Payment_Date   string      `json:"payment_date" binding:"required"`
	Payment_Status string      `json:"payment_status" binding:"required"`
	Basic_Salary   json.Number `json:"basic_salary" binding:"required,number"`
	Bpjs           json.Number `json:"bpjs" binding:"required,number"`
	Tax            float64     `json:"tax" binding:"required"`
	Total_Salary   json.Number `json:"total_salary" binding:"required,number"`
}
