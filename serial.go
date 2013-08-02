package main

import (
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"runtime"
	"strconv"
	"strings"
)

type SerialReader struct {
	quit     chan struct{}
	reader   io.ReadWriteCloser
	Device   string
	Channels []*Channel
}

func NewSerialReader(device string, quit chan struct{}) *SerialReader {
	sr := new(SerialReader)
	sr.Device = device
	sr.quit = quit
	return sr
}

func (r *SerialReader) AddChannel(c *Channel) {
	r.Channels = append(r.Channels, c)
}

func (r *SerialReader) BeginRead() {
	go func() {
		defer r.endRead()
		r.setupSerial()
		buf := make([]byte, 128)
		for {
			in := make([]byte, 1)
			n, err := r.reader.Read(in)
			if err == nil && n == 1 {
				if string(in) == "\n" {
					str := strip(buf)
					if len(str) > 0 {
						values := strings.Split(str, ":")
						if len(values) == len(r.Channels) {
							for i, c := range r.Channels {
								y, _ := strconv.ParseInt(values[i], 10, 64)
								// overwrite the last value in the channel
								if len(c.Buffer) > 0 {
									<-c.Buffer
								}
								c.Buffer <- y
							}
						}
					}
					buf = make([]byte, 128)
				} else {
					buf = append(buf, in[0])
				}
			}
			select {
			case <-r.quit:
				return
			default:
				runtime.Gosched()
				continue
			}
		}
	}()
}
func (r *SerialReader) setupSerial() {
	fmt.Println("Connecting to device:", r.Device)
	c := &serial.Config{Name: r.Device, Baud: 9600}
	reader, err := serial.OpenPort(c)
	if err != nil {
		log.Fatalf("Error opening serial port: %s", err)
	}
	r.reader = reader
}
func (r SerialReader) endRead() {
	fmt.Println("End Read")
}

func strip(bytes []byte) string {
	out := make([]byte, 0)
	for _, b := range bytes {
		if b >= 48 && b <= 58 { // only numbers and colon ":"
			out = append(out, b)
		}
	}
	return string(out)
}
