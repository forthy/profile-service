package service

import (
	A "github.com/IBM/fp-go/array"
	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	J "github.com/IBM/fp-go/json"
	L "github.com/IBM/fp-go/lazy"
	P "github.com/IBM/fp-go/pair"
	W "github.com/IBM/fp-go/writer"

	"time"

	PR "profile-svc/model"
	TC "profile-svc/typeclass"
)

type ServiceEvent interface {
	EventTag() string
}

type Timestamp = time.Time

type RegisterEvent struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	Domain       string    `json:"domain"`
	RegisterTime Timestamp `json:"registerTime"`
}

func (r RegisterEvent) EventTag() string {
	return "RegisterEvent"
}

func registerAction2Event(tFn L.Lazy[time.Time]) func(PR.Profile) RegisterEvent {
	return func(profile PR.Profile) RegisterEvent {
		return RegisterEvent{
			Id:           profile.Id.Value,
			Name:         profile.Name.Value,
			Version:      profile.Version.Value,
			Domain:       profile.Domain.Title(),
			RegisterTime: tFn(),
		}
	}
}

// TODO - should not allow JSON marshalling error (?)
func RegisterEventShow(r RegisterEvent) string {
	return E.Fold(
		func(e error) string {
			return e.Error() // should not be possible
		},
		func(b []byte) string {
			return string(b)
		},
	)(J.Marshal(r))
}

// Writer w a :: () -> (a, w)
// type Writer[W, A any] IO.IO[P.Pair[A, W]]
// TODO - should add a dependency for transforming Profile to RegisterEvent?
func RegisterWithShow(show TC.Show[RegisterEvent]) func(PR.Profile) W.Writer[[]string, E.Either[error, PR.Profile]] {
	return func(profile PR.Profile) W.Writer[[]string, E.Either[error, PR.Profile]] {
		return func() P.Pair[E.Either[error, PR.Profile], []string] {
			description := F.Flow2(
				registerAction2Event(L.Of(time.Now())),
				show,
			)(profile)

			return P.MakePair(E.Right[error](profile), A.Of(description))
		}
	}
}
