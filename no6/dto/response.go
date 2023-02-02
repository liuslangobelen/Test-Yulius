package dto

type ResShoppingCart struct {
	ID          uint   `json:"id"`
	Quantity    int    `json:"quantity"`
	CodeProduct string `json:"codeProduct"`
	NameProduct string `json:"nameProduct"`
}
