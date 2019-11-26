package viewes

import (
	"fmt"
)

const xml_ver string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
const open_tag string = "<sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">"
const close_tag string = "</sitemapindex>"

type IndexXML struct {
	pages int
	host string
}

func NewIndexXML(pages int, host string) *IndexXML  {
	return &IndexXML{pages, host}
}

func (n* IndexXML)XML()[]byte  {
	result := xml_ver
	result += open_tag
	for i:=1; i<=n.pages; i++{
		result += fmt.Sprintf("<sitemap><loc>%s/sitemaps/%d</loc></sitemap>", n.host, i)
	}
	result += close_tag

	return []byte(result)
}