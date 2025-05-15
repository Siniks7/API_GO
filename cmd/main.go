package main

import (
	"api/configs"
	"api/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	println(conf.Db.Dsn)
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
