package avito_test_task

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

const (
	PORT    = 8080
	IP_ADDR = "127.0.0.1"
)

func init() {
	log.SetReportCaller(true)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := http.ListenAndServe(IP_ADDR+":"+fmt.Sprint(PORT), nil)
		if err != nil {
			log.Print("main service, fatal error: ", err)
		}
	}()

	wg.Wait()
	log.Print("Main microservice has stopped")
}
