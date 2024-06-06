package dto

import "time"

type NewCutiDokterDTO struct {
	ID             uint64    `json:"id"`
	TanggalMulai   time.Time `json:"tanggalMulai" binding:"required"`
	TanggalSelesai time.Time `json:"tanggalSelesai" binding:"required"`
	Keterangan     string    `json:"keterangan" binding:"required"`
}