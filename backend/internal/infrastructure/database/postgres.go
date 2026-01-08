package database

import (
	"fmt"
	"os"

	"MediLink/internal/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entity.User{},
		entity.RefreshToken{},
		entity.Clinic{},
		entity.Doctor{},
		entity.DoctorClinicPlacement{},
		entity.Patient{},
		entity.Medicine{},
		entity.ClinicInventory{},
		entity.DoctorSchedule{},
		entity.Appointment{},
		entity.MedicalRecord{},
		entity.Prescription{},
		entity.PrescriptionItem{},
		entity.Billing{},
		entity.Payment{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func ClosDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
