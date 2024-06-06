package dto

type UpdateUlasanDTO struct {
	ID        uint64 `json:"id"`
	UserID    uint   `json:"userID" binding:"required"`
	IsiUlasan string `json:"isiUlasan" binding:"required"`
}