package dto

import "time"

type UpdateObatDTO struct {
	ID          uint64    `json:"id"`
	Nama        string    `json:"nama" binding:"required"`
	ExpiredDate time.Time `json:"expiredDate" binding:"required"`
	JumlahStok  uint64    `json:"jumlahStok" binding:"required"`
	Deskripsi   string    `json:"deskripsi" binding:"required"`
}
