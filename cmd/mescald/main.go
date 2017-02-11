package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bradylove/mescal/pkg/server"
)

func main() {
	mescalAddr := flag.String("addr", ":7321", "Mescal server address")
	healthAddr := flag.String("health", "localhost:4040", "Address for healthcheck and stats")
	flag.Parse()

	actualAddr, err := server.Start(*mescalAddr)
	if err != nil {
		log.Println("Failed to start Mescal:", err)
	}
	log.Printf("Starting Mescal at %s", actualAddr)

	log.Println(http.ListenAndServe(*healthAddr, nil))
}
