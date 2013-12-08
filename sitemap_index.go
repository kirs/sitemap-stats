package main

import (
	"encoding/xml"
	"log"
)

type SitemapIndex struct {
	Items []SitemapIndexItem `xml:"sitemap"`
}

type SitemapIndexItem struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func GetSitemapIndex(url string) SitemapIndex {
	asText := Download(url)

	var i SitemapIndex
	err := xml.Unmarshal([]byte(asText), &i)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
