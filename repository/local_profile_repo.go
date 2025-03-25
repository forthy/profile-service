package repository

import (
	M "profile-svc/model"

	E "github.com/IBM/fp-go/either"
)

type LocalProfileRepo struct {
	profile map[M.Id]M.Profile
}

func LocalProfileRepoOf() LocalProfileRepo {
	return LocalProfileRepo{
		profile: make(map[M.Id]M.Profile),
	}
}

// RegisterProfile Profile -> Either error Profile
func (p *LocalProfileRepo) RegisterProfile(profile M.Profile) E.Either[error, M.Profile] {
	p.profile[profile.Id] = profile

	return E.Right[error](profile)
}

func (p *LocalProfileRepo) ProfileWith(id M.Id) E.Either[error, M.Profile] {
	profile, ok := p.profile[id]

	if !ok {
		return E.Left[M.Profile, error](ProfileNotFound{id})
	}

	return E.Right[error](profile)
}
