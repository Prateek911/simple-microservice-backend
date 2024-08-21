package model

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Name          string `gorm:"not null;size:255"`
	CRNumber      uint   `gorm:"not null;size:32"`
	ContactableID uint
	Contactable   Contactables
}
type Contact struct {
	gorm.Model
	Email    string `gorm:"not null;size:255"`
	PhoneNo  uint   `gorm:"not null;size:255"`
	Location string `gorm:"not null;size:255"`
	Addr1    string `gorm:"not null;column:Address_Line_1;size:255"`
	Addr2    string `gorm:"column:Address_Line_2;size:255"`
	Addr3    string `gorm:"column:Address_Line_3t;size:255"`
}

type Contactables struct {
	gorm.Model
	ContactID uint
	Contact   Contact
	IsActive  bool `gorm:"not null" json:"isActive,omitempty"`
}
