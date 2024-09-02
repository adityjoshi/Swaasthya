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
	DB.AutoMigrate(&Users{}, &PatientInfo{})
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
	User_id       uint   `gorm:"primaryKey"`
	Full_Name     string `gorm:"not null"`
	GenderInfo    Gender
	ContactNumber uint   `gorm:"not null"`
	Email         string `gorm:"not null;unique"`
	Password      string
	User_type     UserType
}

type PatientInfo struct {
	Username  string `gorm:"not null"`
	City      string
	State     string
	PinCode   uint
	Adhar     string
	PatientID uint `gorm:"primaryKey;foreignKey:PatientID;references:Users(User_id);onDelete:CASCADE"`
}

func CloseDatabase() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}
}
