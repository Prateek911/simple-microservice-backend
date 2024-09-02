package model

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Name      string `gorm:"not null;size:255"`
	CRNumber  uint   `gorm:"not null;size:32"`
	ContactID uint
	Contact   Contact
}
