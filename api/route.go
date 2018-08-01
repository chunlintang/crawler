package api

import (
	"github.com/emicklei/go-restful"
	"fmt"
	"github.com/chunlintang/crawler/engine"
	"github.com/chunlintang/crawler/domain"
	"github.com/chunlintang/crawler/parse/gocn.vip"
)

var (
	startURL = "https://gocn.vip/sort_type-new__day-0__is_recommend-0"
	rootURL  = "https://gocn.vip/sort_type-new__day-0__is_recommend-0__page-%s"
)

func GetContents(request *restful.Request, response *restful.Response) {
	var (
		page string
		url  string
	)

	page = request.PathParameter("page")
	if page == "" {
		url = startURL
	} else {
		url = fmt.Sprintf(rootURL, page)
	}

	contents, _ := engine.Run(
		domain.Request{
			Url:       url,
			ParseFunc: gocn_vip.TitleParse,
		})
	response.WriteEntity(contents)
}
