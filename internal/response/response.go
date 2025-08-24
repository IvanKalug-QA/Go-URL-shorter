package response

type ShortUrl struct {
	Status int    `json:"status"`
	Url    string `json:"url"`
}

type OriginalUrl struct {
	Status int    `json:"status"`
	Url    string `json:"url"`
}
