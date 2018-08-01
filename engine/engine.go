package engine

import (
	"errors"
	"github.com/chunlintang/crawler/domain"
	"github.com/PuerkitoBio/goquery"
	"github.com/chunlintang/crawler/download"
	"github.com/chunlintang/crawler/parse/gocn.vip"
)

var (
	ErrorRun    = errors.New("engine run error")
	ContentsNil = domain.Contents{}
)

func Run(request domain.Request) (domain.Contents, error) {
	var (
		document *goquery.Document
		err      error
	)

	if document, err = download.Download(request.Url); err != nil {
		return ContentsNil, ErrorRun
	}

	return gocn_vip.TitleParse(document), nil
}
