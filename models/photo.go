package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Photo's title are required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo's title are required"`
	UserID   uint   `json:"user_id" form:"user_id"`
	User     User   `json:"user" gorm:"constraint:OnDelete:CASCADE;"`
	Comments []Comment `json:"Comments" gorm:"constraint:OnDelete:CASCADE;"`
}
