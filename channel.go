package main

type Channel struct {
	Buffer chan int64
	Name   string
	color  []float32
}

func NewChannel(name string) *Channel {
	c := new(Channel)
	c.Buffer = make(chan int64, 1440)
	c.Name = name
	return c
}

func (c *Channel) SetColor(r, g, b float32) {
	c.color = []float32{r, g, b, 0.0}
}

func (c *Channel) GetColor() []float32 {
	return c.color
}
