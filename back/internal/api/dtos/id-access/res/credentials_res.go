package dto_idaccess_res

import "github.com/google/uuid"

type SigninResDto struct {
	UserId uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Token  string    `json:"token"`
}
