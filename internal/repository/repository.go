package repository

import urlDomain "github.com/arymaulanamalik/shortener/internal/domain/url"

// insert, update, delete, find
type Repository interface {
	InsertURL(req *urlDomain.URL) error
	FindByFilter(filter urlDomain.URLFilter) (*[]urlDomain.URL, error)
}
