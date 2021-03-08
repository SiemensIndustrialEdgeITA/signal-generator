package types

import (
	"time"
)

type DataPoint struct {
	Key string
	Ts  time.Time
	Val float64
}
