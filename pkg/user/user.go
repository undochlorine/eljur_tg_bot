package user

import "errors"

const (
	InitState        = "init"
	WaitingFileState = "waiting for a file"
)

type Users map[int]*User

type User struct {
	state string
}

func (u *User) GetState() string {
	return u.state
}

func (u *User) SetState(state string) error {
	if !(state == InitState || state == WaitingFileState) {
		return errors.New("invalid state")
	}
	u.state = state
	return nil
}
