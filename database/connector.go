package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/maxp007/avito-test-task/config"

	_ "github.com/maxp007/avito-test-task/config"

	log "github.com/sirupsen/logrus"
)

var Pool *pgxpool.Pool

func init() {

	Host := config.GetInstance().Data.Database.DBserver.Host
	Port := config.GetInstance().Data.Database.DBserver.Port
	Database := config.GetInstance().Data.Database.Credentials.DbName
	User := config.GetInstance().Data.Database.Credentials.User
	Password := config.GetInstance().Data.Database.Credentials.Pass

	conn_string := fmt.Sprintf("host=%s port=%d user=%s database=%s password=%s ", Host, Port, User, Database, Password)
	conf, err := pgxpool.ParseConfig(conn_string)

	if err != nil {
		log.Fatal("Config parse error: %s", err)
		return
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %s", err))
		return
	}
	log.Printf("Successfully connected to database on %s:%d", Host, Port)

	Pool = pool
}
func ConnClose() (err error) {
	defer func(err *error) {
		Pool.Close()
	}(&err)
	return err
}
