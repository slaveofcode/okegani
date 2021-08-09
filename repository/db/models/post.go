package models

import (
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

const PostTableName = "Posts"

type Post struct {
	BaseModel
	AccountId      uint            `gorm:"column:accountId;index" json:"accountId"`
	Account        Account         `gorm:"foreignKey:AccountId" json:"account"`
	Title          string          `gorm:"column:title;index" json:"title"`
	Summary        string          `gorm:"column:summary" json:"summary"`
	Content        datatypes.JSON  `gorm:"column:content;index" json:"content"`
	CategoryId     uint            `gorm:"column:categoryId;index" json:"categoryId"`
	Category       Category        `gorm:"foreignKey:CategoryId"`
	Tags           pq.StringArray  `gorm:"type:text[];column:tags;index" json:"tags"`
	VoteSummary    datatypes.JSON  `gorm:"column:voteSummary" json:"voteSummary"`
	Rank           decimal.Decimal `gorm:"column:rank;index" sql:"type:DECIMAL(20.8);" json:"rank"`
	AdditionalInfo datatypes.JSON  `gorm:"column:additionalInfo" json:"additionalInfo"`
	Metadata       datatypes.JSON  `gorm:"column:metadata" json:"metadata"`
	Status         string          `gorm:"column:status;index" json:"status"`
	Visibility     string          `gorm:"column:visibility;index" json:"visibility"`
	Discussions    []Discussion    `gorm:"foreignKey:PostId"`
}

func (m Post) TableName() string {
	return PostTableName
}
