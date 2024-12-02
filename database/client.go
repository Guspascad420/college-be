package database

import (
	"college-be/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB
var err error

func Connect() {
	dbUrl := os.Getenv("DB_URL")
	dsn := dbUrl
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to DB")
	} else {
		log.Println("Connected to Database!")
	}
}

func Seed() {
	majors := []models.Major{
		models.Major{Name: "Teknologi Informasi"},
		models.Major{Name: "Pendidikan Teknologi Informasi"},
		models.Major{Name: "Sistem Informasi"},
		models.Major{Name: "Teknik Informatika"},
		models.Major{Name: "Teknik Komputer"},
	}
	Db.Create(majors)
}

func Migrate() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Major{})
	Db.AutoMigrate(&models.Course{})
	log.Println("Database Migration Completed!")
}
