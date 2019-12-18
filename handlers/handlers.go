package handlers

import (
	"github.com/gorilla/mux"
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

func GetAdvertListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("Unsupported request method: ", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-type", "application/json")

	var page int64
	var date_sort string
	var price_sort string

	str_page := r.URL.Query().Get("page")

	page, err := strconv.ParseInt(str_page, 10, 64)
	if err != nil || page == 0 {
		page = 1
	}

	date_sort = r.URL.Query().Get("date_sort")
	if date_sort == "" {
		date_sort = "none"
	}

	price_sort = r.URL.Query().Get("price_sort")
	if price_sort == "" {
		price_sort = "none"
	}

	//Process Data here
	w.WriteHeader(http.StatusOK)
	return
}

/*
 Метод получения конкретного объявления
- обязательные поля в ответе: название объявления, цена, ссылка на главное фото
- опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото*/

func GetAdvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Unsupported request method: ", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Content-type", "application/json")

	vars := mux.Vars(r)
	str_id, found := vars["id"]
	if !found {
		log.Print("Didn't find `id`.", str_id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		log.Print("Incorrect `id`.", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fields := r.URL.Query()["fields"]

	log.Print(fields)

	//Process Data here
	w.WriteHeader(http.StatusOK)
	return
}

/*
Метод создания объявления:
- принимает все вышеперечисленные поля: название, описание, несколько ссылок на фотографии (сами фото загружать никуда не требуется), цена
- возвращает ID созданного объявления и код результата (ошибка или успех)*/

func CreateAdvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Unsupported request method: ", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	createRequestBody := &models.AdvertCreateBody{}

	body_bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = createRequestBody.UnmarshalJSON([]byte(body_bytes))
	if err != nil {
		log.Println("Error unmashalling createRequesBody", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Process Structure here
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
