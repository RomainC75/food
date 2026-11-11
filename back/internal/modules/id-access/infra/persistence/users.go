package idaccess_infra_persist

import (
	"context"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_app "food/internal/modules/id-access/application"
)

func PopulateUsers(userSrv *idaccess_app.UserSrv) {
	userSrv.SigninUser(context.Background(), dto_idaccess_req.SigninUser{
		Email:    "fake@email.com",
		Password: "azert",
	})
}
