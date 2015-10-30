package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net"
)

const (
	version = "0.1.0"
)

func main() {
	cfg := NewConfig()

	log.SetPrefix("[mescal] ")
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Starting Mescal on port :" + cfg.Port)

	RunServer(cfg)
}

func RunServer(cfg Config) {
	memoryStore := NewStore()

	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Println("Failed to start server:", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Client failed to open new connection:", err.Error())
		}

		id, err := newUUID()
		if err != nil {
			log.Println("Failed to create unique ID for new client")
			conn.Close()
			continue
		}

		log.Printf("New connection opened clientId=%s", id)

		c := NewClient(id, conn, memoryStore)
		go c.Handle()
	}
}

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
