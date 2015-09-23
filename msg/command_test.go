package msg_test

import (
	"bytes"
	"github.com/bradylove/mescal/msg"
	"testing"
	"time"
)

func TestCommandEncode(t *testing.T) {
	t.Log("Given I have a get command")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		cmd := msg.NewCommand("12345", msg.NewGetCommand("foo"))
		if err := cmd.Encode(buf); err != nil {
			t.Errorf("It should be able to encode a GetCommand %v", failed)
		}

		t.Logf("It should be able to encode a GetCommand %v", succeed)

		expectedBytes := []byte{0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 13, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0}
		actualBytes := buf.Bytes()

		slicesMatch := true
		for i := range expectedBytes {
			if expectedBytes[i] != actualBytes[i] {
				slicesMatch = false
			}
		}

		if slicesMatch {
			t.Logf("It should have the correct bytes %v", succeed)
		} else {
			t.Errorf("It should have the correct bytes %v", failed)
		}
	}
}

func TestDecodeCommand(t *testing.T) {
	t.Log("Given I have an encoded get command")
	{
		encodedMsg := []byte{0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 13, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0}
		buf := bytes.NewBuffer(encodedMsg)

		cmd, err := msg.DecodeCommand(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode a GetCommand %v", failed)
		}

		t.Logf("\tIt should be able to decode a GetCommand %v", succeed)

		if cmd.Action == msg.ActionGet {
			t.Logf("\tIt should be a decoded GetCommand %v", succeed)
		} else {
			t.Logf("\tIt should be a decoded GetCommand %v", failed)
		}
	}
}

func TestEncodeAndDecodeHandshakeCommand(t *testing.T) {
	t.Log("Given I have a handshake command")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		cmd := msg.NewCommand("12345", msg.NewHandshakeCommand("Go v0.1.1"))
		if err := cmd.Encode(buf); err != nil {
			t.Fatalf("\tIt should be able to encode the handshake command: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to encode the handshake command %v", succeed)

		decodedCmd, err := msg.DecodeCommand(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode the handshake command: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to decode the handshake command %v", succeed)

		if decodedCmd.Action == msg.ActionHandshake {
			t.Logf("\tIt should be of type ActionHandshake %v", succeed)
		} else {
			t.Fatalf("\tIt should be of type ActionHandshake %v", failed)
		}
	}
}

func TestEncodeAndDecodeSetCommand(t *testing.T) {
	t.Log("Given I have a set command")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		cmd := msg.NewCommand("12345", msg.NewSetCommand("foo", "bar", time.Now().Unix()))
		if err := cmd.Encode(buf); err != nil {
			t.Fatalf("\tIt should be able to encode the set command: %v %v", err, failed)
		}

		t.Logf("\tIt should be able to encode the set command %v", succeed)

		decodedCmd, err := msg.DecodeCommand(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode the set command: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to decode the set command %v", succeed)

		if decodedCmd.Action == msg.ActionSet {
			t.Logf("\tIt should be of type ActionSet %v", succeed)
		} else {
			t.Fatalf("\tIt should be of type ActionSet %v", failed)
		}
	}
}
