package main

import (

	"github.com/codegangsta/negroni"
	"net/http"
	"log"
	"github.com/aleksbgs/TaskManager/common"
	"github.com/aleksbgs/TaskManager/routers"
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
