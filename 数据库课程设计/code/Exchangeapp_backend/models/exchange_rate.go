package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primarykey" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required" gorm:"index:idx_currency_pair"`
	ToCurrency   string    `json:"toCurrency" binding:"required" gorm:"index:idx_currency_pair"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
