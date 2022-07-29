package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"gakujo-notification/server"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	srv, err := server.New(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.Run("8080"); err != nil {
		log.Fatal(err)
	}
}
