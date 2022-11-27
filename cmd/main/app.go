package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"user-transaction-service"
	"user-transaction-service/pkg/handler"
	"user-transaction-service/pkg/repository"
	"user-transaction-service/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("appError: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("appError: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("appError: %s", err.Error())
	}
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	server := new(user_transaction_service.Server)
	if err := server.Run(viper.GetString("port"), h.InitRoutes()); err != nil {
		logrus.Fatalf("appError: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
