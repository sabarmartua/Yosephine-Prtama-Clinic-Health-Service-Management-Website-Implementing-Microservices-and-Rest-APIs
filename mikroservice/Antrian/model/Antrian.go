package model

import "time"

type Antrian struct {
	ID           uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID       uint64    `json:"user_id"`
	Kepentingan  string    `gorm:"type:varchar(255)" json:"kepentingan"`
	Tanggal      time.Time `gorm:"type:date" json:"tanggal"`
	Deskripsi    string    `gorm:"type:varchar(255)" json:"deskripsi"`
	NomorAntrian uint64    `json:"nomor_antrian"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
