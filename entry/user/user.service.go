package user

type userService struct {}

func (service *userService) findUsers(page int, limit int) *User {
	user := &User{}
	user.findUsers(page, limit)
	return user
}

// UserService 用户service层单一实列
var UserService = &userService{}
