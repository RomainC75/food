package dto_idaccess_res

import "github.com/google/uuid"

type MembershipCreated struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
