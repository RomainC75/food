package idaccess_dom

var (
	ErrCouldNotGenerateJwt = "could not generate jwt"
)

type JWT interface {
	Generate(user *User) (string, error)
}
