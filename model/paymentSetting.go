package model

import "time"

// PaymentSetting ...
type PaymentSetting struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"name"`
	Value     string    `json:"value" gorm:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName indicates table name of user
func (PaymentSetting) TableName() string {
	return "payment_settings"
}
