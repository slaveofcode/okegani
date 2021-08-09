package models

import "gorm.io/datatypes"

const AccountTableName = "Accounts"

type Account struct {
	BaseModel
	Username       string         `gorm:"column:username;unique" json:"username"`
	AdditionalInfo datatypes.JSON `gorm:"column:additionalInfo" json:"additionalInfo"`
	Status         string         `gorm:"column:status;index" json:"status"`
	AgreementMeta  datatypes.JSON `gorm:"column:agreementMeta" json:"agreementMeta"`
	Posts          []Post         `gorm:"foreignKey:AccountId"`
	Discussions    []Discussion   `gorm:"foreignKey:AccountId"`
	PostVotes      []PostVote     `gorm:"foreignKey:AccountId"`
}

func (m Account) TableName() string {
	return AccountTableName
}
