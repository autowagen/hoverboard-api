package hoverboard

import (
	"errors"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

type Hoverboard struct {
	port       io.ReadWriteCloser
	steer      int16
	speed      int16
	stop       bool
	receiveBuf []byte
	status     HoverboardStatus
}

type HoverboardStatus struct {
	Cmd1       int16
	Cmd2       int16
	SpeedRMaes int16
	SpeedLMaes int16
	BatVoltage int16
	BoardTemp  int16
	CmdLed     uint16
}

func NewHoverboard(portName string) (*Hoverboard, error) {
	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	port, err := serial.Open(options)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open serial port: %v", err))
	}
	h := &Hoverboard{port, 0, 0, false, make([]byte, 30), HoverboardStatus{0, 0, 0, 0, 0, 0, 0}}
	go h.receiveLoop()
	go h.sendLoop()
	return h, nil
}

func (h *Hoverboard) Close() {
	h.stop = true
	h.port.Close()
}

func (h *Hoverboard) SetSteer(steer int16) {
	h.steer = steer
}

func (h *Hoverboard) SetSpeed(speed int16) {
	h.speed = speed
}

func (h *Hoverboard) GetStatus() HoverboardStatus {
	return h.status
}
