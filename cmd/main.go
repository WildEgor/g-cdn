package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	server "github.com/WildEgor/g-cdn/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Infof("[Main] recieve signal %v", sig)
		done <- true
	}()

	srv, _ := server.NewServer()

	log.Info("[Main] Connect to Mongo")
	srv.Mongo.Connect()

	log.Printf("[Main] HTTP server is listening on port: %s", srv.AppConfig.Port)
	if err := srv.App.Listen(fmt.Sprintf(":%v", srv.AppConfig.Port)); err != nil {
		log.Panicf("[Main] Unable to start server. Reason: %v", err)
	}

	log.Info("[Main] Awaiting signal")
	<-done

	log.Info("[Main] Stopping")

	srv.Mongo.Disconnect()
	err := srv.App.Shutdown()
	if err != nil {
		log.Panicf("[Main] Unable to stop server. Reason: %v", err)
	}
}
