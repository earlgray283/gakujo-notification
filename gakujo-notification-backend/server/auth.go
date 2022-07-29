package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gakujo-notification/lib"
	"gakujo-notification/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (srv *Server) HandleSignup(c *fiber.Ctx) error {
	username := c.FormValue("username")
	if !lib.ValidateString(username, lib.MinLen(4), lib.MaxLen(16), lib.WhiteList(lib.AlphanumericCharacters)) {
		return c.Status(http.StatusBadRequest).SendString("username is invalid")
	}
	password := c.FormValue("password")
	if !lib.ValidateString(password, lib.MinLen(8), lib.MaxLen(128), lib.WhiteList(lib.AlphanumericCharacters)) {
		return c.Status(http.StatusBadRequest).SendString("password is invalid")
	}
	gakujoId := c.FormValue("gakujoId")
	gakujoPassword := c.FormValue("gakujoPassword")

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	encryptedGakujoAccount, err := srv.crypto.Encrypt([]byte(fmt.Sprintf("%s&%s", gakujoId, gakujoPassword)))
	if err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	newUser := repository.NewUser(username, string(encryptedPassword), encryptedGakujoAccount)
	if err := srv.repo.RunInTransaction(func(tx *gorm.DB) error {
		return tx.Create(newUser).Error
	}); err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       newUser.ID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY")))
	if err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Send([]byte(tokenString))
}

func (srv *Server) HandleSignin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := srv.repo.FetchUserByUsername(username)
	if err != nil {
		srv.logger.Println(err)
		switch err {
		case gorm.ErrRecordNotFound:
			return c.SendStatus(http.StatusNotFound)
		default:
			return c.SendStatus(http.StatusInternalServerError)
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)); err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("GAKUJO_NOTIFICATION_SIGNING_KEY")))
	if err != nil {
		srv.logger.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Send([]byte(tokenString))
}
