package models

//go get -u github.com/mailru/easyjson/...
//easyjson -all <file>.go

type AdvertResponse struct {
	Id          int64    `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pictures    []string `json:"pictures,omitempty"`
	MainPicture string   `json:"main_photo,omitempty"`
	Price       int64    `json:"price"`
}

type AdvertCreateBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pictures    []string `json:"pictures,omitempty"`
	Price       int64    `json:"price"`
}
