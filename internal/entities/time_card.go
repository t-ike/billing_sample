package entities

import "time"

type TimeCard struct {
	id         int32
	employeeID int32
	startTime  time.Time
	hours      float32
}
