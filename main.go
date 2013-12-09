package main

import (
	"fmt"
	"flag"
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

var request_stats = map[int]int{}

func CheckSitemap(url string) {
	fmt.Printf("Parsing %s\n", url)
	list := GetSitemap(url)
	for _, item := range list.Items {
    CheckUrl(item.Loc)
		counter++
	}

	fmt.Printf("Total urls: %d\n", counter)

	PrintCodeStats()
}

func CheckUrl(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Println(url, resp.Status)
  request_stats[resp.StatusCode]++
}

func PrintCodeStats() {
	for code, count := range request_stats {
		fmt.Printf("%d: %d\n", code, count)
	}
}

var sitemap_url string
func init() {
	flag.StringVar(&sitemap_url, "url", "", "sitemap url")
}

func main() {
	flag.Parse()

	if len(sitemap_url) == 0 {
		log.Fatal("Sitemap url is empty")
	}

	fmt.Printf("Parsing %s started\n", sitemap_url)

	list := GetSitemapIndex(sitemap_url)

	if len(list.Items) == 0 {
		CheckSitemap(sitemap_url)
		return
	}

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

	PrintCodeStats()
}
