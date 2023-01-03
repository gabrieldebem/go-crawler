package handlers

import (
	"fmt"
	"sort"

	"github.com/gabrieldebem/go-crawler/packages/contracts"
	"github.com/gabrieldebem/go-crawler/packages/services"
	"github.com/gin-gonic/gin"
)

var crawler contracts.ICrawler

func GetProducts(c *gin.Context) {
	crawler = services.MercadoLivreCrawler{}
	p := c.Query("product")
	url := fmt.Sprintf("https://lista.mercadolivre.com.br/%s#D[A:%s]", p, p)
	mProducts := crawler.Search(url)

	sort.Slice(mProducts, func(i, j int) bool {
		return mProducts[i].Price < mProducts[j].Price
	})

	crawler = services.AmazonCrawler{}
	url = fmt.Sprintf("https://www.amazon.com.br/s?k=%s", p)
	aProducts := crawler.Search(url)

	sort.Slice(aProducts, func(i, j int) bool {
		return aProducts[i].Price < aProducts[j].Price
	})

	c.JSON(200, gin.H{
		"mercado_livre": mProducts,
		"amazon":        aProducts,
	})
}
