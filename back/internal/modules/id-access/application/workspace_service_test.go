package idaccess_app

import (
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	userId1       = uuid.MustParse("b574bde3-68ac-4291-94b3-7c7a774ac7b8")
	workspaceId   = uuid.MustParse("54e8598c-900d-4888-b13e-44d9c70560f4")
	workspaceId2  = uuid.MustParse("80a3a165-837a-4c61-b3ba-591aca2c2b51")
	membershipId  = uuid.MustParse("5ce691ed-a8b9-44ac-8e96-d56431c8d6a1")
	membershipId2 = uuid.MustParse("890e71b6-8075-4b51-b8fc-7159d9f20746")
)

func TestUpdateWorkspace(t *testing.T) {
	t.Run("should not update workspace if not found", func(t *testing.T) {
		td := NewTestDriver(t).GivenUuids(workspaceId, membershipId)
		workspaceName := "my-workspace"

		updateCmd := UpdateWorkspaceCmd{
			WorkspaceId: workspaceId,
			Name:        workspaceName,
			UserId:      userId1,
		}
		_, err := td.UpdateWorkspace(updateCmd)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), idaccess_dom.ErrWorkspaceNotFound.Error())
		assert.Contains(t, err.Error(), updateCmd.WorkspaceId.String())
	})
	t.Run("should not update workspace if could not saved", func(t *testing.T) {
		td := NewTestDriver(t)
		td.createWorkspace(workspaceId, membershipId, userId1)
		td.SetupExpectedSaveWorkspaceError(true)
		workspaceName := "my-workspace2"

		updateCmd := UpdateWorkspaceCmd{
			WorkspaceId: workspaceId,
			Name:        workspaceName,
			UserId:      userId1,
		}
		_, err := td.UpdateWorkspace(updateCmd)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), idaccess_infra.ErrCouldNotSaveWorkspace.Error())
		assert.Contains(t, err.Error(), updateCmd.WorkspaceId.String())
	})
	t.Run("should update workspace", func(t *testing.T) {
		td := NewTestDriver(t)
		td.createWorkspace(workspaceId, membershipId, userId1)
		workspaceNewName := "my-workspace2"

		updateCmd := UpdateWorkspaceCmd{
			WorkspaceId: workspaceId,
			Name:        workspaceNewName,
			UserId:      userId1,
		}

		expectedWorkspace := idaccess_dom.NewWorkspace(workspaceId, workspaceNewName)
		updatedWorkspace, err := td.UpdateWorkspace(updateCmd)
		assert.Nil(t, err)
		assert.Equal(t, expectedWorkspace, updatedWorkspace)
	})
}

func TestCreateWorkspace(t *testing.T) {

	t.Run("should create a workspace", func(t *testing.T) {

		td := NewTestDriver(t).GivenUuids(workspaceId, membershipId)
		workspaceName := "my-workspace"

		req := CreateWorkspaceCmd{
			Name:   workspaceName,
			UserId: userId1,
		}
		ws, err := td.CreateWorkspace(req)
		assert.Nil(t, err)
		assert.NotNil(t, ws)

		expectedWorkspace := idaccess_dom.NewWorkspace(workspaceId, workspaceName)
		expectedMembership := idaccess_dom.NewMembership(membershipId, userId1, workspaceId, idaccess_dom.Manager)
		assert.Equal(t, expectedWorkspace, ws)

		assert.Len(t, td.SavedWorkspaces(), 1)
		assert.Equal(t, expectedWorkspace, td.SavedWorkspaces()[expectedWorkspace.Id()])
		assert.Len(t, td.SavedMemberships(), 1)
		assert.Equal(t, expectedMembership, td.SavedMemberships()[expectedMembership.Id()])
	})

	t.Run("should not create a workspace, cause repo error", func(t *testing.T) {
		td := NewTestDriver(t).GivenUuids(workspaceId, membershipId)
		td.SetupExpectedWorkspaceError(true)

		workspaceName := "my-workspace"

		req := CreateWorkspaceCmd{
			Name:   workspaceName,
			UserId: userId1,
		}
		_, err := td.CreateWorkspace(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrTryingToCreateWorkspace)
	})

	t.Run("should not create a duplicate workspace", func(t *testing.T) {
		td := NewTestDriver(t).GivenUuids(workspaceId, membershipId, workspaceId2, membershipId2)
		td.SetupExpectedWorkspaceError(false)

		workspaceName := "my-workspace"

		req := CreateWorkspaceCmd{
			Name:   workspaceName,
			UserId: userId1,
		}

		td.CreateWorkspace(req)

		_, err := td.CreateWorkspace(req)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), ErrWorkspaceAlreadyExists.Error())
		assert.Contains(t, err.Error(), workspaceName)
	})
}

func TestGetWorkspaces(t *testing.T) {
	t.Run("should return an empty list if there is no workspace available", func(t *testing.T) {
		td := NewTestDriver(t)
		foundWS, err := td.GetWorkspaces(userId1)

		assert.Nil(t, err)
		assert.Len(t, foundWS, 0)
		assert.Len(t, td.SavedWorkspaces(), 0)
	})

	t.Run("should return 2 workspaces", func(t *testing.T) {
		/// !!
		workspaceId := uuid.MustParse("54e8598c-900d-4888-b13e-44d9c70560f4")
		membershipId := uuid.MustParse("5ce691ed-a8b9-44ac-8e96-d56431c8d6a1")
		// ws1
		td := NewTestDriver(t).GivenUuids(workspaceId, membershipId)
		workspaceName1 := "my-workspace1"
		req := CreateWorkspaceCmd{
			Name:   workspaceName1,
			UserId: userId1,
		}
		_, err := td.CreateWorkspace(req)

		// ws2
		workspaceId2 := uuid.MustParse("338c130f-30b1-4c8b-b26f-de879a28a22e")
		membershipId2 := uuid.MustParse("d66b3364-0d3f-445e-94e8-72e7dcd73f5b")
		td.GivenUuids(workspaceId2, membershipId2)
		workspaceName2 := "my-workspace2"
		req2 := CreateWorkspaceCmd{
			Name:   workspaceName2,
			UserId: userId1,
		}
		_, err = td.CreateWorkspace(req2)

		fWorkspaces, err := td.GetWorkspaces(userId1)
		assert.Nil(t, err)
		assert.Len(t, fWorkspaces, 2)

		// verification in memory
		assert.Len(t, td.SavedWorkspaces(), 2)
	})
}

// !! delete
