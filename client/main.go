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

	cmd = msg.NewCommand("1233", msg.NewSetCommand("ascii", "characters", time.Now().Unix()))
	if err = cmd.Encode(conn); err != nil {
		fmt.Println(err)
		return
	}

	res, err = msg.DecodeResult(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Action)

	for {
		fmt.Println("\n\n")
		time.Sleep(1 * time.Second)
		cmd = msg.NewCommand("1234", msg.NewGetCommand("ascii"))
		if err = cmd.Encode(conn); err != nil {
			fmt.Println(err)
			return
		}

		res, err = msg.DecodeResult(conn)
		if err != nil {
			panic(err)
		}

		fmt.Println(res.Action)

		switch sb := res.SubResult.(type) {
		case msg.GetResult:
			fmt.Println("SubResult is GetResult")
			fmt.Println("Key:  ", sb.Key)
			fmt.Println("Value:", sb.Value)
		default:
			fmt.Println("Unknown SubCommand")
		}
	}
}
