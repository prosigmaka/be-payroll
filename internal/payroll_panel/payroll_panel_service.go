package payroll_panel

type Service interface {
	Create(CreatePayrollRequest CreatePayrollRequest) (Payroll, error)
	FindAll() ([]Payroll, error)
	FindByID(ID int) (Payroll, error)
	Update(ID int, UpdatePayrollRequest UpdatePayrollRequest) (Payroll, error)
	Delete(ID int) (Payroll, error)
	// DeleteAll() ([]Payroll, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(createPayrollRequest CreatePayrollRequest) (Payroll, error) {
	id_payment, _ := createPayrollRequest.Id_Payment.Int64()
	id_employee, _ := createPayrollRequest.Id_Employee.Int64()
	basic_salary, _ := createPayrollRequest.Basic_Salary.Int64()
	bpjs, _ := createPayrollRequest.Bpjs.Int64()
	// tax, _ := createPayrollRequest.Tax.Float64()
	total_salary, _ := createPayrollRequest.Total_Salary.Int64()

	payroll := Payroll{
		Id_Payment:     int(id_payment),
		Id_Employee:    int(id_employee),
		Full_Name:      createPayrollRequest.Full_Name,
		Job_Title:      createPayrollRequest.Job_Title,
		Payment_Period: createPayrollRequest.Payment_Period,
		Payment_Date:   createPayrollRequest.Payment_Date,
		Payment_Status: createPayrollRequest.Payment_Status,
		Basic_Salary:   int(basic_salary),
		Bpjs:           int(bpjs),
		Tax:            createPayrollRequest.Tax,
		Total_Salary:   int(total_salary),
	}

	newPayroll, err := s.repository.Create(payroll)

	return newPayroll, err
}

func (s *service) FindAll() ([]Payroll, error) {
	payrolls, err := s.repository.FindAll()

	return payrolls, err
}

func (s *service) FindByID(ID int) (Payroll, error) {
	payroll, err := s.repository.FindByID(ID)

	return payroll, err
}

func (s *service) Update(ID int, updatePayrollRequest UpdatePayrollRequest) (Payroll, error) {
	payroll, _ := s.repository.FindByID(ID)

	id_payment, _ := updatePayrollRequest.Id_Payment.Int64()
	id_employee, _ := updatePayrollRequest.Id_Employee.Int64()
	basic_salary, _ := updatePayrollRequest.Basic_Salary.Int64()
	bpjs, _ := updatePayrollRequest.Bpjs.Int64()
	// tax, _ := updatePayrollRequest.Tax.Float64()
	total_salary, _ := updatePayrollRequest.Total_Salary.Int64()

	payroll.Id_Payment = int(id_payment)
	payroll.Id_Employee = int(id_employee)
	payroll.Full_Name = updatePayrollRequest.Full_Name
	payroll.Job_Title = updatePayrollRequest.Job_Title
	payroll.Payment_Period = updatePayrollRequest.Payment_Period
	payroll.Payment_Date = updatePayrollRequest.Payment_Date
	payroll.Payment_Status = updatePayrollRequest.Payment_Status
	payroll.Basic_Salary = int(basic_salary)
	payroll.Bpjs = int(bpjs)
	payroll.Tax = updatePayrollRequest.Tax
	payroll.Total_Salary = int(total_salary)

	newPayroll, err := s.repository.Update(payroll)

	return newPayroll, err
}

func (s *service) Delete(ID int) (Payroll, error) {
	payroll, _ := s.repository.FindByID(ID)

	newPayroll, err := s.repository.Delete(payroll)

	return newPayroll, err
}

// func (s *service) DeleteAll() ([]Payroll, error) {
// 	payrolls, err := s.repository.DeleteAll()

// 	return payrolls, err
// }
