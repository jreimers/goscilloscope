package main

import (
//"log"
)

func main() {
	chan1 := NewChannel("Channel 1")
	chan2 := NewChannel("Channel 2")

	quit := make(chan struct{})

	clock := NewClock()
	clock.SetFPS(30)
	clock.Start()

	s := NewSerialReader("/dev/ttyACM0", quit)
	s.AddChannel(chan1)
	s.AddChannel(chan2)
	s.BeginRead()
	println("postread")
	// rw, err := NewRenderWindow("Goscilloscope", 1024, 768, quit)
	// if err != nil {
	// 	log.Fatalf("Error creating render window: ", err)
	// }
	// rw.AddChannel(chan1)
	// rw.AddChannel(chan2)

	// err = rw.Open()

	// if err != nil {
	// 	log.Fatalf("Error opening render window: ", err)
	// }

	// main loop
	for {
		select {
		case <-clock.Tick:
			println("tick")
			//rw.Render()
		case <-quit:
			//rw.Close()
			break
		}
	}
	println("stopping clock")
	clock.Stop()
}
