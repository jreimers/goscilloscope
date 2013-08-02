package main

import (
	"time"
)

type Clock struct {
	Tick            chan bool
	framesPerSecond float64
	Running         bool
	ticker          *time.Ticker
	quit            chan struct{}
}

func NewClock() *Clock {
	c := new(Clock)
	c.Tick = make(chan bool, 1)
	c.SetFPS(60)
	c.quit = make(chan struct{})
	return c
}

func (c *Clock) Start() {
	c.Running = true
	go c.loop()
}

func (c *Clock) Stop() {
	c.Running = false
	close(c.quit)
}

func (c *Clock) SetFPS(fps float64) {
	c.framesPerSecond = fps
	c.ticker = time.NewTicker(time.Duration(int64((1.0 / fps) * float64(time.Second))))
}

func (c *Clock) GetFPS() float64 {
	return c.framesPerSecond
}

func (c Clock) loop() {
	for {
		select {
		case <-c.ticker.C:
			c.Tick <- true
		case <-c.quit:
			println("clock quit")
			c.ticker.Stop()
			return
		}
	}
}
