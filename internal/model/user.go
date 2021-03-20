package model

type User struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Comments []Comment `json:"comments"`
}
