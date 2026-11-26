package dto_idaccess_res

import (
	idaccess_dom "food/internal/modules/id-access/domain"

	"github.com/google/uuid"
)

type WorkspaceDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type MembershipCreated struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ToWorkspaceDto(workspaces []*idaccess_dom.Workspace) []WorkspaceDTO {
	wsDto := make([]WorkspaceDTO, 0, len(workspaces))
	for _, w := range workspaces {
		wsDto = append(wsDto, WorkspaceDTO{
			Id:   w.Id(),
			Name: w.Name(),
		})
	}
	return wsDto
}
