package service

import (
	"employee-service-gin/model"
	"employee-service-gin/repository"
)

type Service interface {
	AddEmployee(e *model.Employee)
	GetEmployees() []*model.Employee
	DeleteEmployee(id int) bool
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) AddEmployee(e *model.Employee) {
	repository.AddEmployeeRepo(e)
}

func (s *service) GetEmployees() []*model.Employee {
	return repository.DisplayEmployeesRepo()
}

func (s *service) DeleteEmployee(id int) bool {
	return repository.DeleteEmployeeRepo(id)
}
