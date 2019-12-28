package database

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func Db_get_estimated_adverts_number() (advers_estimated int64, err error) {

	row := Pool.QueryRow(context.Background(), `SELECT reltuples FROM pg_class WHERE oid = 'adverts_schema.adverts'::regclass;`)
	err = row.Scan(&advers_estimated)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
