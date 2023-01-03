package handlers

import (
	"github.com/gabrieldebem/go-crawler/packages/contracts"
	"github.com/gabrieldebem/go-crawler/packages/services"
	"github.com/gin-gonic/gin"
)

var crawler contracts.ICrawler

func GetProducts(c *gin.Context) {
	crawler = services.MercadoLivreCrawler{}
	url := "https://lista.mercadolivre.com.br/notebook#D[A:notebook]"
	mProducts := crawler.Search(url)

	crawler = services.AmazonCrawler{}
	url = "https://www.amazon.com.br/s?k=notebook&__mk_pt_BR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=3QPMLAG2FS1F9&sprefix=notebook%2Caps%2C236&ref=nb_sb_noss_1"
	aProducts := crawler.Search(url)

	c.JSON(200, gin.H{
		"mercado_livre": mProducts,
		"amazon":        aProducts,
	})
}
