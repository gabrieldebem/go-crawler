package services

import (
	"fmt"
	s "strings"

	"github.com/gabrieldebem/go-crawler/packages/models"
	"github.com/gocolly/colly/v2"
)

type AmazonCrawler struct {
	products []models.Product
}

func (a AmazonCrawler) Search(url string) []models.Product {
	prefix := "https://www.amazon.com.br"

	c := colly.NewCollector()

	var link string
	var name string
	var price string

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "s-card-container s-overflow-hidden aok-relative puis-expand-height puis-include-content-margin puis s-latency-cf-section s-card-border" {
			e.ForEachWithBreak("a", func(_ int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "a-link-normal s-underline-text s-underline-link-text s-link-style a-text-normal" && e.Attr("href") != "" {
					if s.Contains(e.ChildAttr("a", "href"), prefix) {
						link = e.Attr("href")
					} else {
						link = prefix + e.Attr("href")
					}
					return false
				}
				return true
			})

			e.ForEachWithBreak("span", func(i int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "a-size-base-plus a-color-base a-text-normal" {
					title := e.Text
					name = title
					return false
				}
				return true
			})

			e.ForEachWithBreak("span", func(i int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "a-price-whole" {
					price = s.Replace(s.Replace(e.Text, ",", "", -1), ".", "", -1)
					return false
				}
				return true
			})

			if name != "" && price != "" && link != "" {
				product := models.Product{
					Name:  name,
					Price: "R$ " + price,
					Link:  link,
				}

				a.products = append(a.products, product)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return a.products
}
