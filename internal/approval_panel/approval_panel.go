package approval_panel

import (
	"time"
)

type Approval struct {
	ID              int
	Id_Request      int
	Id_Employee     int
	Full_Name       string
	Leave_Type      string
	Job_Title       string
	Division        string
	Description     string
	Address         string
	Start_Date      string
	End_Date        string
	Approval_Status string
	Created_At      time.Time
	Updated_At      time.Time
}

// type Leave struct {
// 	Leave_Type string
// 	Count      int
// }
