package model

type Property struct {
	ProductId uint   `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
