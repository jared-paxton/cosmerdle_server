package user

import "time"

type User struct {
	Key   int
	UserId    string
	Email      string
	Password   string
	CreatedOn time.Time
	LastActivity time.Time
}

type NewUser struct {
	Email    string
	Password string
}

type AnonymousUser struct {
    Key int
    UserId string
    CreatedOn time.Time
    LastActivity time.Time
}

type NewAnonymousUser struct {
    UserId string
    CreatedOn time.Time
    LastActivity time.Time
}
