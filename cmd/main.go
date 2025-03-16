package main

import (
	todo "MaksJash"
	"MaksJash/pkg/handler"
	"MaksJash/pkg/repository"
	"MaksJash/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Запуск проги

func main() {

	// Кастомизация форматтера: 1) Теперь логи в жсоне, 2) С индентами
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	if err := initConfig(); err != nil {
		logrus.Errorf("Error with config initialization: %s", err.Error()) // .Error() - приведение к строке
	}

	// Ищет .env файл в текущей директории, если параметрами не переданы другие директории
	// И грузит себе во внутрянку
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error with loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Error with database configuration: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	srvc := service.NewService(rep)
	handler := handler.NewHandler(srvc)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Errorf("Error while running the server: %s", err.Error())
	}

}

// Инициализация конфига
func initConfig() error {

	viper.AddConfigPath("configs") // Указание директории, где будет искать файл конфига
	viper.SetConfigName("config")  // Имя файла конфига, что нужно найти
	viper.SetConfigType("yaml")    // Расширение для файла, что мы ищем

	return viper.ReadInConfig()
}
