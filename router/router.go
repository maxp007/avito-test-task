package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/maxp007/avito-test-task/handlers"
)

func init() {
}

func GetRouter() (r *httprouter.Router) {

	r = httprouter.New()

	r.GET("/api/adverts", handlers.GetAdvertListHandler)
	r.GET("/api/advert/:id/", handlers.GetAdvertHandler)

	r.POST("/api/create", handlers.CreateAdvertHandler)

	return
}
