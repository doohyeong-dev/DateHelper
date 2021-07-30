package api

import (
	u "github.com/doohyeong-dev/Tastery/apiHelpers"
	"github.com/doohyeong-dev/Tastery/models"
	res "github.com/doohyeong-dev/Tastery/resources/api"
)

//UserService struct
type UserService struct {
	User models.User
}

//UserList function returns the list of users
func (us *UserService) UserList() map[string]interface{} {
	user := us.User

	userData := res.UserResponse{
		ID:    user.ID,
		Name:  "test",
		Email: "test@gmail.com",
	}
	response := u.Message(0, "This is from version 1 api")
	response["data"] = userData
	return response
}
