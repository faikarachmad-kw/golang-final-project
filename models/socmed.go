package models

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" form:"social_media_url" json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	User           User   `json:"user" gorm:"constraint:OnDelete:CASCADE;"`
}
