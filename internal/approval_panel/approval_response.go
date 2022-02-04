package approval_panel

type ApprovalResponse struct {
	ID              int    `json:"id"`
	Id_Request      int    `json:"id_request"`
	Id_Employee     int    `json:"id_employee"`
	Full_Name       string `json:"full_name"`
	Leave_Type      string `json:"leave_type"`
	Job_Title       string `json:"job_title"`
	Division        string `json:"division"`
	Description     string `json:"description"`
	Address         string `json:"address"`
	Start_Date      string `json:"start_date"`
	End_Date        string `json:"end_date"`
	Approval_Status string `json:"approval_status"`
}

type CountLeave struct {
	Leave_Type string `json:"leave_type"`
	Count      int    `json:"count"`
}
