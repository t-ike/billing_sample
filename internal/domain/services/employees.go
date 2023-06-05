package services

import (
	"awesomeProject/internal/entities"
	"errors"
)

type EmployeeService struct {
	employeeRepo EmployeeRepostiory
}

func (e *EmployeeService) CreateEmployee(id int32, name string, address string, hourlyRate, monthlySalary, commissionRate int32) (*entities.Employee, error) {
	// todo: validation を行う

	employee := &entities.Employee{
		ID:      id,
		Name:    name,
		Address: address,
	}
	employeeRepo.CreateEmployee(employee)
	return employee, nil
}

func (e *EmployeeService) DeleteEmployee(id int32) error {
	if IDInvalid(id) {
		return errors.New("invalid Employee id")
	}

	employeeRepo.DeleteEmploy(id)
	return nil
}
