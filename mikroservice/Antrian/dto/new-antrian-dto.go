package dto

import "time"

type NewAntrianDTO struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id" binding:"required"`
	Kepentingan  string    `json:"kepentingan" binding:"required"`
	Tanggal      time.Time `json:"tanggal" binding:"required"`
	Deskripsi    string    `json:"deskripsi" binding:"required"`
	NomorAntrian *uint64   `json:"nomorAntrian"`
}
