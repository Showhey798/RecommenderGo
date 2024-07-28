package entity

type UserID string
type Password string

type User struct {
	ID       UserID
	Email    Email
	Password Password
}

func (u *User) Validate() bool {
	return u.Email.CheckAvailableEmail()
}
