package service

import (
	"github.com/dqtu39/go-simple-server/internal/models"
	"github.com/dqtu39/go-simple-server/internal/repository"
)

type EmployeeService interface {
	GetAllEmployees(offset int, limit int) ([]models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	AddEmployee(employee models.Employee) (int64, error)
	UpdateEmployee(id int, employee models.Employee) (int64, error)
	DeleteEmployee(id int) (int64, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllEmployees(offset int, limit int) ([]models.Employee, error) {
	return s.repo.GetAll(offset, limit)
}

func (s *employeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) AddEmployee(employee models.Employee) (int64, error) {
	return s.repo.Add(employee)
}

func (s *employeeService) UpdateEmployee(id int, employee models.Employee) (int64, error) {
	return s.repo.Update(id, employee)
}

func (s *employeeService) DeleteEmployee(id int) (int64, error) {
	return s.repo.Delete(id)
}
