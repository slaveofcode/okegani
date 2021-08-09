package models

const PostVoteTableName = "PostVotes"

type PostVote struct {
	BaseModel
	AccountId uint    `gorm:"column:accountId;index" json:"accountId"`
	Account   Account `gorm:"foreignKey:AccountId" json:"account"`
	PostId    uint    `gorm:"column:postId;index" json:"postId"`
	Post      Post    `gorm:"foreignKey:PostId" json:"post"`
	Type      string  `gorm:"column:type" json:"type"`
}

func (m PostVote) TableName() string {
	return PostVoteTableName
}
