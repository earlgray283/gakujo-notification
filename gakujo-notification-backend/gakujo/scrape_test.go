package gakujo

import (
	"os"
	"testing"
)

func TestScrapeReportHtml(t *testing.T) {
	f, err := os.Open("../report.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	assignment, err := scrapeReportHtml(f)
	if err != nil {
		t.Fatal(err)
	}
	if assignment.Title == "" {
		t.Fatal("failed to scrape \"title\"")
	}
}
