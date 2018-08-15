package user

type GetUserById func(id string) *User

func findById() GetUserById {
	return func(id string) *User {
		return &User{id, "Bob"}
	}
}
