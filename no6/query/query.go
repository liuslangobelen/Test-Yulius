package query

type JoinShoppingCart struct {
	ID          uint   `gorm:"column:id" json:"id"`
	UserId      uint   `gorm:"column:user_id" json:"user_id"`
	ProductId   uint   `gorm:"column:product_id;type:bigint(20)"`
	Quantity    int    `gorm:"column:quantity"`
	CodeProduct string `gorm:"column:code_product"`
	NamaProduct string `gorm:"column:nama_product"`
}
