package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"nama"`
	NIM      string   `json:"nim" gorm:"unique"`
	MajorID  int      `json:"prodiId"`
	Angkatan int      `json:"angkatan"`
	Password string   `json:"password"`
	Major    Major    `gorm:"foreignKey:MajorID" json:"prodi"`
	Courses  []Course `gorm:"many2many:user_courses" json:"matakuliah"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
