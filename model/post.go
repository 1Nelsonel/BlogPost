package model

import (
    "time"
    // "github.com/jinzhu/gorm"
)


// User struct
type Author struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	Name string    `json:"name" gorm:"not null;column:name;size:255"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index;not null;column:email"`
    CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}

type Category struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	Name string    `json:"name" gorm:"not null;column:name;size:255"`
    CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}


type Blog struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Title       string `json:"title" gorm:"not null;column:title;size:255"`
    AuthorID    uint   `json:"author_id" gorm:"not null"`
    Author      Author `gorm:"foreignkey:AuthorID"`
    Categories  []Category `gorm:"many2many:blog_categories;"`
    Content     string `json:"content" gorm:"type:text;not null;column:content"`
    PublishedAt time.Time
}
