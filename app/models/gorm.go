package models

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() {
	var err error
	var dsn = os.Getenv("REVELAPP_DBUSER") +
		":" + os.Getenv("REVELAPP_DBPASSWD") +
		"@" + os.Getenv("REVELAPP_DBHOSTNAME") +
		"/" + os.Getenv("REVELAPP_DBNAME") +
		"?parseTime=true&loc=Asia%2FTokyo"

	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.LogMode(true)
	autoMigrate()

	log.Println("Connected to database.")
}

func autoMigrate() {
	// Db.AutoMigrate(&User{})
	// Db.AutoMigrate(&Project{})
	// Db.AutoMigrate(&Item{})
	// Db.AutoMigrate(&DateTran{})
}

type Model struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
}
