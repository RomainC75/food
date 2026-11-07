package idaccess_app_mod

import "github.com/google/uuid"

type SigninResult struct {
	UserID uuid.UUID
	Email  string
	Token  string
}
