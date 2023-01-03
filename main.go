package main

import (
	"github.com/gabrieldebem/go-crawler/packages/services"
	"github.com/gabrieldebem/go-crawler/packages/contracts"
)

var crawler contracts.ICrawler

func main() {
	url := "https://lista.mercadolivre.com.br/notebook#D[A:notebook]"
	crawler = services.MercadoLivreCrawler{}

	crawler.Search(url)
}
