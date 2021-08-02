package model

import "time"

type Measurement struct {
	ID          uint
	Temperature float64
	Humidity    float64
	Date        time.Time
}
