package req_extractor

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

func OpenRequestExtractor[o any](r *http.Request) (o, int, error) {
	var dto o
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return dto, http.StatusBadRequest, err
	}

	err = api_validator.Get().Struct(dto)
	if err != nil {
		return dto, http.StatusBadRequest, err
	}

	return dto, 0, nil
}

func WorkspaceRequestExtractor[o any](r *http.Request) (ExtractedWorkspaceRequestInfo[o], int, error) {
	dto, status, err := OpenRequestExtractor[o](r)
	if err != nil {
		return ExtractedWorkspaceRequestInfo[o]{}, status, err
	}

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		return ExtractedWorkspaceRequestInfo[o]{}, http.StatusBadRequest, err
	}
	return ExtractedWorkspaceRequestInfo[o]{
		WorkspaceId: workspaceId,
		Dto:         dto,
	}, 0, nil
}
