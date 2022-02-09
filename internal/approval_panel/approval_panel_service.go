package approval_panel

type Service interface {
	Create(CreateApprovalRequest CreateApprovalRequest) (Approval, error)
	FindAll() ([]Approval, error)
	FindByID(ID int) (Approval, error)
	Update(ID int, UpdateApprovalRequest UpdateApprovalRequest) (Approval, error)
	Delete(ID int) (Approval, error)
	CountAnnualLeave() ([]Approval, error)
	CountLeavePermission() ([]Approval, error)
	CountSickLeave() ([]Approval, error)
	CreateList(CreateApprovalRequest []CreateApprovalRequest) ([]Approval, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(createApprovalRequest CreateApprovalRequest) (Approval, error) {
	id_request, _ := createApprovalRequest.Id_Request.Int64()
	id_employee, _ := createApprovalRequest.Id_Employee.Int64()

	approval := Approval{
		Id_Request:      int(id_request),
		Id_Employee:     int(id_employee),
		Full_Name:       createApprovalRequest.Full_Name,
		Leave_Type:      createApprovalRequest.Leave_Type,
		Job_Title:       createApprovalRequest.Job_Title,
		Division:        createApprovalRequest.Division,
		Description:     createApprovalRequest.Description,
		Address:         createApprovalRequest.Address,
		Start_Date:      createApprovalRequest.Start_Date,
		End_Date:        createApprovalRequest.End_Date,
		Approval_Status: createApprovalRequest.Approval_Status,
	}

	newApproval, err := s.repository.Create(approval)

	return newApproval, err
}

func (s *service) FindAll() ([]Approval, error) {
	approvals, err := s.repository.FindAll()

	return approvals, err
}

func (s *service) FindByID(ID int) (Approval, error) {
	approval, err := s.repository.FindByID(ID)

	return approval, err
}

func (s *service) Update(ID int, updateApprovalRequest UpdateApprovalRequest) (Approval, error) {
	approval, _ := s.repository.FindByID(ID)

	id_request, _ := updateApprovalRequest.Id_Request.Int64()
	id_employee, _ := updateApprovalRequest.Id_Employee.Int64()

	approval.Id_Request = int(id_request)
	approval.Id_Employee = int(id_employee)
	approval.Full_Name = updateApprovalRequest.Full_Name
	approval.Leave_Type = updateApprovalRequest.Leave_Type
	approval.Job_Title = updateApprovalRequest.Job_Title
	approval.Division = updateApprovalRequest.Division
	approval.Description = updateApprovalRequest.Description
	approval.Address = updateApprovalRequest.Address
	approval.Start_Date = updateApprovalRequest.Start_Date
	approval.End_Date = updateApprovalRequest.End_Date
	approval.Approval_Status = updateApprovalRequest.Approval_Status

	newApproval, err := s.repository.Update(approval)

	return newApproval, err
}

func (s *service) Delete(ID int) (Approval, error) {
	approval, _ := s.repository.FindByID(ID)

	newApproval, err := s.repository.Delete(approval)

	return newApproval, err
}

func (s *service) CountAnnualLeave() ([]Approval, error) {
	approvals, err := s.repository.CountAnnualLeave()

	return approvals, err
}

func (s *service) CountLeavePermission() ([]Approval, error) {
	approvals, err := s.repository.CountLeavePermission()

	return approvals, err
}

func (s *service) CountSickLeave() ([]Approval, error) {
	approvals, err := s.repository.CountSickLeave()

	return approvals, err
}

// Go Routine for Form Create List Approval
func addList(start int, end int, createList []CreateApprovalRequest, channel chan Approval, s *service) {
	for i := start; i < end; i++ {
		approvalList, _ := s.Create(createList[i])
		channel <- approvalList
	}
}

func (s *service) CreateList(createApprovalRequest []CreateApprovalRequest) ([]Approval, error) {
	n := len(createApprovalRequest) / 2
	channel := make(chan Approval)

	go addList(0, n, createApprovalRequest, channel, s)

	go addList(n, len(createApprovalRequest), createApprovalRequest, channel, s)

	var approvalList []Approval
	for i := 0; i < len(createApprovalRequest); i++ {
		approvalCreateList := <-channel
		newList := Approval{
			ID:              approvalCreateList.ID,
			Id_Request:      approvalCreateList.Id_Request,
			Id_Employee:     approvalCreateList.Id_Employee,
			Full_Name:       approvalCreateList.Full_Name,
			Leave_Type:      approvalCreateList.Leave_Type,
			Job_Title:       approvalCreateList.Job_Title,
			Division:        approvalCreateList.Division,
			Description:     approvalCreateList.Description,
			Address:         approvalCreateList.Address,
			Start_Date:      approvalCreateList.Start_Date,
			End_Date:        approvalCreateList.End_Date,
			Approval_Status: approvalCreateList.Approval_Status,
		}
		approvalList = append(approvalList, newList)
	}

	return approvalList, nil
}

// Go Routine for Form Create List Approval
