package ctrl

import (
	json_utils "food/utils/json"
	"net/http"
)

type HealthCtrl struct {
}

func NewHealthCtrl() *HealthCtrl {
	return &HealthCtrl{}
}

func (ictrl *HealthCtrl) Health(w http.ResponseWriter, r *http.Request) {
	json_utils.JsonResponse(w, "healthy", http.StatusOK)
}
