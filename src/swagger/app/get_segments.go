package app

import (
	apiUser "github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) GetSegmentsHandler(params apiUser.GetSegmentsParams) middleware.Responder {
	return middleware.NotImplemented("operation user GetSegments has not yet been implemented")
}
