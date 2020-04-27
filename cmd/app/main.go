package main

import (
	"fmt"
	"github.com/wilian746/go-crud-dynamodb/config"
	"github.com/wilian746/go-crud-dynamodb/internal/repository/adapter"
	"github.com/wilian746/go-crud-dynamodb/internal/repository/instance"
	"github.com/wilian746/go-crud-dynamodb/internal/routes"
	"log"
	"net/http"
)

func main() {
	configs := config.GetConfig()

	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	log.Println("service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}
