package entities

import "time"

type WorkLog struct {
	id         int32
	employeeID int32
	startTime  time.Time
	endTime    time.Time
}
