package services

import (
	"fmt"

	"github.com/gabrieldebem/go-crawler/packages/models"
	"github.com/gocolly/colly/v2"
)

type MercadoLivreCrawler struct {
	products []models.Product
}

func (m MercadoLivreCrawler) Search(url string) []models.Product {
	c := colly.NewCollector()

	var name string
	var price string
	var link string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "ui-search-result__content ui-search-link" {
			link = e.Attr("href")
			name = e.Attr("title")

			e.ForEachWithBreak("span", func(i int, e *colly.HTMLElement) bool {
				if e.Attr("class") == "price-tag-fraction" {
					price = e.Text
					return false
				}
				return true
			})

			product := models.Product{
				Name:  name,
				Price: "R$ " + price,
				Link:  link,
			}

			fmt.Println("Produto:")
			fmt.Println(product.Name)
			fmt.Println(product.Price)
			fmt.Println(product.Link)

			m.products = append(m.products, product)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.Visit(url)

	return m.products
}
