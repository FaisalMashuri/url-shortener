package response

type UrlResponse struct {
	Id          uint   `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
	View        int16  `json:"view"`
}
