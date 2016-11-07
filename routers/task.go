package routers

import (
	"github.com/gorilla/mux"
	"TaskManager/controllers"
	"github.com/codegangsta/negroni"
	"TaskManager/common"
)

func SetTaskRoute(router * mux.Router) * mux.Router{

	taskRouter := mux.NewRouter()

	taskRouter.HandleFunc("/tasks",controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}",controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks",controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}",controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/users/{id}",controllers.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}",controllers.DeleteTask).Methods("DELETE")

	router.PathPrefix("/tasks").Handler(negroni.New(negroni.HandlerFunc(common.Authorize),negroni.Wrap(taskRouter),))

	return router



}

