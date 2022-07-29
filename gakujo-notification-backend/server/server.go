package server

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"gakujo-notification/lib"
	"gakujo-notification/repository"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	pkglogger "github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type Server struct {
	app    *fiber.App
	repo   *repository.Repository
	crypto *lib.Crypto
	logger *log.Logger
}

func middlewareJwt(signingKey []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: signingKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println("!", err)
			return fiber.DefaultErrorHandler(c, err)
		},
	})
}

func getUserIdFromJwtToken(jwtToken *jwt.Token) string {
	claims := jwtToken.Claims.(jwt.MapClaims)
	log.Println(claims)
	userId := claims["id"].(string)
	return userId
}

func New(logWriter io.Writer) (*Server, error) {
	if os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY") == "" || os.Getenv("GAKUJO_NOTIFICATION_ENCRYPT_KEY") == "" {
		return nil, errors.New("environment value GAKUJO_NOTIFICATION_* must be set")
	}

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
	logger := log.New(logWriter, "[gakujo-notification]", log.Flags())

	app := fiber.New()
	app.Use(pkglogger.New(pkglogger.Config{Output: logWriter}))
	srv := &Server{app, repo, crypto, logger}

	srv.app.Post("/auth/signup", srv.HandleSignup)
	srv.app.Post("/auth/signin", srv.HandleSignin)
	srv.app.Use(middlewareJwt([]byte(os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY")))).
		Get("/assignments", srv.HandleGetAllAssignments)

	return srv, nil
}

func (srv *Server) Run(port string) error {
	s := gocron.NewScheduler(time.Local)

	//s.Every(time.Hour).Do(worker.CrawleAssignments(srv.repo, srv.crypto))
	s.StartAsync()

	return srv.app.Listen(":" + port)
}
