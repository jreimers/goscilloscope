package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
)

type RenderWindow struct {
	Title         string
	Width         int64
	Height        int64
	Channels      []*Channel
	quit          chan struct{}
	renderBuffers [][]int64
}

func NewRenderWindow(title string, width, height int, quit chan struct{}) (*RenderWindow, error) {
	rw := new(RenderWindow)
	rw.Title = title
	rw.Width = int64(width)
	rw.Height = int64(height)
	rw.quit = quit
	rw.renderBuffers = make([][]int64, 0)
	err := glfw.Init()
	if err != nil {
		return nil, err
	}
	return rw, nil
}
func (rw *RenderWindow) AddChannel(c *Channel) {
	rw.Channels = append(rw.Channels, c)
	rw.renderBuffers = append(rw.renderBuffers, make([]int64, 0))
}
func (rw *RenderWindow) Open() error {
	err := glfw.OpenWindow(int(rw.Width), int(rw.Height), 8, 8, 8, 8, 0, 0, glfw.Windowed)
	if err != nil {
		return err
	}
	glfw.SetWindowTitle(rw.Title)
	glfw.SetSwapInterval(2)
	glfw.SetKeyCallback(rw.onKey)
	glfw.SetWindowSizeCallback(rw.onResize)
	return nil
}

func (rw *RenderWindow) Render() {
	if glfw.WindowParam(glfw.Opened) == 0 {
		close(rw.quit)
	}
	for i, c := range rw.Channels {
		select {
		case val := <-c.Buffer:
			if int64(len(rw.renderBuffers[i])) > rw.Width {
				rw.renderBuffers[i] = rw.renderBuffers[i][1:]
			}
			rw.renderBuffers[i] = append(rw.renderBuffers[i], val)
		case <-rw.quit:
			rw.Close()
		default:
			continue
		}
	}
	rw.draw()
}
func (rw *RenderWindow) Close() {
	println("quit")
	glfw.CloseWindow()
	glfw.Terminate()
}

func (rw *RenderWindow) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.LINE_SMOOTH)
	for chanIdx, channel := range rw.Channels {
		color := channel.GetColor()
		gl.Color4f(color[0], color[1], color[2], color[3])
		gl.Begin(gl.LINES)
		lastY := int64(0)
		for i := 0; i < len(rw.renderBuffers[chanIdx]); i += 1 {
			y := rw.Height - (rw.renderBuffers[chanIdx][i] * rw.Height / 1024)
			gl.Vertex2i(i-1, int(lastY))
			gl.Vertex2i(i, int(y))
			lastY = y
		}
		gl.Flush()
		gl.End()
	}
	glfw.SwapBuffers()
}

func (rw *RenderWindow) onResize(w, h int) {
	// Write to both buffers, prevent flickering
	gl.DrawBuffer(gl.FRONT_AND_BACK)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Viewport(0, 0, w, h)
	gl.Ortho(0, float64(w), float64(h), 0, -1.0, 1.0)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	rw.Width = int64(w)
	rw.Height = int64(h)
}

func (rw *RenderWindow) onKey(key, state int) {
	switch key {
	case glfw.KeyEsc:
		close(rw.quit)
	case 67: // 'c'
		gl.Clear(gl.COLOR_BUFFER_BIT)
	}
}
