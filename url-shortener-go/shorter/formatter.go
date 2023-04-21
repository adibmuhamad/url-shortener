package shorter

type ShortUrlFormatter struct {
	Title          string `json:"title"`
	Tags           string `json:"tags"`
	DestinationUrl string `json:"detination_url"`
	BackHalf       string `json:"back_half"`
}

func FormatShortUrl(shortUrl ShortUrl) ShortUrlFormatter {
	formatter := ShortUrlFormatter{
		Title:          shortUrl.Title,
		Tags:           shortUrl.Tags,
		DestinationUrl: shortUrl.DestinationUrl,
		BackHalf:       shortUrl.BackHalf,
	}

	return formatter
}
