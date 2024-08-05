package model

type Product struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	Price      string     `json:"price"`
	Currency   string     `json:"currency"`
	ApiCode    string     `json:"apiCode"`
	Preview    string     `json:"preview"`
	Properties []Property `json:"properties"`
}
