package helpers

import (
	"awahids/go-opinia/model/domain"
	"awahids/go-opinia/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Fullname: user.Fullname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}
