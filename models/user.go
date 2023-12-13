package models

type Role string

const (
	ADMIN Role = "Admin"
	USER  Role = "User"
)

type User struct {
	BaseModel
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     Role   `json:"-"`
}
