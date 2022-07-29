package main

import (
	"log"
	"os"
	"time"
	_ "time/tzdata"

	"gakujo-notification/server"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	f, err := os.CreateTemp("", time.Now().Format("2006-01-02_15_04_05"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	srv, err := server.New(f)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := srv.Run("8080"); err != nil {
		log.Fatal(err)
	}
}
