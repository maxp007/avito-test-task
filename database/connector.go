package database

import (
	"context"
	"fmt"
	"github.com/maxp007/avito-test-task/config"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4"
)

func init() {
	host := config.GetInstance().Data.Database.DBserver.Host
	port := config.GetInstance().Data.Database.DBserver.Port

	conn, err := pgx.Connect(context.Background(), host+":"+fmt.Sprint(port))
	if err != nil {
		log.Fatal("Unable to connection to database: %v\n", err)

	}
	defer conn.Close(context.Background())

}
