package main

import (
	todo "MaksJash"
	"MaksJash/pkg/handler"
	"MaksJash/pkg/repository"
	"MaksJash/pkg/service"
	"log"

	"github.com/spf13/viper"
)

// Запуск проги

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error with config initialization: %s", err.Error()) // .Error() - приведение к строке
	}

	rep := repository.NewRepository()
	srvc := service.NewService(rep)
	handler := handler.NewHandler(srvc)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}

}

// Инициализация конфига
func initConfig() error {

	viper.AddConfigPath("configs") // Указание директории, где будет искать файл конфига
	viper.SetConfigName("config")  // Имя файла конфига, что нужно найти

	return viper.ReadInConfig()
}
