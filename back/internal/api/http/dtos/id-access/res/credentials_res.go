package dto_idaccess_res

import "github.com/google/uuid"

type SignupResDto struct {
	UserId uuid.UUID `json:"id"`
	Email  string    `json:"email"`
}

type LoginResDto struct {
	UserId       uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RefreshTokenResDto struct {
	AccessToken string `json:"access_token"`
}
