package payroll_panel

import (
	"time"
)

type Payroll struct {
	ID             int
	Id_Payment     int
	Id_Employee    int
	Full_Name      string
	Job_Title      string
	Payment_Period string
	Payment_Date   string
	Payment_Status string
	Basic_Salary   int
	Bpjs           int
	Tax            float64
	Total_Salary   int
	Created_At     time.Time
	Updated_At     time.Time
}
