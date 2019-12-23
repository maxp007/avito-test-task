package database

import (
	"context"

	"github.com/maxp007/avito-test-task/models"
	"log"
)

func Db_create_advert(body models.AdvertCreateBody) (result int64, err error) {
	tx, err := Pool.Begin(context.Background())
	if err != nil {
		return 0, err
	}

	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "SELECT adverts_schema.createadvert($1, $2, $3, $4);",
		body.Title, body.Description, body.Pictures, body.Price).Scan(&result)
	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return
	}
	return
}
