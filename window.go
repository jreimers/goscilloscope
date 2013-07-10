package main

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
)

type RenderWindow struct {
	Title     string
	Width     int
	Height    int
	Rendering bool
}

func InitRenderWindow(title string, width, height int) (*RenderWindow, error) {
	rw := new(RenderWindow)
	rw.Title = title
	rw.Width = width
	rw.Height = height
	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	err = glfw.OpenWindow(width, height, 8, 8, 8, 8, 0, 0, glfw.Windowed)
	if err != nil {
		return nil, err
	}

	glfw.SetWindowTitle(title)
	glfw.SetSwapInterval(2)
	glfw.SetKeyCallback(rw.onKey)
	glfw.SetWindowSizeCallback(rw.onResize)
	return rw, nil
}

func (rw *RenderWindow) BeginRendering() {
	rw.Rendering = true
	for rw.Rendering && glfw.WindowParam(glfw.Opened) == 1 {
		rw.Render()
	}
	glfw.CloseWindow()
	glfw.Terminate()
}

var buf []int

func newData(buffer []int) {
	buf = buffer
}

func (rw *RenderWindow) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.LINE_SMOOTH)
	gl.Color4f(0.0, 0.9, 0.0, 0.0)
	gl.Begin(gl.LINES)
	lastY := 0
	for i := 0; i < len(buf); i += 1 {
		y := rw.Height - (buf[i] * rw.Height / 1024)
		gl.Vertex2i(i-1, lastY)
		gl.Vertex2i(i, y)
		lastY = y
	}
	gl.Flush()
	gl.End()

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
	rw.Width = w
	rw.Height = h
}

func (rw *RenderWindow) onKey(key, state int) {
	switch key {
	case glfw.KeyEsc:
		rw.Rendering = state == 0
	case 67: // 'c'
		gl.Clear(gl.COLOR_BUFFER_BIT)
	}
}
