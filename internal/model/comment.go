package model

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `json:"content"`
	UserID  uint   `json:"-"`
}

func (Comment) TableName() string {
	return "comment"
}
