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

func (p *Payroll) Validate() map[string]string {
	var errorMessages = make(map[string]string)

	if p.Id_Payment == 0 {
		errorMessages["Id_Payment"] = "Payment ID is required"
	}
	if p.Id_Employee == 0 {
		errorMessages["Id_Employee"] = "Employee ID is required"
	}
	if p.Full_Name == "" {
		errorMessages["Full_Name"] = "Full Name is required"
	}
	if p.Job_Title == "" {
		errorMessages["Job_Title"] = "Job Title is required"
	}
	if p.Payment_Period == "" {
		errorMessages["Payment_Period"] = "Payment Period is required"
	}
	if p.Payment_Date == "" {
		errorMessages["Payment_Date"] = "Payment Date is required"
	}
	if p.Payment_Status == "" {
		errorMessages["Payment_Status"] = "Payment Status is required"
	}
	if p.Basic_Salary == 0 {
		errorMessages["Basic_Salary"] = "Basic Salary is required"
	}
	if p.Bpjs == 0 {
		errorMessages["Bpjs"] = "BPJS is required"
	}
	if p.Tax == 0 {
		errorMessages["Tax"] = "Tax is required"
	}
	if p.Total_Salary == 0 {
		errorMessages["Total_Salary"] = "Total Salary is required"
	}

	return errorMessages
}
