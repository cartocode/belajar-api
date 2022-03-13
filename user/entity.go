package user

import "time"

type User struct {
	ID             int 
	Name           string
	Occupation     string
	Email          string
	PasswordHas    string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt 		 time.Time
}