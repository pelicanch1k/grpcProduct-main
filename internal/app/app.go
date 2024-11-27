package app

import (
	grpcapp "github.com/pelicanch1k/grpcProduct-main/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(port string) *App {
	grpcApp := grpcapp.NewApp(port)

	return &App{
		GRPCServer: grpcApp,
	}
}
