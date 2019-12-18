package main

import (
	"fmt"
	config "github.com/maxp007/avito-test-task/config"
	"github.com/maxp007/avito-test-task/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	log.SetReportCaller(true)

}

func main() {

	err := config.GetInstance().LoadConfig()

	if err != nil {
		log.Fatal("config reader,", err)
	}

	log.Print("Started main microservice on ", config.GetInstance().Data.Server.Host, ":", config.GetInstance().Data.Server.Port)
	err = http.ListenAndServe(config.GetInstance().Data.Server.Host+":"+fmt.Sprint(config.GetInstance().Data.Server.Port), router.GetRouter())
	if err != nil {
		log.Fatal("main service: ", err)
	}

	log.Print("Main microservice has stopped")
}
