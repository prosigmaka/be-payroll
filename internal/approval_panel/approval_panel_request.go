package approval_panel

import "encoding/json"

type CreateApprovalRequest struct {
	Id_Request      json.Number `json:"id_request" binding:"required,number"`
	Id_Employee     json.Number `json:"id_employee" binding:"required,number"`
	Full_Name       string      `json:"full_name" binding:"required"`
	Leave_Type      string      `json:"leave_type" binding:"required"`
	Job_Title       string      `json:"job_title" binding:"required"`
	Division        string      `json:"division" binding:"required"`
	Description     string      `json:"description" binding:"required"`
	Address         string      `json:"address" binding:"required"`
	Start_Date      string      `json:"start_date" binding:"required"`
	End_Date        string      `json:"end_date" binding:"required"`
	Approval_Status string      `json:"approval_status" binding:"required"`
}

type UpdateApprovalRequest struct {
	Id_Request      json.Number `json:"id_request" binding:"required,number"`
	Id_Employee     json.Number `json:"id_employee" binding:"required,number"`
	Full_Name       string      `json:"full_name" binding:"required"`
	Leave_Type      string      `json:"leave_type" binding:"required"`
	Job_Title       string      `json:"job_title" binding:"required"`
	Division        string      `json:"division" binding:"required"`
	Description     string      `json:"description" binding:"required"`
	Address         string      `json:"address" binding:"required"`
	Start_Date      string      `json:"start_date" binding:"required"`
	End_Date        string      `json:"end_date" binding:"required"`
	Approval_Status string      `json:"approval_status" binding:"required"`
}
