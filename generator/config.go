package generator

import "time"

type Config struct {
	SampleRate time.Duration
	Bufflen    int
	MinVal     int
	MaxVal     int
}
