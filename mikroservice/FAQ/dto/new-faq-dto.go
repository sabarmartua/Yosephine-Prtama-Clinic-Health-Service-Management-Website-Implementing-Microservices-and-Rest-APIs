package dto

type NewFAQDTO struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Pertanyaan string `gorm:"type:varchar(255)" json:"pertanyaan"`
	Jawaban    string `gorm:"type:varchar(255)" json:"jawaban"`
}
