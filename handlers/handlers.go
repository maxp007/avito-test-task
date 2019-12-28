package handlers

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/julienschmidt/httprouter"
	"github.com/maxp007/avito-test-task/cache"
	"github.com/maxp007/avito-test-task/config"
	"github.com/maxp007/avito-test-task/database"
	"github.com/maxp007/avito-test-task/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	Expiration_Time  = config.GetInstance().Data.Cache.Expiration
	adverts_per_page = config.GetInstance().Data.AdvertList.MaxAdvertsOnPage
)

func GetAdvertListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Add("Content-type", "application/json;charset=utf-8")

	var page int64
	var order string
	var sort string

	str_page := r.URL.Query().Get("page")
	page, err := strconv.ParseInt(str_page, 10, 64)
	if err != nil || page == 0 || page < 0 {
		page = 1
	}

	order = r.URL.Query().Get("order")
	sort = r.URL.Query().Get("sort")

	if order != "date" && order != "price" {
		order = "date"
	}

	if sort != "desc" && sort != "asc" {
		sort = "desc"
	}

	var result []models.AdvertListElement

	statux_result := cache.Cache.Get(fmt.Sprintf("p:%d,o:%s,s:%s", page, order, sort))
	if statux_result.Err() != redis.Nil {
		res, err := statux_result.Result()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		_, err = w.Write([]byte(res))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if order == "date" {
		result, err = database.Db_get_adverts_list_order_by_date(page, sort)
	} else {
		result, err = database.Db_get_adverts_list_order_by_price(page, sort)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ads_number, err := database.Db_get_estimated_adverts_number()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response_data := models.AdvertsList{
		PagesTotal:  ads_number / int64(adverts_per_page),
		Page:        page,
		AdvsPerPage: int64(adverts_per_page),
		Adverts:     result,
	}

	resp_bytes, err := (response_data).MarshalJSON()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := cache.Cache.Set(fmt.Sprintf("p:%d,o:%s,s:%s", page, order, sort), resp_bytes, time.Minute*time.Duration(Expiration_Time))
	if status.Err() != nil {
		log.Print("cache setting error ", status.Err())
	}
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp_bytes)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetAdvertHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Add("Content-type", "application/json;charset=utf-8")

	id_param := ps.ByName("id")
	if id_param == "" {
		log.Print("Didn't find `id`.", id_param)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.ParseInt(id_param, 10, 64)
	if err != nil {
		log.Print("Incorrect `id`.", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fields := r.URL.Query()["fields"]

	result, err := database.Db_get_advert_by_id(id, fields)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp_bytes, err := result.MarshalJSON()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp_bytes)
	if err != nil {
		log.Println(err)
	}

	return
}

func CreateAdvertHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	response := models.AdvertCreateResponse{}

	defer func(response *models.AdvertCreateResponse) {

		resp_bytes, err := response.MarshalJSON()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(resp_bytes)
		if err != nil {
			log.Println(err)
		}

	}(&response)

	createRequestBody := &models.AdvertCreateBody{}

	w.Header().Add("Content-type", "application/json;charset=utf-8")

	if r.Body == nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Errors = append(response.Errors, `"Request Body is nil"`)
		return
	}

	body_bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Errors = append(response.Errors, err.Error())

		return
	}

	err = createRequestBody.UnmarshalJSON([]byte(body_bytes))
	if err != nil {
		log.Println("Error unmashalling createRequestBody", err)
		w.WriteHeader(http.StatusBadRequest)
		response.Errors = append(response.Errors, err.Error())
		return
	}

	isValid := true
	error_array := make([]string, 0, 0)

	if createRequestBody.Title == "" || len(createRequestBody.Title) > 200 {
		error_array = append(error_array, `"Title must be less than 200 and not empty"`)
		isValid = false
	}
	if createRequestBody.Description == "" || len(createRequestBody.Description) > 1000 {
		error_array = append(error_array, `"Description must be less than 1000 symbols and not empty"`)
		isValid = false
	}
	if createRequestBody.Price <= 0 {
		error_array = append(error_array, `"Price Must be greater than 0"`)
		isValid = false
	}
	if len(createRequestBody.Pictures) < 1 || len(createRequestBody.Pictures) > 3 {
		error_array = append(error_array, `"Picture field must contain at least one file, maximum 3 files"`)
		isValid = false
	}
	response.Errors = error_array
	if isValid == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Id, err = database.Db_create_advert((*createRequestBody))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Errors = append(response.Errors, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}
