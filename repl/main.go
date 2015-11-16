package main

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/bradylove/mescal/msg"
	"github.com/bradylove/mescal/repl/parser"
	"net"
	"os"
)

func newUUID() (string, error) {
	b := make([]byte, 16)

	n, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	if n != len(b) {
		return "", errors.New("Failed to read enough random bytes for UUID")
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func sendGetCommand(conn net.Conn, s parser.Statement) {
	msgId, _ := newUUID()
	cmd := msg.NewCommand(msgId, msg.NewGetCommand(s.Key))

	if err := cmd.Encode(conn); err != nil {
		fmt.Printf("Failed to send command to server: %v\n", err.Error())
	}

	result, err := msg.DecodeResult(conn)
	if err != nil {
		fmt.Printf("Failed to read result from server: %v\n", err.Error())
	}

	sr := result.SubResult.(msg.GetResult)

	fmt.Printf("%v: %v\n", sr.Key, sr.Value)
}

func sendSetCommand(conn net.Conn, s parser.Statement) {
	msgId, _ := newUUID()
	cmd := msg.NewCommand(msgId, msg.NewSetCommand(s.Key, s.Value, s.Expiry))

	if err := cmd.Encode(conn); err != nil {
		fmt.Printf("Failed to send command to server: %v\n", err.Error())
	}

	result, err := msg.DecodeResult(conn)
	if err != nil {
		fmt.Printf("Failed to read result from server: %v\n", err.Error())
	}

	if result.Status == msg.StatusSuccess {
		fmt.Println("OK")
	} else {
		fmt.Println("Failed")
	}
}

func main() {
	fmt.Println("Connecting to localhost:4567")

	conn, err := net.Dial("tcp", "localhost:4567")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	msgId, _ := newUUID()
	cmd := msg.NewCommand(msgId, msg.NewHandshakeCommand("mescal-repl version 0.0.0"))
	if err := cmd.Encode(conn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	r, err := msg.DecodeResult(conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if r.Status != msg.StatusSuccess {
		fmt.Fprintln(os.Stderr, "Failed to complete handshake with server")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')

		p := parser.NewParser(command)
		s, err := p.Parse()
		if err != nil {
			fmt.Printf("Unable to parse command: %s\n", err.Error())
			continue
		}

		switch s.Action {
		case parser.GET:
			sendGetCommand(conn, s)
		case parser.SET:
			sendSetCommand(conn, s)
		default:
			fmt.Println("Unknown command")
		}
	}
}
