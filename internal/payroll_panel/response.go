package payroll_panel

type PayrollResponse struct {
	ID             int     `json:"id"`
	Id_Payment     int     `json:"id_payment"`
	Id_Employee    int     `json:"id_employee"`
	Full_Name      string  `json:"full_name"`
	Job_Title      string  `json:"job_title"`
	Payment_Period string  `json:"payment_period"`
	Payment_Date   string  `json:"payment_date"`
	Payment_Status string  `json:"payment_status"`
	Basic_Salary   int     `json:"basic_salary"`
	Bpjs           int     `json:"bpjs"`
	Tax            float64 `json:"tax"`
	Total_Salary   int     `json:"total_salary"`
}
