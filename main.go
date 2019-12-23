package main

import (
	"fmt"
	"github.com/maxp007/avito-test-task/config"
	"github.com/maxp007/avito-test-task/database"
	"github.com/maxp007/avito-test-task/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	log.SetReportCaller(true)

}

func main() {

	host := config.GetInstance().Data.Server.Host
	port := config.GetInstance().Data.Server.Port
	log.Print("Started main microservice on ", host, ":", port)
	err := http.ListenAndServe(host+":"+fmt.Sprint(port), router.GetRouter())

	if err != nil {
		log.Fatal("main service: ", err)
	}

	defer func() {
		err := database.ConnClose()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Print("Main microservice has stopped")
}
