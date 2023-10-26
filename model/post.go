package model

import (
 "gorm.io/gorm"
)
// User struct
type Blog struct {
 gorm.Model
 ID       uint `json:"id" gorm:"primaryKey"`
 Title string    `json:"title"`
 Category    string    `json:"category"`
 Content string    `json:"content"`
}
