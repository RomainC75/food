package dto_idaccess_res

import "github.com/google/uuid"

type SignupResDto struct {
	UserId uuid.UUID `json:"id"`
	Email  string    `json:"email"`
}

type SigninResDto struct {
	UserId uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	Token  string    `json:"token"`
}
