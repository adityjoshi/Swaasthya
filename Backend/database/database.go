package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=aditya dbname=hospital port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Print("database connected successfully ⚡️")
	}

	// Migrate the schema
	DB.AutoMigrate()
}

var DB *gorm.DB

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type UserType string

const (
	Staff   UserType = "Staff"
	Patient UserType = "Patient"
)

type Users struct {
	user_id       uint   `gorm:"primaryKey"`
	full_Name     string `gorm:"not null"`
	genderInfo    Gender
	contactNumber uint   `gorm:"not null"`
	email         string `gorm:"not null;unique"`
	password      string
	user_type     UserType
}
