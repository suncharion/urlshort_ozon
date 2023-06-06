package storage

import "ozon/internal/model"

type Storage interface {
	Put(*model.ShortenedUrl) error
	Get(string) (*model.ShortenedUrl, error)
}
