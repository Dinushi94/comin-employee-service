// internal/repository/employee_repository.go
package repository

import (
	"fmt"

	"github.com/Axontik/comin-employee-service/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *domain.Employee) error
	GetByID(id uuid.UUID) (*domain.Employee, error)
	GetByUserID(userID uuid.UUID) (*domain.Employee, error)
	GetByOrganization(orgID uuid.UUID) ([]domain.Employee, error)
	Update(employee *domain.Employee) error
	Delete(id uuid.UUID) error
	ListByDepartment(departmentID uuid.UUID) ([]domain.Employee, error)
	ListByManager(managerID uuid.UUID) ([]domain.Employee, error)
	GetByEmployeeID(orgID uuid.UUID, employeeID string) (*domain.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(employee *domain.Employee) error {
	if employee.ID == uuid.Nil {
		employee.ID = uuid.New()
	}
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetByID(id uuid.UUID) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.Preload("Position").First(&employee, "id = ?", id).Error
	return &employee, err
}

func (r *employeeRepository) GetByUserID(userID uuid.UUID) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.Preload("Position").First(&employee, "id = ?", userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &employee, nil
}

func (r *employeeRepository) GetByOrganization(orgID uuid.UUID) ([]domain.Employee, error) {
	var employees []domain.Employee
	err := r.db.Preload("Position").Where("organization_id = ?", orgID).Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) Update(employee *domain.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Employee{}, "id = ?", id).Error
}

func (r *employeeRepository) ListByDepartment(departmentID uuid.UUID) ([]domain.Employee, error) {
	var employees []domain.Employee
	err := r.db.Preload("Position").Where("department_id = ?", departmentID).Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) ListByManager(managerID uuid.UUID) ([]domain.Employee, error) {
	var employees []domain.Employee
	err := r.db.Preload("Position").Where("manager_id = ?", managerID).Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) GetByEmployeeID(orgID uuid.UUID, employeeID string) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.Preload("Position").
		Where("organization_id = ? AND employee_id = ?", orgID, employeeID).
		First(&employee).Error
	return &employee, err
}
