package hoverboard

import (
	"encoding/binary"
	"errors"
)

func decodeReceiveMessage(buffer []byte) (*HoverboardStatus, error) {
	if len(buffer) != 18 {
		return nil, errors.New("invalid length")
	}

	startFrame := binary.LittleEndian.Uint16(buffer[0:2])
	if startFrame != START_FRAME {
		return nil, errors.New("start_frame doesn't match")
	}

	cmd1 := int16(binary.LittleEndian.Uint16(buffer[2:4]))
	cmd2 := int16(binary.LittleEndian.Uint16(buffer[4:6]))
	speedRMaes := int16(binary.LittleEndian.Uint16(buffer[6:8]))
	speedLMaes := int16(binary.LittleEndian.Uint16(buffer[8:10]))
	batVoltage := int16(binary.LittleEndian.Uint16(buffer[10:12]))
	boardTemp := int16(binary.LittleEndian.Uint16(buffer[12:14]))
	cmdLed := binary.LittleEndian.Uint16(buffer[14:16])
	checksum := binary.LittleEndian.Uint16(buffer[16:18])

	calculatedChecksum := startFrame ^ uint16(cmd1) ^ uint16(cmd2) ^ uint16(speedRMaes) ^ uint16(speedLMaes) ^ uint16(batVoltage) ^ uint16(boardTemp) ^ cmdLed

	if checksum != calculatedChecksum {
		return nil, errors.New("invalid message. checksum didn't match")
	}

	message := HoverboardStatus{cmd1, cmd2, speedRMaes, speedLMaes, batVoltage, boardTemp, cmdLed}
	return &message, nil
}
