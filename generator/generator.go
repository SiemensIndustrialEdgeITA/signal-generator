package generator

import (
	//	"context"
	//	"errors"
	//	"io/ioutil"
	//	"os"
	//	"path/filepath"
	//	"strings"
	//	"errors"
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
)

type Generator interface {
	Start()
}

type gentype int

const (
	LINEAR gentype = iota
	SINE
)

func NewGenerator(gtype gentype, c *Config) (Generator, error) {
	t := time.NewTicker(c.SampleRate * time.Millisecond)

	switch gtype {
	case LINEAR:
		return &LinearGenerator{
			interval: c.SampleRate,
			minVal:   c.MinVal,
			maxVal:   c.MaxVal,
			ticker:   t,
		}, nil
	case SINE:
		return &SineGenerator{
			interval: c.SampleRate,
			ticker:   t,
		}, nil

	}
	return nil, fmt.Errorf("could not create generator type: %d", gtype)
}

type LinearGenerator struct {
	log      *logger.Logger
	interval time.Duration
	minVal   int
	maxVal   int
	ticker   *time.Ticker
}

type SineGenerator struct {
	log      *logger.Logger
	interval time.Duration
	ampl     int
	period   int
	ticker   *time.Ticker
}

func (l *LinearGenerator) Start() {
}

func (s *SineGenerator) Start() {
}
