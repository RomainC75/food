package dto_idaccess_req

type CreateWorkspace struct {
	Name string `json:"name" validate:"required,gte=8,lte=30"`
}
