package hoverboard

import (
	"encoding/binary"
	"log"
)

func (h *HoverboardImpl) receiveLoop() {
	for !h.stop {
		h.receive()
	}
}

func (h *HoverboardImpl) receive() []byte {
	buf := make([]byte, 1)
	n, err := h.port.Read(buf)
	if err != nil {
		log.Fatalf("error reading from serial port: %v\n", err)
		return nil
	}
	_ = n

	h.receiveBuf = append(h.receiveBuf, buf[0])

	if len(h.receiveBuf) > 1 {
		if binary.LittleEndian.Uint16(h.receiveBuf[len(h.receiveBuf)-2:]) == START_FRAME {
			if len(h.receiveBuf) == 20 { // 20 = length of message (18) + START_FRAME of next message (2)
				status, err := decodeReceiveMessage(h.receiveBuf[:18])
				if err != nil {
					log.Printf("error decoding message: %v", err)
				} else {
					h.status = *status
				}
			}

			// TODO: check, if this line can be moved up into the if to prevent dropping messages containing the START_FRAME in position > 0
			h.receiveBuf = h.receiveBuf[len(h.receiveBuf)-2:]
		}
	}

	return nil
}
