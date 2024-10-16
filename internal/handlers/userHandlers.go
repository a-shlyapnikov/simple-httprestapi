package handlers

import (
	"context"

	"github.com/a-shlyapnikov/simple-httprestapi/internal/userService"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/web/users"
	"github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	Service *userService.UserService
}

// DeleteUsersId implements users.StrictServerInterface.
func (u *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := u.Service.DeleteUser(request.Id)
	if err != nil {
		return users.DeleteUsersId404Response{}, err
	}
	return users.DeleteUsersId204Response{}, nil
}

// GetUsers implements users.StrictServerInterface.
func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		u := users.User{
			Email:    types.Email(user.Email),
			Id:       &user.ID,
			Password: user.Password,
		}
		response = append(response, u)
	}
	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (u *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	userToUpdate := userService.User{
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}

	updatedUser, err := u.Service.UpdateUser(request.Id, userToUpdate)
	if err != nil {
		return users.PatchUsersId404Response{}, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    types.Email(updatedUser.Email),
		Password: updatedUser.Password,
	}
	return response, nil

}

// PostUsers implements users.StrictServerInterface.
func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email: string(userRequest.Email),
		Password: userRequest.Password,
	}

	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id: &createdUser.ID,
		Email: types.Email(createdUser.Email),
		Password: createdUser.Password,
	}
	return response, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}
