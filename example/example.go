package main

import (
	"github.com/autowagen/hoverboard-api"
	"log"
	"time"
)

func main() {
	h, err := hoverboard.NewHoverboard("/dev/ttyUSB0")
	if err != nil {
		panic(err)
	}
	defer h.Close()

	max := 150

	go (func() {
		for true {
			status := h.GetStatus()
			log.Printf("%v\n", status)
			time.Sleep(500 * time.Millisecond)
		}
	})()

	i := 0
	for true {
		for i < max {
			i += 20
			h.SetSpeed(int16(i))
			time.Sleep(500 * time.Millisecond)
		}
		for i > -max {
			i -= 20
			h.SetSpeed(int16(i))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
