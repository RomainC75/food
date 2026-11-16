package idaccess_dom

import "github.com/google/uuid"

type Membership struct {
	id          uuid.UUID
	userId      uuid.UUID
	workspaceId uuid.UUID
	role        Role
}

func NewMembership(id uuid.UUID, userId uuid.UUID, workspaceId uuid.UUID, role Role) *Membership {
	return &Membership{
		id:          id,
		userId:      userId,
		workspaceId: workspaceId,
		role:        role,
	}
}

func (m *Membership) Id() uuid.UUID {
	return m.id
}

func (m *Membership) WorkspaceId() uuid.UUID {
	return m.workspaceId
}

func (m *Membership) IsSameWorkspace(workspaceId uuid.UUID) bool {
	return m.workspaceId == workspaceId
}

func (m *Membership) IsSameUser(userId uuid.UUID) bool {
	return m.userId == userId
}

func (m *Membership) Clone() *Membership {
	return &Membership{
		id:          m.id,
		userId:      m.userId,
		workspaceId: m.workspaceId,
		role:        m.role,
	}
}
