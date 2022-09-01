package main

import (
	handler_impl "backend/app/handler"
	"backend/app/middleware"
	repository_impl "backend/app/repository/impl"
	usecase_impl "backend/app/usecase/impl"
	"backend/cmd/api/config"
	handler_email "backend/cmd/email/handler"
	handler_session "backend/cmd/session/handler"
	"backend/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
	"strings"
	"time"
)

func initDB(cfg config.Config, log *logrus.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		strings.Join([]string{"host=", cfg.Db.Host, " user=", cfg.Db.User, " password=", cfg.Db.Pass, " dbname=", cfg.Db.Name, " port=", cfg.PostgresPort}, "")), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Run() {
	logger := logrus.New()
	cfg, err := config.ParseConfig("cmd/api/config/config.json")
	if err != nil {
		logger.Errorf("Error while parsing config: %s", err.Error())
		return
	}

	f, err := os.Create(cfg.LogFile)
	if err != nil {
		logger.Errorf("Error while creating log file: %s", err.Error())
		return
	}

	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	time.Sleep(10 * time.Second)
	db, err := initDB(cfg, logger)
	if err != nil {
		logger.Errorf("Error while init Db: %s", err.Error())
		return
	}

	grpcConn, err := grpc.Dial(
		cfg.SessionContainer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Errorf("Error while grpcConn to sessionContainer: %s", err.Error())
		return
	}

	sessService := repository_impl.CreateSessionRepo(handler_session.NewSessionCheckerClient(grpcConn))

	grpcConnEmail, err := grpc.Dial(
		cfg.EmailContainer,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Errorf("Error while grpcConn to sessionContainer: %s", err.Error())
		return
	}

	emailService := repository_impl.CreateEmailRepo(handler_email.NewEmailCheckerClient(grpcConnEmail))

	userRepo := repository_impl.CreateUserRep(db)

	authMiddleware := middleware.CreateMiddleware(sessService)
	router.Use(middleware.CheckError())
	fmt.Println("CFG.EXPIRATION: ", cfg.Expiration)
	userHandler := handler_impl.MakeUserHandler(usecase_impl.MakeUserUsecase(userRepo, sessService, emailService), cfg.Expiration)
	mainRoutes := router.Group(cfg.Urls.RootUrl)
	{
		mainRoutes.POST(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.LoginUrl}, ""), userHandler.Login)
		mainRoutes.POST(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.RegisterUrl}, ""), userHandler.Register)
		mainRoutes.DELETE(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.LogoutUrl}, ""), userHandler.Logout)
		mainRoutes.GET(cfg.Urls.ProfileUrl, authMiddleware.CheckAuth, userHandler.GetInfoByCookie)
		mainRoutes.PUT(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.UploadUrl}, ""), authMiddleware.CheckAuth, userHandler.SaveAvatar)
		mainRoutes.PUT(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.RefactorUrl}, ""), authMiddleware.CheckAuth, userHandler.RefactorProfile)
		mainRoutes.POST(strings.Join([]string{cfg.Urls.ProfileUrl, cfg.Urls.VerifyUrl}, ""), authMiddleware.CheckAuth, userHandler.VerificationProfile)
	}
	err = router.Run(strings.Join([]string{"", cfg.ServerPort}, ":"))
	if err != nil {
		logger.Errorf("Error while starting server: %s", err.Error())
	}
}
