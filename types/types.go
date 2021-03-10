package types

import (
	"time"
)

// DataPoint schema for internal message passing and for SimpleSink output
type DataPoint struct {
	Key string
	Ts  time.Time
	Val float64
}
