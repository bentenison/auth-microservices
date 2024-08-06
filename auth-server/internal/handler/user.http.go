package handler

import (
	"context"
	"grpc-auth-app/auth-server/api"
	"grpc-auth-app/auth-server/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *AuthServer) SigninHttpHandler(c *gin.Context) {
	var payload api.UserRequest
	if err := c.BindJSON(&payload); err != nil {
		log.Println("error while binding payload", err)
		return
	}
	user, err := a.userSvc.SigninService(c.Request.Context(), &payload)
	if err != nil {
		log.Println("error while login user", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
func (a *AuthServer) SignupHttpHandler(c *gin.Context) {
	var payload api.User
	if err := c.BindJSON(&payload); err != nil {
		log.Println("error while binding payload", err)
		return
	}
	res, err := a.userSvc.SignUpService(c.Request.Context(), &payload)
	if err != nil {
		log.Println("error while registering user", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *AuthServer) DeleteUserHttpHandler(c *gin.Context) {
	var payload api.DeleteUserRequest
	if err := c.BindJSON(&payload); err != nil {
		log.Println("error while binding payload", err)
		return
	}
	res, err := a.userSvc.DeleteUserService(c.Request.Context(), &payload)
	if err != nil {
		log.Println("error while deleting user", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *AuthServer) ListUserHttpHandler(c *gin.Context) {
	var payload api.ListUserRequest
	if err := c.BindJSON(&payload); err != nil {
		log.Println("error while binding payload", err)
		return
	}
	res, err := a.userSvc.ListUsersService(c.Request.Context(), &payload)
	if err != nil {
		log.Println("error while listing users", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func RegisterHttpServer(h *AuthServer) {
	// api.RegisterUserServiceServer(grpcServer, svcServer)
	h.router.GET("/check", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status":     "running",
			"serverTime": time.Now(),
			"handled-by": os.Getenv("POD_NAME"),
			"on":         os.Getenv("POD_NAMESPACE"),
		})
	})
	h.router.Group("/user")
	h.router.POST("/signin", h.SigninHttpHandler)
	h.router.POST("/signup", h.SignupHttpHandler)
	h.router.POST("/delete", h.DeleteUserHttpHandler)
	h.router.POST("/list", h.ListUserHttpHandler)
}

func StartHttpServer(h *AuthServer) {
	RegisterHttpServer(h)
	serverPort := config.Config.HTTPPort
	svr := http.Server{
		Addr:    serverPort,
		Handler: h.router,
	}
	log.Printf("http server started on port %s\n", serverPort)
	// go func() {
	//server alredy started in a goroutne from its parent function
	if err := svr.ListenAndServe(); err != nil {
		log.Println("error in starting http server.. fallback to grpc", err)
		return
	}
	// }()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Println("error in shutting down the server, server misehaving", err)
		// return
	}
}
