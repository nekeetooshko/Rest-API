package main

import (
	todo "MaksJash"
	"MaksJash/pkg/handler"
	"MaksJash/pkg/repository"
	"MaksJash/pkg/service"
	"log"
)

// Запуск проги

func main() {

	rep := repository.NewRepository()
	srvc := service.NewService(rep)
	handler := handler.NewHandler(srvc)

	server := new(todo.Server)

	if err := server.Run("9090", handler.InitRoutes()); err != nil {
		log.Fatalf("Error while running the server: %s", err.Error()) // .Error() - приведение к строке
	}

}
