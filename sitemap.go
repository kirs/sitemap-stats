package main

import (
	"encoding/xml"
	"log"
)

type Sitemap struct {
	Items []SitemapItem `xml:"url"`
}

type SitemapItem struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func GetSitemap(url string) Sitemap {
	asText := Download(url)

	var i Sitemap
	err := xml.Unmarshal([]byte(asText), &i)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
