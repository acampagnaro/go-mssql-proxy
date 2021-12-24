package main

import (
	"context"
	proxy2 "go-mssql-proxy/proxy"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	proxy := proxy2.NewProxy("alto-da-xv.clinic.inf.br", ":1433", ctx)
	proxy.EnableDecoding()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			log.Printf("Signal received %v, stopping and exiting...", sig)
			cancel()
		}
	}()

	err := proxy.Start("1433")
	if err != nil {
		log.Fatal(err)
	}
}
