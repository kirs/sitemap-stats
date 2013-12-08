package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Sitemap struct {
	List []Url `xml:"url"`
}

type Url struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func GetSitemap() Sitemap {
	xmlFile, err := os.Open("rent2.xml")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer xmlFile.Close()

	var i Sitemap
	err = xml.NewDecoder(xmlFile).Decode(&i)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func main() {
	list := GetSitemap()

	fmt.Println(len(list.List))
}
