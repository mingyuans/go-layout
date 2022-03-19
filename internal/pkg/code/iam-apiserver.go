package code

//go:generate codegen -type=int

// apiserver: user errors.
const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 409: User already exist.
	ErrUserAlreadyExist
)
