package contracts

import "github.com/gabrieldebem/go-crawler/packages/models"

type ICrawler interface {
	Search(url string) []models.Product
}

