package routes

import (
    //"net/http"
   
    "Registration/handlers"
)

func InitRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/register", handlers.Register).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")
    return r
}
