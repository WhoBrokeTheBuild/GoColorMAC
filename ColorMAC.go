package main

import (
    "net"
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    var event sdl.Event
    var running bool

    sdl.Init(sdl.INIT_EVERYTHING)

    window, err := sdl.CreateWindow("Color MAC", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 
                                    640, 640, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()

    surface, err := window.GetSurface()
    if err != nil {
        panic(err)
    }

    interfaces, err := net.Interfaces()
    if err != nil {
        panic(err)
    }

    colors := make([]sdl.Color, 0)

    for _, inter := range interfaces {
        if len(inter.HardwareAddr) == 0 {
            colors = append(colors, sdl.Color{0, 0, 0, 255}, sdl.Color{0, 0, 0, 255})
            continue
        }
        colors = append(colors, sdl.Color{
            inter.HardwareAddr[0],
            inter.HardwareAddr[1],
            inter.HardwareAddr[2],
            255,
        }, sdl.Color{
            inter.HardwareAddr[3],
            inter.HardwareAddr[4],
            inter.HardwareAddr[5],
            255,
        })
    }

    running = true
    for running {
        for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
            }
        }

        rect := sdl.Rect{0, 0, 128, 128}
        for _, color := range colors {
            surface.FillRect(&rect, color.Uint32())
            rect.X += rect.W
            if rect.X >= 640 {
                rect.X = 0
                rect.Y += rect.H
            }
        }
        window.UpdateSurface()

        sdl.Delay(1000 / 60)
    }

    sdl.Quit()
}
