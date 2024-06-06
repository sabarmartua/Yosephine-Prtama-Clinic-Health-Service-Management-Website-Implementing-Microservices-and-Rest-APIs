package dto

type NewUlasanDTO struct {
	ID        uint64 `json:"id"`
	UserID    uint   `json:"userID" binding:"required"`
	IsiUlasan string `json:"isiUlasan" binding:"required"`
}