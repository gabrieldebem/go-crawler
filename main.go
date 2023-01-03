package main

import (
	"github.com/gabrieldebem/go-crawler/packages/services"
	"github.com/gabrieldebem/go-crawler/packages/contracts"
)

var crawler contracts.ICrawler

func main() {
	url := "https://www.amazon.com.br/s?k=notebook+acer+nitro+5&sprefix=notebook+acer+%2Caps%2C272&ref=nb_sb_ss_ts-doa-p_2_14"
	crawler = services.AmazonCrawler{}

	crawler.Search(url)
}
