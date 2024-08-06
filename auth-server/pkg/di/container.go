package di

import (
	"grpc-auth-app/auth-server/internal/handler"
	"grpc-auth-app/auth-server/internal/repository"
	"grpc-auth-app/auth-server/internal/service"
	"grpc-auth-app/auth-server/pkg/auth"
	"grpc-auth-app/auth-server/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
)

func BuildContainer() *handler.AuthServer {
	repo := repository.NewUserRepository()
	expirationHrs := config.Config.TokenExpiryHour
	// log.Println("expiration Hours", expirationHrs)
	expirationDuration := time.Duration(expirationHrs) * time.Hour
	tokenSvc := auth.NewAuthService(&auth.AuthConf{
		ExpirationDuration: expirationDuration,
		SecretKey:          []byte(config.Config.TokenSecret),
	})
	svc := service.NewUserService(&service.USConf{
		UserRepo: repo,
		TokenSvc: tokenSvc,
	})
	router := gin.Default()
	authServer := handler.NewAuthServer(svc, router)
	// go handler.StartHttpServer(authServer)
	return authServer
}
