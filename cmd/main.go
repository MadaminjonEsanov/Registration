package main

import (
    "fmt"
    "log"
    "net/http"
    "registration/config"
    "registration/handlers"
    "registration/routes"
  
)

func main() {
    conf := config.NewConfig()
    dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
        conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)
    handlers.InitDB(dbConnection)

    r := routes.InitRoutes()

    log.Fatal(http.ListenAndServe(conf.HttpPort, r))
}
