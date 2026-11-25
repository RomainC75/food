package ctrl

import (
	"errors"
	dto_idaccess_req "food/internal/api/http/dtos/id-access/req"
	dto_idaccess_res "food/internal/api/http/dtos/id-access/res"
	req_extractor "food/internal/api/http/req-extractor"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_dom "food/internal/modules/id-access/domain"
	platform_http "food/internal/platform/http"
	json_utils "food/utils/json"
	"net/http"
	"strings"
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
	req, status, err := req_extractor.OpenRequestExtractor[dto_idaccess_req.SignupUser](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	found, err := ictrl.userSrv.SignupUser(r.Context(), idaccess_app.SignupCmd{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.SignupResDto{
		UserId: found.Id(),
		Email:  found.Email(),
	}
	json_utils.JsonResponse(w, res, http.StatusCreated)

}

func (ictrl *IdAccessCtrl) Login(w http.ResponseWriter, r *http.Request) {
	req, status, err := req_extractor.OpenRequestExtractor[dto_idaccess_req.LoginUser](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	found, err := ictrl.userSrv.LoginUser(r.Context(), idaccess_app.LoginCmd{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.LoginResDto{
		UserId:       found.UserID,
		Email:        found.Email,
		AccessToken:  found.AccessToken,
		RefreshToken: found.RefreshToken,
	}
	json_utils.JsonResponse(w, res, http.StatusOK)
}

func (ictrl *IdAccessCtrl) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawToken, err := ictrl.extractBearer(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenData, err := ictrl.userSrv.Verify(rawToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := platform_http.WithUserId(r.Context(), tokenData.UserId)
		ctx = platform_http.WithUserEmail(ctx, tokenData.UserEmail)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (ictrl *IdAccessCtrl) Refresh(w http.ResponseWriter, r *http.Request) {
	req, status, err := req_extractor.OpenRequestExtractor[dto_idaccess_req.RefreshToken](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	accessToken, err := ictrl.userSrv.Refresh(req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto_idaccess_res.RefreshTokenResDto{
		AccessToken: accessToken,
	}
	json_utils.JsonResponse(w, res, http.StatusOK)
}

func (ictrl *IdAccessCtrl) ReturnAuthData(w http.ResponseWriter, r *http.Request) {
	email := platform_http.GetEmailFromContext(r)

	id, ok := platform_http.GetUserIdFromContext(r.Context())
	if !ok {
		http.Error(w, "could not parse token ()", http.StatusUnauthorized)
		return
	}
	data := idaccess_dom.VerifyInfo{
		UserId:    id,
		UserEmail: email,
	}

	json_utils.JsonResponse(w, data, http.StatusOK)
}

func (ictrl *IdAccessCtrl) extractBearer(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return "", errors.New("unauthorized")
	}

	return strings.TrimPrefix(header, "Bearer "), nil
}
