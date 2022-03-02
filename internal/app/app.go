package app

import (
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
	"wblvl0/internal/handler"
	"wblvl0/internal/nuts"
	"wblvl0/internal/repository"
	"wblvl0/internal/service"
	"wblvl0/pkg/db"
)

func Run() {

	if err := initConfig(); err != nil {
		log.Fatalf("init config error: %s", err.Error())
	}

	dataBase, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("init db error: %s", err.Error())
	}

	newRepos := repository.NewRepository(dataBase)
	newService := service.NewService(newRepos)
	newHandler := handler.NewHandler(newService)

	natsConnection, err := stan.Connect(viper.GetString("nuts.stanclusterid"), viper.GetString("nuts.clientid"))
	if err != nil {
		log.Fatalf("nuts connection error: %s", err.Error())
	}
	defer func(natsConnection stan.Conn) {
		err = natsConnection.Close()
		if err != nil {
			log.Printf("close nuts error: %s", err.Error())
		}
	}(natsConnection)
	nuts.NewNutsSubscriber(natsConnection, newRepos)

	server := new(Server)
	err = server.Run(viper.GetString("httpserver.port"), newHandler.InitRoutes())
	if err != nil {
		log.Fatalf("start http server error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
