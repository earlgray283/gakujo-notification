package gakujo

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var client *Client

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	var err error
	client, err = NewClient(context.Background(), os.Getenv("GAKUJO_ID"), os.Getenv("GAKUJO_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
}

func TestAssignments(t *testing.T) {
	assignments, err := client.ReportAssignments()
	if err != nil {
		t.Fatal(err)
	}
	for _, assignmant := range assignments {
		t.Log(assignmant)
	}
}
