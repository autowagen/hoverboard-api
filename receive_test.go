package hoverboard

import (
	"bytes"
	"encoding/binary"
	"hoverboard/mock"
	"testing"
)

func TestReceiveCaseDropIncompleteMessage(t *testing.T) {
	mockRwc := mock.NewMockReadWriteCloser()
	h := Hoverboard{mockRwc, 0, 0, false, []byte{0x2A, 0x01, 0x00, 0x00, 0xAF, 0x5A, 0xCD}, HoverboardStatus{0, 0, 0, 0, 0, 0, 0}}
	mockRwc.B = bytes.NewBuffer([]byte{0xAB})

	h.receive()

	if len(h.receiveBuf) != 2 || binary.LittleEndian.Uint16(h.receiveBuf) != START_FRAME {
		t.Errorf("expected receiveBuf to contain only the start-frame, but was: %X", h.receiveBuf)
	}
}

func TestReceiveCaseHandleMessageAndDropItFromTheBuffer(t *testing.T) {
	mockRwc := mock.NewMockReadWriteCloser()
	h := Hoverboard{mockRwc, 0, 0, false, []byte{0xCD, 0xAB, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x16, 0x00, 0xE9, 0x0F, 0x1E, 0x01, 0x00, 0x00, 0xAA, 0xA5, 0xCD}, HoverboardStatus{0, 0, 0, 0, 0, 0, 0}}
	mockRwc.B = bytes.NewBuffer([]byte{0xAB})

	h.receive()

	if len(h.receiveBuf) != 2 || binary.LittleEndian.Uint16(h.receiveBuf) != START_FRAME {
		t.Errorf("expected receiveBuf to contain only the start-frame, but was: %X", h.receiveBuf)
	}

	expectedStatus := HoverboardStatus{-191, 160, -153, 22, 4073, 286, 0}
	if h.status != expectedStatus {
		t.Errorf("expected status to be %v, but was: %v", expectedStatus, h.status)
	}
}

func TestReceiveCaseMidMessage(t *testing.T) {
	mockRwc := mock.NewMockReadWriteCloser()
	h := Hoverboard{mockRwc, 0, 0, false, []byte{0x2A, 0x01, 0x00, 0x00, 0xAF, 0x5A, 0xCD}, HoverboardStatus{0, 0, 0, 0, 0, 0, 0}}
	mockRwc.B = bytes.NewBuffer([]byte{0x12})

	h.receive()

	receiveBufLength := len(h.receiveBuf)
	if receiveBufLength != 8 {
		t.Errorf("expected size of receiveBuf to be 8, but was %v", receiveBufLength)
	}

	lastByte := h.receiveBuf[7]
	if lastByte != 0x12 {
		t.Errorf("expected last byte to be 0x12, but was %X", lastByte)
	}
}
