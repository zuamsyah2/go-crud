package model

type Data struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
	Umur string `json:"umur"`
}

type DataBodyRequest struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Umur string `json:"umur"`
}
