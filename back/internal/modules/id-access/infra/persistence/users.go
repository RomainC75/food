package idaccess_infra_persist

import (
	"context"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_app "food/internal/modules/id-access/application"
	shared_infra "food/internal/shared/infra"

	"github.com/google/uuid"
)

func PopulateUsers(uuidGen *shared_infra.InMemUuidGen, userSrv *idaccess_app.UserSrv) {
	uuidGen.SetNext(uuid.MustParse("28ebf9da-e385-40b7-bf9a-33b1b1ef735f"))
	userSrv.SigninUser(context.Background(), dto_idaccess_req.SigninUser{
		Email:    "fake@email.com",
		Password: "azert",
	})
}
