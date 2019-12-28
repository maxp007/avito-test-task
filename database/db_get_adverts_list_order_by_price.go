package database

import (
	"context"
	"github.com/maxp007/avito-test-task/config"
	"github.com/maxp007/avito-test-task/models"
	"log"
)

func Db_get_adverts_list_order_by_price(page int64, order_price string) (response []models.AdvertListElement, err error) {
	ads_per_page := config.GetInstance().Data.AdvertList.MaxAdvertsOnPage
	//getadvertslistorderedbyprice(page integer, ads_per_page integer, order_by_price text) returns SETOF adverts_schema.shortadvertdata

	rows, err := Pool.Query(context.Background(), "select * FROM adverts_schema.getadvertslistorderedbyprice($1,$2,$3);", int32(page), int32(ads_per_page), order_price)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	response = make([]models.AdvertListElement, 0, 0)

	var i int = 0
	for rows.Next() {
		response = append(response, models.AdvertListElement{})
		err = rows.Scan(
			&response[i].Id,
			&response[i].Title,
			&response[i].MainPicture,
			&response[i].Price,
			&response[i].Date)
		if err != nil {
			log.Println(err)
		}
		i += 1
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		log.Println(err)
	}

	return
}
