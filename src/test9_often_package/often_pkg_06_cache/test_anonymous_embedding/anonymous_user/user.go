package anonymous_user

type AnonymousUser struct {
	*user
}

type user struct {
	name string
	age  int
}

func NewUser(name string, age int) *AnonymousUser {
	return &AnonymousUser{user: &user{
		name: name,
		age:  age,
	}}
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}
