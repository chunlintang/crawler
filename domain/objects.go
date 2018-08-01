package domain

import "github.com/PuerkitoBio/goquery"

type Request struct {
	Url       string
	ParseFunc func(*goquery.Document) Contents
}

type Content struct {
	URL          string // 链接
	Title        string // 标题
	Tag          string // 标签
	Reporter     string // 楼主
	ReporterHome string // 主页
	Followers    int    // 关注数
	Answers      int    // 回复数
	Answerer     string // 回复人
	See          int    // 浏览数
	Passage      string // 正文
}

type Contents []Content
