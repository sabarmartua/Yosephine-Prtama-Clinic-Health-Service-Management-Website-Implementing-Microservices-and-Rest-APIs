package model

import "time"

type CutiDokter struct {
	ID             uint64    `gorm:"primary_key:auto_increment" json:"id"`
	TanggalMulai   time.Time `gorm:"type:date" json:"tanggalMulai"`
	TanggalSelesai time.Time `gorm:"type:date" json:"tanggalSelesai"`
	Keterangan     string    `gorm:"type:varchar(255)" json:"keterangan"`
}
