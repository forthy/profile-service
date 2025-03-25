package fp_samples

import (
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
)

type UserId struct {
	id string
}
type User struct {
	Id   UserId
	Name string
}

type FindUserById = func(UserId) O.Option[User]

func findUserBy(uId UserId) O.Option[User] {
	return F.Ternary(
		func(uId UserId) bool {
			return uId.id == "UT-28474"
		},
		func(uId UserId) O.Option[User] {
			return O.Some(User{uId, "Richard Chuo"})
		},

		func(uId UserId) O.Option[User] {
			return O.None[User]()
		},
	)(uId)
}

func findUserByIdWithCount(count int) FindUserById {
	var counter = 0

	return func(uId UserId) O.Option[User] {
		var r O.Option[User] = O.None[User]()

		if counter > count {
			r = findUserBy(uId)
		}
		counter += 1
		return r
	}
}
