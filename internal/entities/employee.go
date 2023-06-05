package entities

import "time"

type Employee struct {
	ID                     int32
	Name                   string
	HourlyRate             int32
	MonthlySalary          int32
	CommissionRate         int32
	ReceivingType          ReceivingType
	Address                string
	BankAccountInformation string
	UnionDue               int32
	Payday                 time.Time
}

type ReceivingType int32

const (
	ReceivingTypeMailCheck ReceivingType = iota
	ReceivingTypeHandOverCheck
	ReceivingTypeDeposit
)
