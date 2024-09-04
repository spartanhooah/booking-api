package models

type User struct {
	ID       int64  `json:"id"` //binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `binding:"required"`
}
