package dto

import "test/forumtest-api/model"

type Dto struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Telphone string `json:"telphone"`
}
func UserDto(user model.User) Dto {
	return Dto{
		Id:       user.ID,
		Name:     user.Name,
		Telphone: user.Telphone,
	}
}
