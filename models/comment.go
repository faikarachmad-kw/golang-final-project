package models

type Comment struct {
	GormModel
	UserID  uint   `json:"UserID"`
	User    User   `json:"user" gorm:"constraint:onDelete:SET NULL;"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Photo   Photo  `gorm:"constraint:onDelete:SET NULL;" json:"photo"`
	Message string `gorm:"not null" json:"message" form:"message"`
}
