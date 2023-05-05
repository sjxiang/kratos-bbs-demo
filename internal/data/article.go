package data

import (
	"gorm.io/gorm"
)


type Article struct {
	gorm.Model

	Slug        string
	Title       string
	Description string
	Body        string
	Tags        []Tag  `gorm:"many2many:article_tags"`
}


type Tag struct {
	gorm.Model

	Name string
}

type Following struct {
	gorm.Model

	User      string
	Following uint
}
