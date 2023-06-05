package entities

import "time"

type Employee struct {
	id                     int32
	hourlyWage             int32
	monthlySalary          int32
	bonus                  int32
	receivingType          ReceivingType
	address                string
	bankAccountInformation string
	unionDue               int32
	payday                 time.Time
}

type ReceivingType int32

const (
	ReceivingTypeMailCheck ReceivingType = iota
	ReceivingTypeHandOverCheck
	ReceivingTypeDeposit
)
