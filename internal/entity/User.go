package entity

import "github.com/ribeirosaimon/Mockly/internal/auth"

type User struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     auth.Role `json:"role"`
}
