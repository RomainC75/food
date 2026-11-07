package idaccess_dom

var (
	ErrHashingPassword  = "error trying to hash the password"
	ErrNotValidPassword = "error password not valid"
)

type Bcrypt interface {
	Generate(password string) (string, error)
	IsValid(hash string, password string) bool
}
