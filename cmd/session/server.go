package main

import (
	"backend/cmd/session/config"
	"backend/cmd/session/handler"
	"backend/cmd/session/handler/impl"
	"backend/cmd/session/repository"
	"backend/cmd/session/usecase"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func Run() {
	log := logrus.New()
	cfg, err := config.ParseConfig("cmd/session/config/config.json")
	if err != nil {
		log.Errorf("Error while parsing config: %s", err.Error())
	}
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisContainer,
		Password: cfg.RedisPass,
		DB:       0,
	})

	sessionRepo := repository.CreateSessRepo(redisCli)
	sessionUseCase := usecase.CreateSessionUseCase(sessionRepo)
	sessionHandler := impl.CreateSessionServer(sessionUseCase, int(cfg.CookieExpiration*60*60*24))

	listener, err := net.Listen("tcp", cfg.SessionContainer)
	if err != nil {
		log.Errorf("Error while net.Listen: %s", err.Error())
		return
	}

	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterSessionCheckerServer(srv, sessionHandler)
	logrus.Println("Successfully started session microservice")
	if err = srv.Serve(listener); err != nil {
		log.Errorf("Error while grpc.Serve(listener): %s", err.Error())
		return
	}
}
