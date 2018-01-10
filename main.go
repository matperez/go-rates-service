package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/matperez/cbr-http-service/cache"
	cbr "github.com/matperez/go-cbr-client"
)

var (
	httpListening string
)

func init() {
	flag.StringVar(&httpListening, "L", ":8080", "HTTP host and port to listen at")
	flag.Parse()
}

func main() {
	svc := NewService(cbr.NewClient(), cache.New())

	getRateHandler := httpTransport.NewServer(
		makeGetRateEndpoint(svc),
		decodeRateRequest,
		endcodeResponse,
	)

	http.Handle("/get-rate", getRateHandler)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go stopper(gracefulStop)

	log.Fatal(http.ListenAndServe(httpListening, nil))
}

// function to perform a graceful shutdown
func stopper(gracefulStop chan os.Signal) {
	sig := <-gracefulStop
	log.Printf("Caught sig: %+v", sig)
	log.Println("Wait for 2 seconds to finish processing")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
