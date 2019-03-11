package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User user information
type User struct {
	gorm.Model
	Name     string
	Mail     string
	Password []byte
	Projects []Project
}

// Project Project information
type Project struct {
	gorm.Model
	Name        string
	Description string
	UserID      int
	User        User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Items       []Item
}

type Item struct {
	gorm.Model
	Name      string
	Weight    int8
	ProjectID int
	Project   Project `gorm:"ForeignKey:ProjectID;AssociationForeignKey:ID"`
}

type DateTran struct {
	gorm.Model
	Date   time.Time
	score  int8
	ItemID int
	Item   Item `gorm:"ForeignKey:ProjectID;AssociationForeignKey:ID"`
}
