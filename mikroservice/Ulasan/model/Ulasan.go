package model

type Ulasan struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID    uint   `json:"user_id"`
	IsiUlasan string `json:"isi_ulasan"`
}
