package app

import (
	"github.com/WalterSmith9/trainee-assignment/src/models/user"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/models"
	apiUser "github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) UserEditHandler(params apiUser.UserEditParams) middleware.Responder {
	//return middleware.NotImplemented("operation user UserEdit has not yet been implemented")

	// Edit User
	err := user.EditUser(params.UserInfo)
	if err != nil {
		return apiUser.NewUserEditInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}

	return apiUser.NewUserEditOK().WithPayload(&models.SuccessResponse{
		Code:    1,
		Message: "success",
	})
}
