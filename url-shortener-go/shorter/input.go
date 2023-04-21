package shorter

type UrlInput struct {
	Title          string `json:"title" binding:"required"`
	DestinationUrl string `json:"destination_url" binding:"required"`
	Tags           string `json:"tags"`
}

type ShorterUrlInput struct {
	Title          string `json:"title" binding:"required"`
	DestinationUrl string `json:"destination_url" binding:"required"`
	Tags           string `json:"tags"`
	BackHalf       string `json:"back_half" binding:"required"`
}
