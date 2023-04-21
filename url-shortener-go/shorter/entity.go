package shorter

import "time"

type ShortUrl struct {
	ID             int
	Title          string
	Tags           string
	DestinationUrl string
	BackHalf       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
