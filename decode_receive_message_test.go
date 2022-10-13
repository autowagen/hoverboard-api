package hoverboard

import (
	"testing"
)

func TestDecodeReceiveMessageCaseValid(t *testing.T) {
	cases := []struct {
		buffer          []byte
		expectedMessage HoverboardStatus
	}{
		{
			buffer:          []byte{0xCD, 0xAB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x00, 0xC4, 0x0F, 0x03, 0x01, 0x00, 0x00, 0x19, 0xA5},
			expectedMessage: HoverboardStatus{0, 0, 0, 19, 4036, 259, 0},
		},
		{
			buffer:          []byte{0xCD, 0xAB, 0x9B, 0x00, 0xC3, 0x00, 0xC6, 0xFF, 0xA6, 0x00, 0xD4, 0x0F, 0x0B, 0x01, 0x00, 0x00, 0x2A, 0x5A},
			expectedMessage: HoverboardStatus{155, 195, -58, 166, 4052, 267, 0},
		},
		{
			buffer:          []byte{0xCD, 0xAB, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x16, 0x00, 0xE9, 0x0F, 0x1E, 0x01, 0x00, 0x00, 0xAA, 0xA5},
			expectedMessage: HoverboardStatus{-191, 160, -153, 22, 4073, 286, 0},
		},
		{
			buffer:          []byte{0xCD, 0xAB, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x15, 0x00, 0xE9, 0x0F, 0x20, 0x01, 0x00, 0x00, 0x97, 0xA5},
			expectedMessage: HoverboardStatus{-191, 160, -153, 21, 4073, 288, 0},
		},
		{
			buffer:          []byte{0xCD, 0xAB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xEF, 0x0F, 0x20, 0x01, 0x00, 0x00, 0x02, 0xA5},
			expectedMessage: HoverboardStatus{0, 0, 0, 0, 4079, 288, 0},
		},
	}

	for i, test := range cases {
		message, err := decodeReceiveMessage(test.buffer)

		if err != nil {
			t.Errorf("expected error to be nil, but was: %v", err)
		}

		if *message != test.expectedMessage {
			t.Errorf("expected message to be %v, but was: %v (case %v)", test.expectedMessage, *message, i)
		}
	}
}

func TestDecodeReceiveMessageCaseWrongLength(t *testing.T) {
	buffer := []byte{0xCD, 0xAB, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x16, 0x00, 0xE9, 0x0F, 0x1E, 0x01, 0x00, 0x00, 0xAA}
	_, err := decodeReceiveMessage(buffer)

	if err == nil {
		t.Errorf("expected error to be 'invalid length', but was: nil")
	}

	if err.Error() != "invalid length" {
		t.Errorf("expected error to be 'invalid length', but was: %v", err.Error())
	}
}

func TestDecodeReceiveMessageCaseInvalidStartFrame(t *testing.T) {
	buffer := []byte{0xCD, 0xAA, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x16, 0x00, 0xE9, 0x0F, 0x1E, 0x01, 0x00, 0x00, 0xAA, 0x00}
	_, err := decodeReceiveMessage(buffer)

	if err == nil {
		t.Errorf("expected error to be 'start_frame doesn't match', but was: nil")
	}

	if err.Error() != "start_frame doesn't match" {
		t.Errorf("expected error to be 'start_frame doesn't match', but was: %v", err.Error())
	}
}

func TestDecodeReceiveMessageCaseInvalidChecksum(t *testing.T) {
	buffer := []byte{0xCD, 0xAB, 0x41, 0xFF, 0xA0, 0x00, 0x67, 0xFF, 0x16, 0x00, 0xE9, 0x0F, 0x1E, 0x01, 0x00, 0x00, 0xAA, 0x00}
	_, err := decodeReceiveMessage(buffer)

	if err == nil {
		t.Errorf("expected error to be 'invalid message. checksum didn't match', but was: nil")
	}

	if err.Error() != "invalid message. checksum didn't match" {
		t.Errorf("expected error to be 'invalid message. checksum didn't match', but was: %v", err.Error())
	}
}
