package models

type Course struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"nama"`
}
