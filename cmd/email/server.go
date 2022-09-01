package main

import (
	"backend/cmd/email/config"
	"backend/cmd/email/handler"
	"backend/cmd/email/handler/impl"
	"backend/cmd/email/repository"
	"backend/cmd/email/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"gopkg.in/gomail.v2"
	"net"
	"time"
)

func Run() {
	log := logrus.New()
	cfg, err := config.ParseConfig("cmd/email/config/config.json")
	if err != nil {
		log.Errorf("Error while parsing config: %s", err.Error())
	}
	mailDealer := gomail.NewDialer(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUsername,
		cfg.MailPassword,
	)
	emailRepo := repository.CreateEmailRepo(mailDealer)
	emailUseCase := usecase.CreateEmailUseCase(emailRepo, cfg.MailUsername)
	emailHandler := impl.CreateEmailServer(emailUseCase)

	listener, err := net.Listen("tcp", cfg.EmailContainer)
	if err != nil {
		log.Errorf("Error while net.Listen: %s", err.Error())
		return
	}

	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	handler.RegisterEmailCheckerServer(srv, emailHandler)
	logrus.Println("Successfully started email microservice")
	if err = srv.Serve(listener); err != nil {
		log.Errorf("Error while grpc.Serve(listener): %s", err.Error())
		return
	}
}
