package models

import "gorm.io/datatypes"

const CategoryTableName = "Categories"

type Category struct {
	BaseModel
	Name        string         `gorm:"column:name;index" json:"name"`
	Description string         `gorm:"column:description" json:"description"`
	Metadata    datatypes.JSON `gorm:"column:metadata" json:"metadata"`
	IsVisible   bool           `gorm:"column:isVisible;index" json:"isVisible"`
	Children    []*Category    `gorm:"many2many:SubCategories"`
}

func (m Category) TableName() string {
	return CategoryTableName
}
