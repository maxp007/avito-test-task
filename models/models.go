package models

//go get -u github.com/mailru/easyjson/...
//easyjson -all <file>.go

type AdvertResponse struct {
	Id          int64    `json:"id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Pictures    []string `json:"pictures,omitempty"`
	MainPicture string   `json:"main_picture"`
	Price       int64    `json:"price"`
}

type AdvertCreateBody struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Pictures    []string `json:"pictures"`
	Price       int64    `json:"price"`
}

type AdvertCreateResponse struct {
	Id    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type AdvertListElement struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	MainPicture string `json:"main_picture"`
	Price       int64  `json:"price"`
}

type AdvertsList struct {
	Adverts []AdvertListElement `json:"adverts"`
}
