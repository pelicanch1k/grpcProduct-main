package main

import (
	"github.com/pelicanch1k/grpcProduct-main/internal/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	application := app.NewApp(viper.GetString("port"))

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	logrus.Info("Gracefully stopped")
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
