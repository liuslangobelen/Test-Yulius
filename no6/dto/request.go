package dto

type ReqAddCart struct {
	CodeProduct string `json:"codeProduct" binding:"required"`
	NameProduct string `json:"nameProduct" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
