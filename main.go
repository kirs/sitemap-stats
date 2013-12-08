package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Download(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	asText, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return asText
}

var counter int

func CheckSitemap(url string) {
	list := GetSitemap(url)
	for _, item := range list.Items {
    CheckUrl(item.Loc)
		counter++
	}

	fmt.Println(counter)
}

func CheckUrl(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}

  if resp.StatusCode != 200 {
    fmt.Println(url, resp.Status)
  }
}

func main() {
	url := "http://www.idinaidi.ru/sitemap.xml"

	list := GetSitemapIndex(url)

	for _, item := range list.Items {
		index := GetSitemapIndex(item.Loc)

		if len(index.Items) > 0 {
			for _, item := range index.Items {
				CheckSitemap(item.Loc)
			}
		} else {
			CheckSitemap(item.Loc)
		}
	}
}
