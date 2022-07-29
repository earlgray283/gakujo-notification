package gakujo

import (
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type Assignment struct {
	Kind        AssignmentKind
	SubjectName string
	Semester    Semester
	Status      AssignmentStatus
	Title       string
	Since       time.Time
	Until       time.Time
	Description string
	Message     string
	Year        int
}

// レポートの提出一覧を取得する
func (c *Client) ReportAssignments() ([]*Assignment, error) {
	year, err := c.processAccessReportPageAndGetYear()
	if err != nil {
		return nil, err
	}

	assignments := make([]*Assignment, 0)
	for row := 1; ; row++ {
		assignmentSelector := fmt.Sprintf("#searchList > tbody > tr:nth-child(%d)", row)
		var text string
		if err := chromedp.Run(c.ctx, chromedp.Text(assignmentSelector, &text, chromedp.ByQuery, chromedp.AtLeast(0))); err != nil {
			break
		}
		if text == "" {
			break
		}

		var subjectText, statusText, html string
		if err := chromedp.Run(c.ctx,
			chromedp.Text(assignmentSelector+" > td:nth-child(1)", &subjectText, chromedp.NodeVisible),
			chromedp.Text(assignmentSelector+" > td:nth-child(3)", &statusText, chromedp.NodeVisible),
			chromedp.Click(assignmentSelector+" > td:nth-child(2) > a"),      // 課題詳細のリンクをクリック
			chromedp.OuterHTML("#area > table", &html, chromedp.NodeVisible), // 課題詳細ページの html を取得
		); err != nil {
			return nil, err
		}

		assignment, err := scrapeReportHtml(strings.NewReader(html))
		if err != nil {
			return nil, err
		}

		// 科目名・学期の取得
		subjectName, semester, err := parseSubjectFormat(subjectText)
		if err != nil {
			return nil, err
		}
		assignment.Status = NewAssignmentStatus(statusText)
		assignment.SubjectName = subjectName
		assignment.Semester = semester
		assignment.Year = year
		assignments = append(assignments, assignment)

		fmt.Println(assignment.SubjectName, assignment.Title)

		if err := chromedp.Run(c.ctx,
			chromedp.Click("#gnav-menu > li:nth-child(2) > a", chromedp.NodeVisible),
			chromedp.SetValue("#searchList_length > label > select", "-1", chromedp.NodeVisible),
		); err != nil {
			return nil, err
		}

		// 学情に負荷をかけないように、1sec ごとに処理を行う
		time.Sleep(time.Second)
	}

	return assignments, nil
}
