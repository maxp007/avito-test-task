package database

import (
	"context"
	"github.com/maxp007/avito-test-task/models"
)

func Db_get_advert_by_id(id int64, fields []string) (response models.AdvertResponse, err error) {
	//getadvert(id_arg bigint, fields_str text DEFAULT ''::text) returns adverts_schema.adverts

	row := Pool.QueryRow(context.Background(), "select * FROM adverts_schema.getadvert($1,$2);", id, fields)
	response = models.AdvertResponse{}

	err = row.Scan(
		&response.Id,
		&response.Title,
		&response.Description,
		&response.Pictures,
		&response.MainPicture,
		&response.Price,
	)
	if err != nil {
		return
	}
	return
}
