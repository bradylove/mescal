package main

import (
	"github.com/bradylove/mescal/msg"
	"net"
	"testing"
	"time"
)

const (
	succeed = "\033[0;32m\u2713\033[0m"
	failed  = "\033[0;31m\u2717\033[0m"
)

func startServer() {
	cfg = NewConfig()
	go runServer()

	time.Sleep(100 * time.Millisecond)
}

func TestClientServerHandshake(t *testing.T) {
	t.Log("Given I have a running server")
	{
		startServer()

		t.Log("\tAnd I connect to it as a new client")
		{
			conn, err := net.Dial("tcp", "localhost:8899")
			if err != nil {
				t.Fatalf("\t\tFailed to open connection: %v %s", failed, err.Error())
			}

			t.Logf("\t\tSuccessfully connected to server %v", succeed)

			cmd := msg.NewCommand("1234", msg.NewHandshakeCommand("mescal go version 1.2.3"))
			if err = cmd.Encode(conn); err != nil {
				t.Fatalf("\t\tFailed to send client handshake %v %s", failed, err.Error())
			}

			t.Logf("\t\tSuccessfully sent client handshake %v", succeed)

			res, err := msg.DecodeResult(conn)
			if err != nil {
				t.Fatalf("\t\tIt successfully received and decoded server handshake %v %s", failed, err.Error())
			}

			t.Logf("\t\tIt successfully received and decoded server handshake %v", succeed)

			if res.Id == "1234" {
				t.Logf("\t\tIt receives a response with proper id %s", succeed)
			} else {
				t.Fatalf("\t\tIt receives a response with proper id %s", failed)
			}

			subRes := res.SubResult.(msg.HandshakeResult)
			if len(subRes.ClientId) == 36 {
				t.Logf("\t\tIt received a client id %v", succeed)
			} else {
				t.Fatalf("\t\tIt received a client id %v", failed)
			}
		}
	}
}

// func TestClientSendsGetCommand(t *testing.T) {
//	t.Log("Given I have a running server")
//	{
//		startServer()

//		t.Log("\tAs a client I send a GET command")
//		{
//			conn, err := net.Dial("tcp", "localhost:8899")
//			if err != nil {
//				t.Fatalf("\t\tFailed to open connection: %v %s", failed, err.Error())
//			}

//			t.Logf("\t\tSuccessfully connected to server %v", succeed)

//			cmd := msg.NewCommand("1234", msg.NewHandshakeCommand("mescal go version 1.2.3"))
//			if err = cmd.Encode(conn); err != nil {
//				t.Fatalf("\t\tFailed to send client handshake %v %s", failed, err.Error())
//			}

//			res, err := msg.DecodeResult(conn)
//			if err != nil {
//				t.Fatalf("\t\tIt successfully received and decoded server handshake %v %s", failed, err.Error())
//			}
//		}
//	}
// }
