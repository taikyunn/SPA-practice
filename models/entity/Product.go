package entity

type Product struct {
	ID          int    `gorm:"primary_key;not null"       json:"id"`
	ProductName string `gorm:"type:varchar(200);not null" json:"name"`
	Memo        string `gorm:"type:varchar(400)"          json:"memo"`
	Status      int    `gorm:"not null"                   json:"state"`
}
