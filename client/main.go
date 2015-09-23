package main

import (
	"fmt"
	"github.com/bradylove/mescal/msg"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		fmt.Println("Failed to open connection:", err)
	}

	cmd := msg.NewCommand("1234", msg.NewHandshakeCommand("mescal go version 1.2.3"))
	if err = cmd.Encode(conn); err != nil {
		fmt.Println(err)
		return
	}

	res, err := msg.DecodeResult(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Id)
	fmt.Printf("%+v", res.SubResult.(msg.HandshakeResult))

	for {
		fmt.Println("\n\n")
		time.Sleep(1 * time.Second)
		cmd := msg.NewCommand("1234", msg.NewGetCommand("foo"))
		if err = cmd.Encode(conn); err != nil {
			fmt.Println(err)
			return
		}

		decodedCmd, err := msg.DecodeCommand(conn)
		if err != nil {
			panic(err)
		}

		fmt.Println(decodedCmd.Action)

		switch sb := decodedCmd.SubCommand.(type) {
		case msg.GetCommand:
			fmt.Println("SubCommand is GetCommand")
			fmt.Println("Key:", sb.Key)
		default:
			fmt.Println("Unknown SubCommand")
		}
	}
}
