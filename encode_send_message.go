package hoverboard

import (
	"bytes"
	"encoding/binary"
)

func encodeSendMessage(steer int16, speed int16) ([]byte, error) {
	checksum := START_FRAME ^ uint16(steer) ^ uint16(speed)
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, START_FRAME); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, steer); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, speed); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, checksum); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
