package services

import (
	"awesomeProject/internal/entities"
	"errors"
	"time"
)

type TimeCardService struct {
	timeCardRepo TimeCardRepostiory
	employeeRepo EmployeeRepository
}

func (r *TimeCardService) CreateTimeCard(employeeID int32, startTime time.Time, hours float32) (*entities.TimeCard, error) {
	// todo: id validation
	employee := r.employeeRepo.FindByID(employeeID)

	if !employee.IsHourlyWorking {
		return nil, errors.New()
	}

	timeCard := &entities.TimeCard{}
	timeCardRepo.CreateTimeCard(timeCard)
	return timeCard, nil
}
