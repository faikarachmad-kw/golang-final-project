package models

type Comment struct {
	GormModel
	UserID  uint   `json:"UserID" gorm:"onDelete:SET NULL;"`
	User    User   `json:"user" gorm:"onDelete:SET NULL;"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Photo   Photo  `gorm:"onDelete:CASCADE" json:"photo"`
	Message string `gorm:"not null" json:"message" form:"message"`
}
