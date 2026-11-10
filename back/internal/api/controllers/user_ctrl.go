package ctrl

import (
	"context"
	"encoding/json"
	"errors"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	dto_idaccess_res "food/internal/api/dtos/id-access/res"
	api_validator "food/internal/api/validator"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_dom "food/internal/modules/id-access/domain"
	json_utils "food/utils/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type IdAccessCtrl struct {
	userSrv *idaccess_app.UserSrv
}

func NewIdAccessCtrl(userSrv *idaccess_app.UserSrv) *IdAccessCtrl {
	return &IdAccessCtrl{
		userSrv: userSrv,
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

	found, err := ictrl.userSrv.SignupUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.SignupResDto{
		UserId: found.Id(),
		Email:  found.Email(),
	}
	json_utils.JsonResponse(w, res)

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

	found, err := ictrl.userSrv.SigninUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.SigninResDto{
		UserId: found.UserID,
		Email:  found.Email,
		Token:  found.Token,
	}
	json_utils.JsonResponse(w, res)
}

func (ictrl *IdAccessCtrl) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken, err := ictrl.extractBearer(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenData, err := ictrl.userSrv.Verify(rawToken)
		log.Println("---> ", rawToken, tokenData, err)
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

	id, ok := userId.(uuid.UUID)
	if !ok {
		http.Error(w, "could not parse token ()", http.StatusUnauthorized)
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
