package main

import (
	todo "MaksJash"
	"MaksJash/pkg/handler"
	"log"
)

// Запуск проги

func main() {

	server := new(todo.Server)
	router := new(handler.Handler) // Не знаю, в чем отличие между handler.Handler{}

	if err := server.Run("9090", router.InitRoutes()); err != nil {
		log.Fatalf("Error while running the server: %s", err.Error()) // .Error() - приведение к строке
	}

}
