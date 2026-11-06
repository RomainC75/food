package idaccess_dom

var (
	ErrGeneratingBcrypt = "error trying to hash the password"
)

type Bcrypt interface {
	Generate(password string) (string, error)
	IsValid(hash string, password string) bool
}
