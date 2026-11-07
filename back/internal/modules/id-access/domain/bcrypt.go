package idaccess_dom

var (
	ErrHashingPassword = "error trying to hash the password"
)

type Bcrypt interface {
	Generate(password string) (string, error)
	IsValid(hash string, password string) bool
}
