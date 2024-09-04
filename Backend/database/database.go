package database

import (
	"fmt"
	"time"

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
	DB.AutoMigrate(&Users{}, &PatientInfo{}, &HospitalAdmin{}, &Hospitals{}, &Doctors{}, &Appointment{})
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
	Admin   UserType = "Admin"
)

type Users struct {
	User_id       uint     `gorm:"primaryKey"`
	Full_Name     string   `json:"Full_Name" gorm:"not null"`
	GenderInfo    Gender   `json:"GenderInfo"`
	ContactNumber string   `json:"ContactNumber" gorm:"not null"`
	Email         string   `json:"Email" gorm:"not null;unique"`
	Password      string   `json:"Password"`
	User_type     UserType `json:"User_type"`
}

type PatientInfo struct {
	Username  string `gorm:"not null"`
	City      string
	State     string
	PinCode   uint
	Adhar     string
	PatientID uint `gorm:"primaryKey;foreignKey:PatientID;references:Users(User_id);onDelete:CASCADE"`
}

type HospitalAdmin struct {
	AdminID       uint     `gorm:"primaryKey;autoIncrement"`
	FullName      string   `gorm:"not null" json:"full_name"`
	Email         string   `gorm:"unique;not null" json:"email"`
	Password      string   `gorm:"not null" json:"password"`
	ContactNumber string   `gorm:"not null" json:"contact_number"`
	Usertype      UserType `json:"user_type" gorm:"not null"`
}

type Hospitals struct {
	HospitalId    uint   `json:"hospital_id" gorm:"primaryKey;autoIncrement"`
	HospitalName  string `json:"hospital_name" gorm:"not null"`
	Address       string `json:"address" gorm:"not null"`
	City          string `json:"city" gorm:"not null"`
	State         string `json:"state" gorm:"not null"`
	PinCode       string `json:"pincode" gorm:"not null"`
	ContactNumber string `json:"contact_number" gorm:"not null"`
	Email         string `json:"email" gorm:"not null"`
	AdminID       uint   `json:"admin_id" gorm:"primaryKey;foreignKey:AdminID;references:HospitalAdmin(AdminID);onDelete:CASCADE"`
	Username      string `json:"username" gorm:"unique;not null"`
	Description   string `json:"description"`
}

type Department string

const (
	Cardiology  Department = "Cardiology"
	Neurology   Department = "Neurology"
	Orthopedics Department = "Orthopedics"
	Pediatrics  Department = "Pediatrics"
	Radiology   Department = "Radiology"
	Surgery     Department = "Surgery"
	InternalMed Department = "Internal Medicine"
)

type Doctors struct {
	DoctorID      uint       `json:"doctor_id" gorm:"primaryKey;autoIncrement"`
	FullName      string     `json:"full_name" gorm:"not null"`
	Description   string     `json:"description" gorm:"not null"`
	ContactNumber string     `json:"contact_number" gorm:"not null"`
	Email         string     `json:"email" gorm:"unique;not null"`
	HospitalID    uint       `json:"hospital_id" gorm:"not null;foreignKey:HospitalID;references:Hospitals(HospitalId)"`
	Hospital      string     `json:"hospital_name" gorm:"not null;foreignKey:Hospital;references:Hospitals(HospitalName);onDelete:CASCADE"`
	Department    Department `json:"department" gorm:"not null"`
	Username      string     `json:"username" gorm:"unique;not null"`
}

type Appointment struct {
	AppointmentID   uint      `json:"appointment_id" gorm:"primaryKey;autoIncrement"`
	PatientID       uint      `json:"patient_id" gorm:"not null;foreignKey:PatientID;references:PatientInfo(PatientID);onDelete:CASCADE"`
	DoctorID        uint      `json:"doctor_id" gorm:"not null;foreignKey:DoctorID;references:Doctors(DoctorID);onDelete:CASCADE"`
	AppointmentDate time.Time `json:"appointment_date" gorm:"not null"`
	AppointmentTime time.Time `json:"appointment_time" gorm:"not null"`
	Description     string    `json:"description"`
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
