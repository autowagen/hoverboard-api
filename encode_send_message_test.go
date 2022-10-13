package hoverboard

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncodeSendMessage(t *testing.T) {
	message, err := encodeSendMessage(100, 200)
	if err != nil {
		t.Errorf("expected err to be nil: %v", err)
	}

	fmt.Printf("%04X\n", message)

	messageLength := len(message)
	expectedMessageLength := 8
	if messageLength != expectedMessageLength {
		t.Errorf("expected message to have length %v, but was %v", messageLength, expectedMessageLength)
	}

	expected := []byte{0xCD, 0xAB, 0x64, 0x00, 0xC8, 0x00, 0x61, 0xAB}
	if bytes.Compare(message, expected) != 0 {
		t.Errorf("expected message to be %X, but was %X", expected, message)
	}
}
