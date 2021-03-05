package generator

import "time"

type Config struct {
	SampleRate time.Duration
	MinVal     int
	MaxVal     int
}
