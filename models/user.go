package models



type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Role     string `gorm:"type:user_role"` // Use the custom enum type
}