package dto

type UpdateKategoriArtikelDTO struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string    `gorm:"type:varchar(255)" json:"nama"`
}
