package gocn_vip

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/chunlintang/crawler/domain"
	"github.com/chunlintang/crawler/util"
	"github.com/chunlintang/crawler/download"
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
		user = util.StringSplit(selection.Find("a").Eq(0).AttrOr("href", "None"))
		userHome = selection.Find("a").Eq(0).AttrOr("href", "None")
		url = selection.Find("div h4 a").AttrOr("href", "None")
		doc, _ := download.Download(url)
		passageContent, _ = PassageParse(doc)
		title = util.StringTrim(selection.Find("div h4").Text())
		tag = selection.Find("p a").Eq(0).Text()
		lastAnswer = selection.Find("p a").Eq(1).Text()
		one, two, three = util.StringSplitByDot(selection.Find("p span").Eq(0).Text())

		OneContent = domain.Content{
			URL:          url,
			Title:        title,
			Tag:          tag,
			Reporter:     user,
			ReporterHome: userHome,
			Followers:    one,
			Answers:      two,
			Answerer:     lastAnswer,
			See:          three,
			Passage:      passageContent,
		}

		AllContents = append(AllContents, OneContent)
	})

	return AllContents
}

func PassageParse(document *goquery.Document) (string, error) {
	var content string
	content = util.StringTrim(document.Find("div.content.markitup-box").Text())
	return util.StringReplace(content), nil
}
