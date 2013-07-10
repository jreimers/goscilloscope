package main

import (
	"bytes"
	"github.com/tarm/goserial"
	"log"
	"strconv"
)

func ReadSerial(data chan int) {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalf("Error opening serial port: %s", err)
	}
	buf := make([]byte, 128)
	running := true
	for running {

		in := make([]byte, 1)
		n, err := s.Read(in)
		if err == nil && n == 1 {
			if string(in) == "\n" {
				str := string(bytes.Trim(buf, "\x00"))
				if len(str) > 0 {
					y, _ := strconv.ParseInt(str, 10, 64)
					data <- int(y)
				}
				buf = make([]byte, 128)
			} else {
				buf = append(buf, in[0])
			}
		}
	}
}
