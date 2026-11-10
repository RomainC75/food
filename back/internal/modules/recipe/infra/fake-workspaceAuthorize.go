package recipe_infra

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type FakeWorkspaceAuthorizer struct {
	expectedErrorMsg string
}

func NewFakeWorkspaceAuthorizer() *FakeWorkspaceAuthorizer {
	return &FakeWorkspaceAuthorizer{}
}

func (wa *FakeWorkspaceAuthorizer) CanCreateInWorkspace(
	ctx context.Context,
	userID uuid.UUID,
	workspaceID uuid.UUID,
) error {
	return errors.New(wa.expectedErrorMsg)
}

func (wa *FakeWorkspaceAuthorizer) SetExpectedError(errorMsg string) *FakeWorkspaceAuthorizer {
	wa.expectedErrorMsg = wa.expectedErrorMsg
	return wa
}
