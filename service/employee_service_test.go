package service

import (
	"employee-service-gin/model"
	"testing"
)

func resetRepo(s Service) {
	emps := s.GetEmployees()
	for _, emp := range emps {
		s.DeleteEmployee(emp.ID)
	}
}

func TestAddEmployee(t *testing.T) {
	s := NewService()
	resetRepo(s)

	emp := &model.Employee{ID: 1, Name: "Alice", Location: "Singapore"}
	s.AddEmployee(emp)

	list := s.GetEmployees()
	if len(list) != 1 {
		t.Errorf("Expected 1 employee, got %d", len(list))
	}
}

func TestDisplayEmployees(t *testing.T) {
	s := NewService()
	resetRepo(s)

	s.AddEmployee(&model.Employee{ID: 2, Name: "Bob", Location: "Tokyo"})
	s.AddEmployee(&model.Employee{ID: 3, Name: "Carol", Location: "Paris"})

	list := s.GetEmployees()
	if len(list) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(list))
	}
}

func TestDeleteEmployee(t *testing.T) {
	s := NewService()
	resetRepo(s)

	s.AddEmployee(&model.Employee{ID: 2, Name: "Bob", Location: "Tokyo"})
	deleted := s.DeleteEmployee(2)
	if !deleted {
		t.Errorf("Expected deletion to succeed")
	}

	list := s.GetEmployees()
	if len(list) != 0 {
		t.Errorf("Expected employee list to be empty after deletion")
	}
}

func TestDeleteEmployee_NotExist(t *testing.T) {
	s := NewService()
	resetRepo(s)

	result := s.DeleteEmployee(404)
	if result {
		t.Errorf("Expected deletion to fail, got success")
	}
}
