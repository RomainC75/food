package dto_idaccess_req

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=30"`
}

type SignupUser = Credentials
type SigninUser = Credentials

type VerifyUser struct {
	Token string `json:"token" validate:"required"`
}
