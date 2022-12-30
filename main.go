package main

import (
	"fmt"
	"net/http"

	"github.com/gabrieldebem/go-crawler/packages"
	"golang.org/x/net/html"
)

func main() {
	url := "https://www.amazon.com.br/s?k=notebook+acer+nitro+5&sprefix=notebook+acer+%2Caps%2C272&ref=nb_sb_ss_ts-doa-p_2_14"
	a := packages.AmazonCrawler{}

	a.Search(url)
//	fmt.Println(links)
}

func visit() {
	url := "https://www.amazon.com.br/gp/browse.html?node=16243890011&ref_=nav_em__wireless_smartphones_0_2_15_3"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	extractInfo(doc)
}

func extractInfo(doc *html.Node) {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for i, attr := range doc.Attr {
			fmt.Println(i)
			if attr.Key == "class" && attr.Val == "a-link-normal octopus-pc-item-link"{
				fmt.Println(attr.Val)
				
			}
		}
	}

	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		extractInfo(child)
	}
}
