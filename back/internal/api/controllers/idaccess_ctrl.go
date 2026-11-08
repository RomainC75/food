package ctrl

import (
	"context"
	"encoding/json"
	"errors"
	"food/config"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	api_validator "food/internal/api/validator"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	shared_infra "food/internal/shared/infra"
	json_utils "food/utils/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type IdAccessCtrl struct {
	idaccessSrv *idaccess_app.IdAccessSrv
}

func NewIdAccessCtrl() *IdAccessCtrl {
	return &IdAccessCtrl{
		idaccessSrv: idaccess_app.NewIdAccessSrv(
			shared_infra.NewUuidGenerator(),
			idaccess_infra.NewBcrypt(),
			idaccess_infra.NewJWT(config.Get().Jwt.Secret, time.Duration(config.Get().Jwt.ExpiresInMin)),
			idaccess_infra.NewInMemoryUsers(),
		),
	}
}

func (ictrl *IdAccessCtrl) Signup(w http.ResponseWriter, r *http.Request) {
	var req dto_idaccess_req.SignupUser
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

	ctx := context.Background()
	found, err := ictrl.idaccessSrv.SignupUser(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found)
}

func (ictrl *IdAccessCtrl) Signin(w http.ResponseWriter, r *http.Request) {
	var req dto_idaccess_req.SigninUser
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

	ctx := context.Background()
	found, err := ictrl.idaccessSrv.SigninUser(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found)
}

func (ictrl *IdAccessCtrl) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken, err := ictrl.extractBearer(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tokenData, err := ictrl.idaccessSrv.Verify(rawToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), idaccess_dom.TokenUserIdKey, tokenData.UserId)
		ctx = context.WithValue(ctx, idaccess_dom.TokenEmailKey, tokenData.UserEmail)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ictrl *IdAccessCtrl) ReturnAuthData(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value(idaccess_dom.TokenEmailKey)
	userId := r.Context().Value(idaccess_dom.TokenUserIdKey)
	id, err := uuid.Parse(userId.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	data := idaccess_dom.VerifyInfo{
		UserId:    id,
		UserEmail: email.(string),
	}

	json_utils.JsonResponse(w, data)
}

func (ictrl *IdAccessCtrl) extractBearer(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("unauthorized")
	}

	return strings.TrimPrefix(header, "Bearer "), nil
}
