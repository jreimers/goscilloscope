package main

import (
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
				str := strip(buf)
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

func strip(bytes []byte) string {
	out := make([]byte, 0)
	for _, b := range bytes {
		if b >= 48 && b <= 57 { // only numbers
			out = append(out, b)
		}
	}
	return string(out)
}
