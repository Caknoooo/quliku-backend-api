package entities

type ListBank struct {
	ID   int    `gorm:"type:int;primary_key;auto_increment" json:"id"` // Primary key
	Nama string `json:"nama"`
}