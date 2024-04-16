package main

import (
	"fis/configs"
	"fis/internal/handler"
	"fis/internal/pb"
	"fis/internal/repository"
	"fis/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/skbt-ecom/logging"
	"google.golang.org/grpc"
	"net"
)

func main() {
	log := logging.InitLogger()

	var cfg configs.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.WithError(err).Fatal("failed to read env")
	}

	log.AddGraylogHook(cfg.GraylogUrl, "fis-service")

	// Clean Architecture Pattern
	repo := repository.NewRepository(&cfg, log.WithExtraField("layer", "repository"))
	svc := service.NewService(repo, log.WithExtraField("layer", "service"))
	s := handler.NewHandler(svc, log.WithExtraField("layer", "handler"))

	// GRPC Server
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.WithError(err).Fatalf("failed to listen %s port", cfg.Port)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFisServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.WithError(err).Fatalf("failed to serve %s port", cfg.Port)
	}
}
