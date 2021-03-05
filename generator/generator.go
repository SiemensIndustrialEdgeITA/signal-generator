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
	Stop()
}

type gentype int

const (
	LINEAR gentype = iota
	SINE
)

type datapoint struct {
	key string
	ts  time.Time
	val float64
}

func NewGenerator(gtype gentype, c *Config) (Generator, error) {
	t := time.NewTicker(c.SampleRate)
	q := make(chan struct{})             // Unbuffered
	o := make(chan datapoint, c.Bufflen) // Buffered

	switch gtype {
	case LINEAR:
		return &LinearGenerator{
			interval: c.SampleRate,
			value:    0,
			coeff:    1,
			minVal:   c.MinVal,
			maxVal:   c.MaxVal,
			ticker:   t,
			quit:     q,
			out:      o,
		}, nil
	case SINE:
		return &SineGenerator{
			interval: c.SampleRate,
			value:    0,
			ticker:   t,
			quit:     q,
			out:      o,
		}, nil

	}
	return nil, fmt.Errorf("could not create generator type: %d", gtype)
}

type LinearGenerator struct {
	log      *logger.Logger
	interval time.Duration
	value    float64
	coeff    float64
	minVal   int
	maxVal   int
	ticker   *time.Ticker
	quit     chan struct{}
	out      chan datapoint
}

type SineGenerator struct {
	log      *logger.Logger
	interval time.Duration
	value    float64
	ampl     int
	period   int
	ticker   *time.Ticker
	quit     chan struct{}
	out      chan datapoint
}

func (l *LinearGenerator) Start() {
	fmt.Println("starting linear generation")
	fmt.Println("interval:", l.interval)
	for {
		select {
		case t := <-l.ticker.C:
			fmt.Println("key:", "linear")
			fmt.Println("ts:", t)
			fmt.Println("value:", l.value)
			intervalsec := float64(l.interval) / 1000000000
			l.value = l.value + (l.coeff * intervalsec)
			msg := datapoint{
				key: "linear",
				ts:  t,
				val: l.value,
			}
			l.out <- msg
		case <-l.quit:
			fmt.Println("received close")
			l.ticker.Stop()
			return
		}
	}
}

func (l *LinearGenerator) Stop() {
	close(l.quit)
}

func (s *SineGenerator) Start() {
}

func (s *SineGenerator) Stop() {
}
