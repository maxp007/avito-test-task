package main

import (
	"context"
	"fmt"
	"github.com/maxp007/avito-test-task/cache"
	"github.com/maxp007/avito-test-task/config"
	"github.com/maxp007/avito-test-task/database"
	"github.com/maxp007/avito-test-task/router"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	log.SetReportCaller(true)
}

var (
	host = config.GetInstance().Data.Server.Host
	port = config.GetInstance().Data.Server.Port
)

func main() {
	log.Print("Starting main microservice on ", host, ":", port)

	server := &http.Server{Addr: fmt.Sprintf("%s:%d", host, port), Handler: router.GetRouter()}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("main service: ", err)

		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	defer func() {
		err := database.ConnClose()
		if err != nil {
			log.Println(err)
		}

		err = cache.ConnClose()
		if err != nil {
			log.Println(err)
		}
	}()
}
