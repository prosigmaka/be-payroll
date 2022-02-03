package payroll_panel

import "gorm.io/gorm"

type Repository interface {
	Create(payroll Payroll) (Payroll, error)
	FindAll() ([]Payroll, error)
	FindByID(ID int) (Payroll, error)
	Update(payroll Payroll) (Payroll, error)
	Delete(payroll Payroll) (Payroll, error)
	// DeleteAll() ([]Payroll, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(payroll Payroll) (Payroll, error) {
	err := r.db.Debug().Create(&payroll).Error

	return payroll, err
}

func (r *repository) FindAll() ([]Payroll, error) {
	var payrolls []Payroll

	err := r.db.Debug().Order("id desc").Find(&payrolls).Error

	return payrolls, err
}

func (r *repository) FindByID(ID int) (Payroll, error) {
	var payroll Payroll

	// err := r.db.Debug().Where("id = ?", ID).(&payroll).Error
	err := r.db.Debug().Find(&payroll, ID).Error

	return payroll, err
}

func (r *repository) Update(payroll Payroll) (Payroll, error) {
	err := r.db.Debug().Save(&payroll).Error

	return payroll, err
}

func (r *repository) Delete(payroll Payroll) (Payroll, error) {
	err := r.db.Debug().Delete(&payroll).Error

	return payroll, err
}

// func (r *repository) DeleteAll() ([]Payroll, error) {
// 	var payrolls []Payroll

// 	err := r.db.Debug().Delete(&payrolls).Error

// 	return payrolls, err
// }
