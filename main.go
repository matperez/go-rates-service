package main

import (
	"flag"
	"log"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/matperez/cbr-http-service/cache"
	cbr "github.com/matperez/go-cbr-client"
)

var (
	httpListening string
)

func init() {
	flag.StringVar(&httpListening, "L", ":8080", "HTTP port")
	flag.Parse()
}

func main() {
	log.Println("Started...")

	svc := NewService(cbr.NewClient(), cache.New())

	getRateHandler := httpTransport.NewServer(
		makeGetRateEndpoint(svc),
		decodeRateRequest,
		endcodeResponse,
	)

	http.Handle("/get-rate", getRateHandler)
	log.Fatal(http.ListenAndServe(httpListening, nil))

	log.Println("Stoped...")
}
