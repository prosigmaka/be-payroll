package approval_panel

import "gorm.io/gorm"

type Repository interface {
	Create(approval Approval) (Approval, error)
	FindAll() ([]Approval, error)
	FindByID(ID int) (Approval, error)
	Update(approval Approval) (Approval, error)
	Delete(approval Approval) (Approval, error)
	// DeleteAll() ([]Approval, error)
	CountAnnualLeave() ([]Approval, error)
	CountLeavePermission() ([]Approval, error)
	CountSickLeave() ([]Approval, error)
}

var countAnnualLeave int64
var countLeavePermission int64
var CountSickLeave int64

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(approval Approval) (Approval, error) {
	err := r.db.Debug().Create(&approval).Error

	return approval, err
}

func (r *repository) FindAll() ([]Approval, error) {
	var approvals []Approval

	err := r.db.Debug().Order("id desc").Find(&approvals).Error

	return approvals, err
}

func (r *repository) FindByID(ID int) (Approval, error) {
	var approval Approval

	// err := r.db.Debug().Where("id = ?", ID).(&approval).Error
	err := r.db.Debug().Find(&approval, ID).Error

	return approval, err
}

func (r *repository) Update(approval Approval) (Approval, error) {
	err := r.db.Debug().Save(&approval).Error

	return approval, err
}

func (r *repository) Delete(approval Approval) (Approval, error) {
	err := r.db.Debug().Delete(&approval).Error

	return approval, err
}

func (r *repository) CountAnnualLeave() ([]Approval, error) {
	var approvals []Approval

	err := r.db.Debug().Where("leave_type = ?", "Annual").Find(&approvals).Count(&countAnnualLeave).Error

	return approvals, err
}

func (r *repository) CountLeavePermission() ([]Approval, error) {
	var approvals []Approval

	err := r.db.Debug().Where("leave_type = ?", "Permission").Find(&approvals).Count(&countLeavePermission).Error

	return approvals, err
}

func (r *repository) CountSickLeave() ([]Approval, error) {
	var approvals []Approval

	err := r.db.Debug().Where("leave_type = ?", "Sick").Find(&approvals).Count(&CountSickLeave).Error

	return approvals, err
}
