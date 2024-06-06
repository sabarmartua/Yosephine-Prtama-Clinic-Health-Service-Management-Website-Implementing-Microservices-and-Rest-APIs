package model

type Artikel struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Nama       string `gorm:"type:varchar(255)" json:"nama"`
	Konten     string `gorm:"type:text" json:"konten"`
	KategoriID uint64 `gorm:"column:kategori_id" json:"kategori_id" validate:"required"`
	Gambar     string `gorm:"type:varchar(255)" json:"gambar" validate:"required,oneof= jpg jpeg png"`
}
