package model

import (
	"gorm.io/gorm"
)

type AccountMaster struct {
	gorm.Model
	AccountNo  uint          `gorm:"not null;size:32" json:"accountNo,omitempty"`
	Balance    uint          `gorm:"not null;size:32" json:"balance,omitempty"`
	Hold       uint          `gorm:"not null;size:32" json:"hold,omitempty"`
	Type       AccountType   `gorm:"not null;size:32" json:"accountType,omitempty"`
	Health     AccountHealth `gorm:"not null;size:32" json:"accountHealth,omitempty"`
	AccOwnerID uint
	AccOwner   Owner
}

type AccountHealth string
type AccountType string

const (
	OK         AccountHealth = "ok"
	WARN       AccountHealth = "Warning"
	SUB1       AccountHealth = "Substandard 1"
	SUB2       AccountHealth = "Substandard 2"
	SUB3       AccountHealth = "Substandard 3"
	OVERDUE30  AccountHealth = "Overdue by 30 days"
	OVERDUE90  AccountHealth = "Overdue by 90 days"
	DELINQUENT AccountHealth = "Delinquent"
)

const (
	SAVING    AccountType = "Savings Account"
	CURRENT   AccountType = "Current Account"
	GL        AccountType = "GL Account"
	OVERDRAFT AccountType = "Overdraft Account"
)
