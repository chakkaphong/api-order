package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderStatusEnum string

const (
	OrderStatusEnumCreated   OrderStatusEnum = "CREATED"
	OrderStatusEnumRejected  OrderStatusEnum = "REJECTED"
	OrderStatusEnumCompleted OrderStatusEnum = "COMPLETED"
)

type Orders struct {
	Id           int             `gorm:"primaryKey;autoIncrement"`
	CustomerName string          `gorm:"type:varchar(100);"`
	TotalAmount  decimal.Decimal `gorm:"type:decimal(10,2); not null"`
	Status       OrderStatusEnum `gorm:"type:varchar(20);"`
	CreatedAt    time.Time       `gorm:"autoCreateTime;not null"`
	UKpdatedAt   time.Time       `gorm:"autoCreateTime;not null"`

	// FK
	OrderItems []OrderItems `gorm:"foreignKey:OrderId;references:Id"`
}
