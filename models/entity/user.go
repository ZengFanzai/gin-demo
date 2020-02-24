package entity

import (
	"src/global"
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-" gorm:"not null;default:NOW()"`
	UpdatedAt time.Time  `json:"-" gorm:"not null;default:NOW()"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Username  string     `json:"username" gorm:"not null;DEFAULT:''"`
	Password  string     `json:"-" gorm:"not null;DEFAULT:''"`
	Mobile    string     `json:"mobile" gorm:"not null;DEFAULT:''"`
	Email     string     `json:"email" gorm:"not null;DEFAULT:''"`
}

func init() {
	global.DB.Set("gorm:table_options", "AUTO_INCREMENT = 10000").AutoMigrate(User{})
}
