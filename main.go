package main

import (
	"log"
)

var buffer []int

func main() {
	rw, err := InitRenderWindow("Goscilloscope", 1440, 900)
	if err != nil {
		log.Fatalf("Error initializing render window: %s", err)
	}
	go rw.BeginRendering()
	dataChan := make(chan int)
	go ReadSerial(dataChan)
	buffer = make([]int, 0)
	receiveData(dataChan)
}

func receiveData(data chan int) {
	val := <-data
	if len(buffer) > 1440 {
		buffer = buffer[1:]
	}
	buffer = append(buffer, val)
	newData(buffer)
	receiveData(data)
}
