package gakujo

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const DateFormat = "2006/01/02 15:04"

func scrapeReportHtml(r io.Reader) (*Assignment, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	tbodySection := doc.Find("tbody")

	title := tbodySection.Find("tr:nth-child(1) > td").Text()
	dateRangeText := strings.TrimSpace(tbodySection.Find("tr:nth-child(2) > td").Text())
	tokens := strings.Split(dateRangeText, " ～ ")
	sinceText, untilText := tokens[0], tokens[1]
	since, _ := time.Parse(DateFormat, sinceText)
	until, _ := time.Parse(DateFormat, untilText)
	description := tbodySection.Find("tr:nth-child(4) > td").Text()
	message := tbodySection.Find("tr:nth-child(6) > td").Text()

	return &Assignment{
		Kind:        AssignmentKindReport,
		Title:       title,
		Since:       since,
		Until:       until,
		Description: description,
		Message:     message,
	}, nil
}

func scrapeMinitestHtml(r io.Reader) (*Assignment, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	title := doc.Find("tr:nth-child(1) > td").Text()
	dateRangeText := doc.Find("tr:nth-child(2) > td").Text()
	sinceText, untilText := "", ""
	fmt.Sscanf(dateRangeText, "%s ～ %s", &sinceText, &untilText)
	since, _ := time.Parse(DateFormat, sinceText)
	until, _ := time.Parse(DateFormat, untilText)
	description := doc.Find("tr:nth-child(5) > td").Text()
	message := doc.Find("tr:nth-child(7) > td").Text()

	return &Assignment{
		Kind:        AssignmentKindMinitest,
		Title:       title,
		Since:       since,
		Until:       until,
		Description: description,
		Message:     message,
	}, nil
}

// 授業名・学期を取得
func parseSubjectFormat(s string) (string, Semester, error) {
	s = strings.TrimSpace(s)
	tokens := strings.Split(s, "\n")
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("invalid subject name format: ===%s===", s)
	}
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}
	subjectName := tokens[0]
	semester := NewSemester(strings.Split(tokens[1], "/")[0])

	return subjectName, semester, nil
}
