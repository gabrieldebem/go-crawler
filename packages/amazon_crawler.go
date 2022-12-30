package packages

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
			fmt.Println("achou algo")

			if e.ChildAttr("a", "href") != "" {
				if s.Contains(e.ChildAttr("a", "href"), prefix) {
					link = e.ChildAttr("a", "href")
				} else {
					link = prefix + e.ChildAttr("a", "href")
				}
			}

			e.ForEachWithBreak("span", func(i int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "a-size-base-plus a-color-base a-text-normal" {
					title := e.Text
					name = title
					return false
				}
				return true
			})

			e.ForEachWithBreak("div", func(i int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "a-row" {
					price = e.Text
					return false
				}
				return true
			})

			if name != "" && price != "" && link != "" {
				product := models.Product{
					Name:  name,
					Price: price,
					Link:  link,
				}

				fmt.Println("produto:")
				fmt.Println(product.Name)
				fmt.Println(product.Price)
				fmt.Println(product.Link)
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
