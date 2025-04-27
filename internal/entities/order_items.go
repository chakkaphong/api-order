package entities

import "github.com/shopspring/decimal"

type OrderItems struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	OrderId     int    `gorm:"not null"`
	ProductName string `gorm:"type:varchar(100);"`
	Quantity    int
	Price       decimal.Decimal `gorm:"type:decimal(10,2); not null"`
}
