package main

import (
	"errors"
	"github.com/bradylove/mescal/msg"
	"io"
	"log"
	"net"
)

type Client struct {
	id   string
	conn net.Conn
}

func NewClient(id string, conn net.Conn) Client {
	return Client{id, conn}
}

func (c Client) close() {
	log.Printf("Client disconnected clientId=%s", c.id)
	c.conn.Close()
}

func (c Client) handle() {
	err := c.handshake()
	if err != nil {
		log.Printf("Client handshake failed clientId=%s reason=%s", c.id, err.Error())
		c.close()
		return
	}

HandlerLoop:
	for {
		decodedCmd, err := msg.DecodeCommand(c.conn)
		if err != nil {
			if err == io.EOF {
				c.close()
				break HandlerLoop
			}

			// TODO: DO NOT PANIC!
			panic(err)
		}

		switch sb := decodedCmd.SubCommand.(type) {
		case msg.GetCommand:
			log.Printf("Command received action=%d subCommand=GetCommand key=%s\n",
				decodedCmd.Action,
				sb.Key)

			store.HandleCommand(decodedCmd, c.conn)
		case msg.SetCommand:
			log.Printf("Command received action=%d subCommand=SetCommand key=%s\n",
				decodedCmd.Action,
				sb.Key)

			store.HandleCommand(decodedCmd, c.conn)
		default:
			log.Println("Unknown sub command received")
		}
	}
}

func (c Client) handshake() error {
	decodedCmd, err := msg.DecodeCommand(c.conn)
	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed before handshake could be completed")
			c.conn.Close()
		}

		return err
	}

	switch decodedCmd.SubCommand.(type) {
	case msg.HandshakeCommand:
		res := msg.NewResult(decodedCmd.Id, msg.StatusSuccess, msg.NewHandshakeResult(c.id))

		if err := res.Encode(c.conn); err != nil {
			return errors.New("Failed to send handshake response to client")
		}

		log.Printf("Handshake successful clientId=%s", c.id)

		return nil
	default:
		return errors.New("First message from client must be handshake")
	}
}
