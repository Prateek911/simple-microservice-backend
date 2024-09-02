package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Firstname  string       `gorm:"not null;size:255" json:"firstName,omitempty"`
	LastName   string       `gorm:"not null;size:255" json:"LastName,omitempty"`
	MiddleName string       `gorm:"size:255" json:"MiddleName,omitempty"`
	Mnemonic   string       `gorm:"not null;size:32" json:"mnemonic,omitempty"`
	Department WorkGroup    `gorm:"not null;size:32" json:"department,omitempty"`
	Role       EmployeeRole `gorm:"not null;size:32" json:"role,omitempty"`
	ContactID  uint
	Contact    Contact
}

type EmployeeRole string
type WorkGroup string

const (
	INFRASTRUCTURE_FINANCE WorkGroup = "Infrastructure Finance"
	GOVERNMENT_FINANCE     WorkGroup = "Government Project Finance"
	RETAIL_FINANCE         WorkGroup = "Retail and Investment Finance"
	PERSONAL_FINANCE       WorkGroup = "Personal Finance"
	APP_SUPPORT            WorkGroup = "Application Support"
	COMPLIANCE_GROUP       WorkGroup = "Compliance Group"
	ASSET_FINANCE          WorkGroup = "Asset Finance"
)

const (
	TRANSACTION_MANAGER EmployeeRole = "Transaction Manager"
	TEAM_LEAD           EmployeeRole = "Team Lead"
	DEAL_MANAGER        EmployeeRole = "Deal Manager"
	CREDIT_ANALYST      EmployeeRole = "Credit Analyst"
	TRANSACTOR          EmployeeRole = "Transactor"
	SYS_USER            EmployeeRole = "System User"
	DEVELOPER           EmployeeRole = "Developer"
	ADMIN               EmployeeRole = "Admin User"
	QUALITY_ANALYST     EmployeeRole = "Quality Analyst"
	AUTOMATION          EmployeeRole = "Automation Tool"
)
