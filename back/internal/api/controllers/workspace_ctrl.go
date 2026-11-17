package ctrl

import (
	"encoding/json"
	dto_idaccess_req "food/internal/api/http/dtos/id-access/req"
	dto_idaccess_res "food/internal/api/http/dtos/id-access/res"
	api_validator "food/internal/api/validator"
	idaccess_app "food/internal/modules/id-access/application"
	platform_http "food/internal/platform/http"
	json_utils "food/utils/json"
	"net/http"
)

type WorkspaceCtrl struct {
	workspaceSrv *idaccess_app.WorkspaceSrv
}

func NewWorkspaceCtrl(workspaceSrv *idaccess_app.WorkspaceSrv) *WorkspaceCtrl {
	return &WorkspaceCtrl{
		workspaceSrv: workspaceSrv,
	}
}

func (ictrl *WorkspaceCtrl) Create(w http.ResponseWriter, r *http.Request) {
	var req dto_idaccess_req.CreateWorkspace
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api_validator.Get().Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, _ := platform_http.GetUserIdFromContext(r.Context())
	createdW, err := ictrl.workspaceSrv.Create(r.Context(), idaccess_app.CreateWorkspaceCmd{
		Name:   req.Name,
		UserId: userId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.MembershipCreated{
		Id:   createdW.Id(),
		Name: createdW.Name(),
	}
	json_utils.JsonResponse(w, res, http.StatusCreated)

}
