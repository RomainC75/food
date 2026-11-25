package idaccess_infra_persist

import (
	"context"
	idaccess_app "food/internal/modules/id-access/application"
)

func PopulateUsers(userSrv *idaccess_app.UserSrv) {
	userSrv.LoginUser(context.Background(), idaccess_app.LoginCmd{
		Email:    "fake@email.com",
		Password: "azertyazerty",
	})
}
