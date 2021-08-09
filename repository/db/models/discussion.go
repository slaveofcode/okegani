package models

import (
	"gorm.io/datatypes"
)

const DiscussionTableName = "Discussions"

type Discussion struct {
	BaseModel
	AccountId      uint           `gorm:"column:accountId;index" json:"accountId"`
	Account        Account        `gorm:"foreignKey:AccountId" json:"account"`
	PostId         uint           `gorm:"column:postId;index" json:"postId"`
	Post           Post           `gorm:"foreignKey:PostId" json:"post"`
	Content        string         `gorm:"column:content" json:"content"`
	AdditionalInfo datatypes.JSON `gorm:"column:additionalInfo" json:"additionalInfo"`
	Metadata       datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	Status         string         `gorm:"column:status;index" json:"status"`
}

func (m Discussion) TableName() string {
	return DiscussionTableName
}
