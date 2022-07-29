package gakujo

import (
	"strconv"

	"github.com/chromedp/chromedp"
)

// chrome で学情にログインする
func (c *Client) processLogin() error {
	return chromedp.Run(c.ctx,
		chromedp.Navigate("https://gakujo.shizuoka.ac.jp/portal/"),
		chromedp.Click("#left_container > div.left-module-top.bg_color > div > div > a", chromedp.NodeVisible),
		chromedp.SendKeys("#username", c.id),
		chromedp.SendKeys("#password", c.password),
		chromedp.Click("body > div > div > div > div > form > div:nth-child(3) > button", chromedp.NodeVisible),
	)
}

// レポートページにアクセスする。また、年を取得する。
func (c *Client) processAccessReportPageAndGetYear() (int, error) {
	if err := c.processLogin(); err != nil {
		return 0, err
	}

	var schoolYearText string
	if err := chromedp.Run(c.ctx,
		chromedp.Click("#header-menu > li.header-menulist.menu > a", chromedp.NodeVisible),
		chromedp.Click("#header-menu-sub > li:nth-child(1) > a", chromedp.NodeVisible),
		chromedp.Click("#gnav-menu > li:nth-child(2) > a", chromedp.NodeVisible),
		chromedp.SetValue("#searchList_length > label > select", "-1", chromedp.NodeVisible),
		chromedp.Value("#schoolYear", &schoolYearText, chromedp.ByID),
	); err != nil {
		return 0, err
	}

	year, err := strconv.Atoi(schoolYearText)
	if err != nil {
		return 0, err
	}

	return year, nil
}
