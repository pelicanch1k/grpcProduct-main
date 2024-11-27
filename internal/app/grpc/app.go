package grpcapp

import (
	"fmt"
	server "github.com/pelicanch1k/grpcProduct-main/internal/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func NewApp(port string) *App {
	grpcServer := grpc.NewServer()

	server.Register(grpcServer)
	return &App{port: port, gRPCServer: grpcServer}
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err.Error())
	}
}

func (a *App) run() error {
	const op = "internal/app/grpc/app.go Run()"

	lis, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	logrus.Info("starting grpc server at " + a.port)

	if err := a.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "internal/app/grpc/app.go Stop()"

	logrus.Info("stopping grpc server at " + a.port)

	a.gRPCServer.GracefulStop()
}
