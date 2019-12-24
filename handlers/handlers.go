package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/maxp007/avito-test-task/database"
	"github.com/maxp007/avito-test-task/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 3 метода: получение списка объявлений, получение одного объявления, создание объявления
// - валидация полей (не больше 3 ссылок на фото, описание не больше 1000 символов, название не больше 200 символов)

/*
Метод получения списка объявлений
- нужна пагинация, на одной странице должно присутствовать 10 объявлений
- нужна возможность сортировки: по цене (возрастание/убывание) и по дате создания (возрастание/убывание)
- поля в ответе: название объявления, ссылка на главное фото (первое в списке), цена*/
func GetAdvertListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Add("Content-type", "application/json")

	var page int64
	var order string
	var sort string

	str_page := r.URL.Query().Get("page")
	page, err := strconv.ParseInt(str_page, 10, 64)
	if err != nil || page == 0 {
		page = 1
	}

	order = r.URL.Query().Get("order")
	sort = r.URL.Query().Get("sort")

	if order != "date" && order != "price" {
		order = "date"
	}

	// default order by date desc
	if sort != "desc" && sort != "asc" {
		sort = "desc"
	}

	var result []models.AdvertListElement

	if order == "date" {
		result, err = database.Db_get_adverts_list_order_by_date(page, sort)
	} else {
		result, err = database.Db_get_adverts_list_order_by_price(page, sort)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp_bytes, err := (models.AdvertsList{Adverts: result}).MarshalJSON()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp_bytes)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

/*
 Метод получения конкретного объявления
- обязательные поля в ответе: название объявления, цена, ссылка на главное фото
- опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото*/
func GetAdvertHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Add("Content-type", "application/json")

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

/*
Метод создания объявления:
- принимает все вышеперечисленные поля: название, описание, несколько ссылок на фотографии (сами фото загружать никуда не требуется), цена
- возвращает ID созданного объявления и код результата (ошибка или успех)*/
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

	w.Header().Add("Content-type", "application/json")

	body_bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
		return
	}

	err = createRequestBody.UnmarshalJSON([]byte(body_bytes))
	if err != nil {
		log.Println("Error unmashalling createRequestBody", err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
		return
	}

	isValid := true
	error_array := make([]string, 0, 0)

	if createRequestBody.Title == "" || len(createRequestBody.Title) > 200 {
		error_array = append(error_array, "Title must be less than 200 and not empty")
		isValid = false
	} else if createRequestBody.Description == "" || len(createRequestBody.Description) > 1000 {
		error_array = append(error_array, "Description must be less than 1000 symbols and not empty")
		isValid = false
	} else if createRequestBody.Price == 0 {
		error_array = append(error_array, "Price Must be greater than 0")
		isValid = false
	} else if len(createRequestBody.Pictures) < 1 || len(createRequestBody.Pictures) > 3 {
		error_array = append(error_array, "Picture field must contain at least one file, maximum 3 files")
		isValid = false
	}

	if isValid == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Id, err = database.Db_create_advert((*createRequestBody))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}
