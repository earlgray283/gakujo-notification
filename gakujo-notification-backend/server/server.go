package server

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"gakujo-notification/lib"
	"gakujo-notification/repository"
	"gakujo-notification/server/worker"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type Server struct {
	app    *fiber.App
	repo   *repository.Repository
	crypto *lib.Crypto
	logger *log.Logger
}

func middlewareJwt() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY"),
	})
}

func getUserIdFromJwtToken(jwtToken *jwt.Token) uint {
	claims := jwtToken.Claims.(jwt.MapClaims)
	userId := claims["id"].(uint)
	return userId
}

func New(logWriter io.Writer) (*Server, error) {
	app := fiber.New()
	app.Use(logger.New(logger.Config{Output: logWriter}))
	crypto, err := lib.NewCrypto([]byte(os.Getenv("GAKUJO_NOTIFICATION_ENCRYPT_KEY")))
	if err != nil {
		return nil, err
	}
	repo, err := repository.New(
		os.Getenv("GAKUJO_NOTIFICATION_HOST"),
		os.Getenv("GAKUJO_NOTIFICATION_USER"),
		os.Getenv("GAKUJO_NOTIFICATION_PASSWORD"),
	)
	if err != nil {
		return nil, err
	}
	if os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY") == "" {
		return nil, errors.New("GAKUJO_NOTIFICATION_SIGNING_KEY must be set")
	}
	logger := log.New(logWriter, "[gakujo-notification]", log.Flags())
	srv := &Server{app, repo, crypto, logger}

	srv.app.Post("/auth/signup", srv.HandleSignup)
	srv.app.Post("/auth/signin", srv.HandleSignin)
	srv.app.Use(middlewareJwt()).Get("/assignments", srv.HandleGetAllAssignments)

	return srv, nil
}

func (srv *Server) Run(port string) error {
	s := gocron.NewScheduler(time.Local)

	s.Every(time.Hour).Do(worker.CrawleAssignments(srv.repo, srv.crypto))
	s.StartAsync()

	return srv.app.Listen(":" + port)
}
