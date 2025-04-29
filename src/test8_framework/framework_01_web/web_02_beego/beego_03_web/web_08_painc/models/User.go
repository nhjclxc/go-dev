package models

import "errors"

var (
	UserMap map[string]*User
)

func init() {
	UserMap = make(map[string]*User)
	u := User{"123", "astaxie", "11111"}
	UserMap["123"] = &u
}

type User struct {
	UserId   string
	Username string
	Password string
}

func GetUser(userId string) (u *User, err error) {
	if u, ok := UserMap[userId]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}
