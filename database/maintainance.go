package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func AnalyzeTable() {
	fmt.Println("analyze table adverts, cron job ")
	_, err := Pool.Exec(context.Background(), `ANALYZE adverts_schema.adverts;`)
	if err != nil {
		log.Println("analyze table adverts, cron job, err:", err)
		return
	}

	return
}

func VacuumTable() {
	fmt.Println("vacuum table adverts, cron job ")
	_, err := Pool.Exec(context.Background(), `VACUUM ANALYZE adverts_schema.adverts;`)
	if err != nil {
		log.Println("vacuum table adverts , cron job, err:", err)
		return
	}

	return
}
