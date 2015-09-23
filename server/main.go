package main

import (
	"github.com/bradylove/mescal/msg"
	"io"
	"log"
	"net"
)

var cfg Config

const (
	version = "0.1.0"
)

func handleConnection(conn net.Conn) {
HandlerLoop:
	for {
		decodedCmd, err := msg.DecodeCommand(conn)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
				conn.Close()
				break HandlerLoop
			}
			panic(err)
		}

		switch sb := decodedCmd.SubCommand.(type) {
		case msg.GetCommand:
			log.Printf("Command received action=%d subCommand=GetCommand key=%s\n",
				decodedCmd.Action,
				sb.Key)

			decodedCmd.Encode(conn)
		default:
			log.Println("Unknown sub command received")
		}
	}
}

func main() {
	cfg = NewConfig()

	log.SetPrefix("[mescal] ")
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting Mescal on port :" + cfg.Port)

	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Println("Failed to start server:", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Client failed to open new connection:", err.Error())
		}

		log.Println("New connection opened")

		go handleConnection(conn)
	}
}
