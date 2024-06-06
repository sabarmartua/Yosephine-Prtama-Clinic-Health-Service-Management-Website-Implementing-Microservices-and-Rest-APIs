package model

import "time"

type Obat struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama        string    `gorm:"type:varchar(255)" json:"nama"`
	ExpiredDate time.Time `gorm:"type:date" json:"expiredDate"`
	JumlahStok  uint64    `gorm:"type:int" json:"jumlahStok"`
	Deskripsi   string    `gorm:"type:text" json:"deskripsi"`
}
