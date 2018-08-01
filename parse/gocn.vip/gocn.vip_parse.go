package gocn_vip

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/chunlintang/crawler/domain"
)

var (
	ErrorParse = errors.New("parse error")
)

func TitleParse(document *goquery.Document) domain.Contents {
	var (
		AllContents domain.Contents
		OneContent  domain.Content
	)
	document.Find("div.aw-common-list div.aw-item").Each(func(i int, selection *goquery.Selection) {
		var (
			user           string
			userHome       string
			url            string
			passageContent string
			title          string
			tag            string
			lastAnswer     string
			one            int
			two            int
			three          int
		)


	})
}
