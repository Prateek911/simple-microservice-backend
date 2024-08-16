package model

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	BeneAcc         uint   `gorm:"column:Beneficiary_account;not null;size:255" json:"beneficiary_account,omitempty"`
	PayeeAcc        uint   `gorm:"column:Payee_account;not null;size:255" json:"payee_account,omitempty"`
	PaymentResponse string `gorm:"column:Payment_response_string;not null;size:1024" json:"payment_response,omitempty"`
	Status          bool   `gorm:"not null;size:255" json:"status,omitempty"`
}
