package repos

import "fmt"

type ErrorCode int

const (
	SqlErrorCode ErrorCode = iota
	SaveErrorCode
	UpdateErrorCode
	UpsertErrorCode
	DeleteErrorCode
	HashPasswordErrorCode
	CustomerLoginErrorCode
	GetCustomerByEmailErrorCode
)

type RepoError struct {
	StatusCode ErrorCode
	Err        error
}

func (r *RepoError) Error() string {
	return fmt.Sprintf("status %d: error %v", r.StatusCode, r.Err)
}
