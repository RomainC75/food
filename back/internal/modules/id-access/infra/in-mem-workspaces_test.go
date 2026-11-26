package idaccess_infra

import (
	"context"
	idaccess_dom "food/internal/modules/id-access/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaces(t *testing.T) {
	userId := uuid.MustParse("37e88ec7-ef9f-4c36-82cf-d8485a273da7")
	user := idaccess_dom.NewUser(userId, "fake@email.com", "azerty")
	workspaceId := uuid.MustParse("afd0de35-95d8-4ebc-bbed-e0f2be4291b4")
	workspaceName := "my-workspace"
	w := idaccess_dom.NewWorkspace(workspaceId, workspaceName)

	membershipId := uuid.MustParse("27f11144-d080-4ff6-a12b-810560c2b035")
	m := idaccess_dom.NewMembership(membershipId, user.Id(), w.Id(), idaccess_dom.Manager)

	t.Run("should create workspaces", func(t *testing.T) {
		workspaces := NewInMemWorkspaces()
		ctx := context.Background()
		err := workspaces.CreateWithMembership(ctx, w, m)
		assert.Nil(t, err)

		foundWorkspaces, err := workspaces.GetWorkspaces(ctx, user.Id())
		assert.Nil(t, err)
		assert.Len(t, foundWorkspaces, 1)
		assert.Equal(t, w, foundWorkspaces[0])
	})

}
