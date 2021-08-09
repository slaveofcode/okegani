package models

import "time"

const AccountCredentialTableName = "AccountCredentials"

type AccountCredential struct {
	BaseModel
	AccountId   uint       `gorm:"column:accountId;index" json:"accountId"`
	Account     Account    `gorm:"foreignKey:AccountId" json:"account"`
	CredType    string     `gorm:"column:credType;index" json:"credType"`
	CredValue   string     `gorm:"column:credValue;index" json:"credValue"`
	Creds       string     `gorm:"creds" json:"-"`
	IsActive    bool       `gorm:"isActive;index" json:"isActive"`
	ActivatedAt *time.Time `gorm:"activatedAt;index" json:"activatedAt"`
	ReplacedAt  *time.Time `gorm:"replacedAt;index" json:"replacedAt"`
}

func (m AccountCredential) TableName() string {
	return AccountCredentialTableName
}
