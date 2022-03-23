package models

type Comment struct {
	GormModel
	UserID  uint   `json:"UserID"`
	User    User   `json:"user"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Photo   Photo  `json:"photo"`
	Message string `gorm:"not null" json:"message" form:"message"`
}
