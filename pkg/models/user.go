package models

type Users []User

type User struct {
	IdU         uint   `json:"idu" gorm:"primaryKey"`
	Email       string `json:"email" gorm:"not null;unique;index"`
	Password    string `json:"password" gorm:"not null;"`
	ImgAvatar   string `json:"img_avatar"`
	Role        string `json:"role"`
	IsConfirmed bool   `json:"is_confirmed"`
}
