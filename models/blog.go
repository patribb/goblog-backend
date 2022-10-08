package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"content"`
	Image  string `json:"image"`
	UserID string `json:"userID"`
	User   User   `json:"user";gorm:"foreingkey:UserID"`
}
