package main

import (
	"TaskManager/common"
	"TaskManager/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main() {

	common.StartUp()
	router := routers.InitRoutes()

	n:=negroni.Classic()

	n.UseHandler(router)

	server := &http.Server{
		Addr:common.AppConfig.Server,
		Handler:n,


	}
	log.Println("listening....")
	server.ListenAndServe()


}
