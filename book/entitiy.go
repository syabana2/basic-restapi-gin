package book

import "time"

type Book struct {
	ID          int
	Title       string `gorm:"size:100;unique;not null"`
	Description string `gorm:"size:300"`
	Price       int    `gorm:"size:10"`
	Rating      int    `gorm:"size:10"`
	Discount    int    `gorm:"size:10"'`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
