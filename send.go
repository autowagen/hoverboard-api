package hoverboard

import (
	"log"
	"time"
)

const sendLoopInterval = 50

func (h *Hoverboard) sendLoop() {
	lastLoop := time.Now()
	for !h.stop {
		duration := time.Since(lastLoop)
		if duration.Milliseconds() > sendLoopInterval+2 {
			log.Printf("sendLoop not reliable (%v instead of %v)", duration.Milliseconds(), sendLoopInterval)
		}
		lastLoop = time.Now()
		h.send()
		time.Sleep(sendLoopInterval * time.Millisecond)
	}
}

func (h *Hoverboard) send() {
	//log.Printf("sending steer=%v speed=%v", h.steer, h.speed)
	message, err := encodeSendMessage(h.steer, h.speed)
	if err != nil {
		log.Printf("failed to encode message: %v", err)
	}
	_, err = h.port.Write(message)
	if err != nil {
		log.Printf("failed to send message: %v", err)
	}
}
