package models

import (
	"time"
)

//go get -u github.com/mailru/easyjson/...
//easyjson -all <file>.go

type AdvertResponse struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Pictures    []string  `json:"pictures,omitempty"`
	MainPicture string    `json:"main_picture"`
	Price       int64     `json:"price"`
	Date        time.Time `json:"date_created"`
}

type AdvertCreateBody struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Pictures    []string `json:"pictures"`
	Price       int64    `json:"price"`
}

type AdvertCreateResponse struct {
	Id     int64    `json:"id,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

type AdvertListElement struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	MainPicture string    `json:"main_picture"`
	Price       int64     `json:"price"`
	Date        time.Time `json:"date_created"`
}

type AdvertsList struct {
	PagesTotal  int64               `json:"pages_total"`
	Page        int64               `json:"page"`
	AdvsPerPage int64               `json:"adverts_per_page"`
	Adverts     []AdvertListElement `json:"adverts"`
}
