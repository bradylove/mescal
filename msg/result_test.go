package msg_test

import (
	"bytes"
	"github.com/bradylove/mescal/msg"
	"testing"
)

func TestEncodeResult(t *testing.T) {
	t.Log("Given I have a get result")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		r := msg.NewResult("12345", msg.StatusSuccess, msg.NewGetResult("foo", "bar", int64(1442537301)))
		if err := r.Encode(buf); err != nil {
			t.Errorf("It should be able to encode a GetResult %v", failed)
		}
		t.Logf("It should be able to encode a GetResult %v", succeed)

		expectedBytes := []byte{0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 0, 1, 0, 0, 0, 25, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 85, 95, 251, 85, 0, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 98, 97, 114, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0}
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

func TestDecodeResult(t *testing.T) {
	t.Log("Given I have an encoded get result")
	{
		encodedMsg := []byte{0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 0, 0, 0, 1, 0, 0, 0, 25, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 85, 95, 251, 85, 0, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 5, 0, 0, 0, 34, 0, 0, 0, 102, 111, 111, 0, 0, 0, 0, 0, 98, 97, 114, 0, 0, 0, 0, 0, 49, 50, 51, 52, 53, 0, 0, 0}
		buf := bytes.NewBuffer(encodedMsg)

		res, err := msg.DecodeResult(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode a GetResult: %v %v", err, failed)
		}

		t.Logf("\tIt should be able to decode a GetResult %v", succeed)

		if res.Action == msg.ActionGet {
			t.Logf("\tIt should be a decoded GetResult %v", succeed)
		} else {
			t.Logf("\tIt should be a decoded GetResult %v", failed)
		}
	}
}

func TestEncodeAndDecodeHandshakeResult(t *testing.T) {
	t.Log("Given I have a handshake result")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		res := msg.NewResult("12345", msg.StatusSuccess, msg.NewHandshakeResult("123"))
		if err := res.Encode(buf); err != nil {
			t.Fatalf("\tIt should be able to encode the handshake result: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to encode the handshake result %v", succeed)

		decodedRes, err := msg.DecodeResult(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode the handshake result: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to decode the handshake result %v", succeed)

		if decodedRes.Action == msg.ActionHandshake {
			t.Logf("\tIt should be of type ActionHandshake %v", succeed)
		} else {
			t.Fatalf("\tIt should be of type ActionHandshake %v", failed)
		}
	}
}

func TestEncodeAndDecodeSetResult(t *testing.T) {
	t.Log("Given I have a set result")
	{
		var b []byte
		buf := bytes.NewBuffer(b)

		res := msg.NewResult("12345", msg.StatusSuccess, msg.NewSetResult())
		if err := res.Encode(buf); err != nil {
			t.Fatalf("\tIt should be able to encode the set result: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to encode the set result %v", succeed)

		decodedRes, err := msg.DecodeResult(buf)
		if err != nil {
			t.Fatalf("\tIt should be able to decode the set result: %v %v", err, failed)
		}
		t.Logf("\tIt should be able to decode the set result %v", succeed)

		if decodedRes.Action == msg.ActionSet {
			t.Logf("\tIt should be of type ActionSet %v", succeed)
		} else {
			t.Fatalf("\tIt should be of type ActionSet %v", failed)
		}
	}
}
