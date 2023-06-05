package entities

import "time"

type UnionServiceFeeLog struct {
	id         int32
	employeeID int32
	date       time.Time
	price      int32
}
