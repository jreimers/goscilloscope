package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	chan1 := NewChannel("Channel 1")
	chan1.SetColor(0.0, 1.0, 0.0)
	chan2 := NewChannel("Channel 2")
	chan2.SetColor(0.0, 0.0, 1.0)
	chan3 := NewChannel("Channel 3")
	chan3.SetColor(1.0, 0.0, 0.0)

	quit := make(chan struct{})

	clock := NewClock()
	clock.SetFPS(60)
	clock.Start()

	s := NewSerialReader("/dev/ttyACM0", quit)
	s.AddChannel(chan1)
	s.AddChannel(chan2)
	s.AddChannel(chan3)
	s.BeginRead()

	rw, err := NewRenderWindow("Goscilloscope", 1024, 768, quit)

	if err != nil {
		log.Fatalf("Error creating render window: ", err)
	}
	rw.AddChannel(chan1)
	rw.AddChannel(chan2)
	rw.AddChannel(chan3)

	err = rw.Open()

	if err != nil {
		log.Fatalf("Error opening render window: ", err)
	}

	// main loop
	for {
		select {
		case <-clock.Tick:
			rw.Render()
		case <-quit:
			rw.Close()
			break
		}
	}
	println("stopping clock")
	clock.Stop()
}
