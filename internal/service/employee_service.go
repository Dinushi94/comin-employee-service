// internal/service/employee_service.go
package service

import (
	"errors"
	"fmt"

	"github.com/Axontik/comin-employee-service/internal/domain"
	"github.com/Axontik/comin-employee-service/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeService interface {
	Create(req *domain.CreateEmployeeRequest) (*domain.Employee, error)
	GetByID(id uuid.UUID) (*domain.Employee, error)
	GetByUserID(userID uuid.UUID) (*domain.Employee, error)
	Update(id uuid.UUID, req *domain.UpdateEmployeeRequest) (*domain.Employee, error)
	Delete(id uuid.UUID) error
	ListByOrganization(orgID uuid.UUID) ([]domain.Employee, error)
	ListByDepartment(departmentID uuid.UUID) ([]domain.Employee, error)
	ListByManager(managerID uuid.UUID) ([]domain.Employee, error)
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo,
	}
}

func (s *employeeService) Create(req *domain.CreateEmployeeRequest) (*domain.Employee, error) {
	// Check if employee already exists
	existingEmployee, err := s.employeeRepo.GetByUserID(req.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			employee := &domain.Employee{
				UserID:         req.UserID,
				OrganizationID: req.OrganizationID,
				PositionID:     req.PositionID,
				DepartmentID:   req.DepartmentID,
				FirstName:      req.FirstName,
				LastName:       req.LastName,
				Email:          req.Email,
				Phone:          req.Phone,
				DateOfBirth:    req.DateOfBirth,
				HireDate:       req.HireDate,
				EmployeeID:     req.EmployeeID,
				WorkType:       req.WorkType,
				ManagerID:      req.ManagerID,
				Status:         "active",
			}

			if err := s.employeeRepo.Create(employee); err != nil {
				return nil, fmt.Errorf("failed to create employee: %w", err)
			}

			return employee, nil
		}
		return nil, fmt.Errorf("error checking existing employee: %w", err)
	}

	if existingEmployee != nil {
		return nil, errors.New("employee already exists for this user")
	}

	return nil, errors.New("unexpected error")
}

func (s *employeeService) GetByID(id uuid.UUID) (*domain.Employee, error) {
	return s.employeeRepo.GetByID(id)
}

func (s *employeeService) GetByUserID(userID uuid.UUID) (*domain.Employee, error) {
	return s.employeeRepo.GetByUserID(userID)
}

func (s *employeeService) Update(id uuid.UUID, req *domain.UpdateEmployeeRequest) (*domain.Employee, error) {
	employee, err := s.employeeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		employee.FirstName = req.FirstName
	}
	if req.LastName != "" {
		employee.LastName = req.LastName
	}
	if req.Phone != "" {
		employee.Phone = req.Phone
	}
	if req.PositionID != nil {
		employee.PositionID = req.PositionID
	}
	if req.DepartmentID != nil {
		employee.DepartmentID = req.DepartmentID
	}
	if req.ManagerID != nil {
		employee.ManagerID = req.ManagerID
	}
	if req.WorkType != "" {
		employee.WorkType = req.WorkType
	}

	if err := s.employeeRepo.Update(employee); err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *employeeService) Delete(id uuid.UUID) error {
	return s.employeeRepo.Delete(id)
}

func (s *employeeService) ListByOrganization(orgID uuid.UUID) ([]domain.Employee, error) {
	return s.employeeRepo.GetByOrganization(orgID)
}

func (s *employeeService) ListByDepartment(departmentID uuid.UUID) ([]domain.Employee, error) {
	return s.employeeRepo.ListByDepartment(departmentID)
}

func (s *employeeService) ListByManager(managerID uuid.UUID) ([]domain.Employee, error) {
	return s.employeeRepo.ListByManager(managerID)
}
