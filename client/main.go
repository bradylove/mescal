package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/bradylove/mescal/msg"
	"io/ioutil"
	// "net"
	"time"
)

func main() {
	certPool := x509.NewCertPool()
	caFile, err := ioutil.ReadFile("../test/ca.pem")

	block, _ := pem.Decode(caFile)
	caCert, err := x509.ParseCertificate(block.Bytes)

	certPool.AddCert(caCert)
	tlsCfg := tls.Config{
		RootCAs:    certPool,
		ServerName: "localhost",
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:4333", &tlsCfg)

	// conn, err := net.Dial("tcp", "localhost:4333")
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
