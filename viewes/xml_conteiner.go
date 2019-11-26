package viewes

import "sitemap/models/entity"

const xml_ver_sitemaps string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
const open_tag_sitemaps string = "<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">"
const close_tag_sitemaps string = "</urlset>"

func XMLConteiner(url string,xml *[]entity.ResQuery)[]byte{
	result := xml_ver_sitemaps + open_tag_sitemaps
	for _, entity := range *xml {
		result += *entity.XML(url)
	}
	result += close_tag_sitemaps
	return []byte(result)
}