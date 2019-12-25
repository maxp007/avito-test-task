package handlers

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"github.com/maxp007/avito-test-task/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAdvertListWrongPageParam(t *testing.T) {

	cases := []struct {
		Name  string
		order string
		sort  string
		page  string
	}{
		{"page param is not a number", "price", "asc", "not_a_number"},
		{"page param doesn't exist", "price", "asc", ""},
		{"page param is less than 0", "price", "asc", "-1"},
		{"page param is equal to  0", "price", "asc", "0"},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/api/adverts", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("order", tt.order)
			q.Add("sort", tt.sort)
			q.Add("page", tt.page)
			req.URL.RawQuery = q.Encode()

			r := httprouter.New()
			r.GET("/api/adverts", GetAdvertListHandler)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}

func TestCreateAdvertHandler_MissingBody(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	router := httprouter.New()
	router.POST("/api/create", CreateAdvertHandler)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateAdvertHandlerValidation(t *testing.T) {
	cases := []struct {
		Name string
		Body models.AdvertCreateBody
	}{

		{"empty adv title", models.AdvertCreateBody{"", "description field", []string{"pic.com/pic1.jpg", "pic.com/pic2.jpg", "pic.com/pic3.jpg"}, 1000}},
		{"empty adv description", models.AdvertCreateBody{"title", "", []string{"pic.com/pic1.jpg", "pic.com/pic2.jpg", "pic.com/pic3.jpg"}, 1000}},
		{"empty adv pictures", models.AdvertCreateBody{"title", "description field", []string{}, 1000}},
		{"empty adv price", models.AdvertCreateBody{"title", "description field", []string{"pic.com/pic1.jpg", "pic.com/pic2.jpg", "pic.com/pic3.jpg"}, 0}},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			body_bytes, err := tt.Body.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("POST", "/api/create", bytes.NewReader(body_bytes))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

			router := httprouter.New()
			router.POST("/api/create", CreateAdvertHandler)

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}
