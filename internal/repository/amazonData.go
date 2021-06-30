package repository

import (
	"time"
)

type AmazonData struct {
	Asin      string
	Maker     string
	Price     int64
	Name      string
	Reason    string
	Url       string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
