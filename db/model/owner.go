package model

import "gorm.io/gorm"

type Owner struct {
	gorm.Model
	Name          string `gorm:"not null;size:255" json:"name,omitempty"`
	CRNumber      uint   `gorm:"not null;size:32" json:"crNo,omitempty"`
	ContactableID uint
	Contactable   Contactables
}
type Contact struct {
	gorm.Model
	Email    string `gorm:"not null;size:255" json:"email,omitempty"`
	PhoneNo  uint   `gorm:"not null;size:255" json:"phoneNo,omitempty"`
	Location string `gorm:"not null;size:255" json:"location,omitempty"`
	Addr1    string `gorm:"not null;column:Address_Line_1;size:255" json:"address_1,omitempty"`
	Addr2    string `gorm:"column:Address_Line_2;size:255" json:"address_2,omitempty"`
	Addr3    string `gorm:"column:Address_Line_3t;size:255" json:"address_3,omitempty"`
}

type Contactables struct {
	gorm.Model
	ContactID uint
	Contact   Contact
	IsActive  bool `gorm:"not null" json:"isActive,omitempty"`
}
