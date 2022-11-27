package model

type User struct {
	Id       int64  `json:"-" db:"id"`
	FullName string `json:"full_name" db:"full_name" binding:"required"`
	Balance  int64  `json:"balance" db:"balance" binding:"required"`
}
