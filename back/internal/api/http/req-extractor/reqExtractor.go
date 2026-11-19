package req_Extractor

import (
	"encoding/json"
	api_validator "food/internal/api/validator"
	platform_http "food/internal/platform/http"
	"net/http"

	"github.com/google/uuid"
)

type ExtractedWorkspaceRequestInfo[o any] struct {
	WorkspaceId uuid.UUID
	Dto         o
}

func WorkspaceRequestExtractor[o any](r *http.Request) (ExtractedWorkspaceRequestInfo[o], int, error) {
	var req o
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return ExtractedWorkspaceRequestInfo[o]{}, http.StatusBadRequest, err
	}

	err = api_validator.Get().Struct(req)
	if err != nil {
		return ExtractedWorkspaceRequestInfo[o]{}, http.StatusBadRequest, err
	}

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		return ExtractedWorkspaceRequestInfo[o]{}, http.StatusBadRequest, err
	}
	return ExtractedWorkspaceRequestInfo[o]{
		WorkspaceId: workspaceId,
		Dto:         req,
	}, 0, nil
}
