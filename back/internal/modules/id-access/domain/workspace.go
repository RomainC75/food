package idaccess_dom

import "github.com/google/uuid"

type Workspace struct {
	id   uuid.UUID
	name string
}

func NewWorkspace(id uuid.UUID, name string) *Workspace {
	return &Workspace{
		id:   id,
		name: name,
	}
}

func (w *Workspace) Id() uuid.UUID {
	return w.id
}

func (w *Workspace) Name() string {
	return w.name
}

func (w *Workspace) Clone() *Workspace {
	return &Workspace{
		id:   w.id,
		name: w.name,
	}
}
