package idaccess_app

import (
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_dom "food/internal/modules/id-access/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateWorkspace(t *testing.T) {
	t.Run("should create a workspace", func(t *testing.T) {
		userId := uuid.MustParse("b574bde3-68ac-4291-94b3-7c7a774ac7b8")
		workspaceId := uuid.MustParse("54e8598c-900d-4888-b13e-44d9c70560f4")
		mmembershipId := uuid.MustParse("5ce691ed-a8b9-44ac-8e96-d56431c8d6a1")

		td := NewTestDriver(t).GivenUserIdContext(userId).GivenUuids(workspaceId, mmembershipId)
		workspaceName := "my-workspace"

		req := dto_idaccess_req.CreateWorkspace{
			Name: workspaceName,
		}
		ws, err := td.CreateWorkspace(req)
		assert.Nil(t, err)
		assert.NotNil(t, ws)

		expectedWorkspace := idaccess_dom.NewWorkspace(workspaceId, workspaceName)
		expectedMembership := idaccess_dom.NewMembership(mmembershipId, userId, workspaceId, idaccess_dom.Manager)
		assert.Equal(t, expectedWorkspace, ws)

		assert.Len(t, td.SavedWorkspaces(), 1)
		assert.Equal(t, expectedWorkspace, td.SavedWorkspaces()[0])
		assert.Len(t, td.SavedMemberships(), 1)
		assert.Equal(t, expectedMembership, td.SavedMemberships()[0])
	})
}
