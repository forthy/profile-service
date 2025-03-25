package repository

import (
	"fmt"
	M "profile-svc/model"

	E "github.com/IBM/fp-go/either"
)

type ProfileNotFound struct {
	id M.Id
}

func (p ProfileNotFound) Error() string {
	return fmt.Sprintf("Profile not found with id: %s", p.id.Value)
}

type ProfileRepo interface {
	// RegisterProfile Create a new profile
	RegisterProfile(profile M.Profile) E.Either[error, M.Profile]
	ProfileWith(id M.Id) E.Either[error, M.Profile]
}
