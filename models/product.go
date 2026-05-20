package models

// Tempat deklarasi struct/skema tabel
type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Price int64  `json:"price"`
	Stock int    `json:"stock"`
}