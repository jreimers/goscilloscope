package main

type Channel struct {
	Buffer chan int64
	Name   string
}

func NewChannel(name string) *Channel {
	c := new(Channel)
	c.Buffer = make(chan int64, 1440)
	c.Name = name
	return c
}

func (c *Channel) GetColor() []float32 {
	return []float32{0.0, 0.9, 0.0, 0.0}
}
