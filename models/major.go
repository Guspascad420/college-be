package models

type Major struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"nama"`
}
