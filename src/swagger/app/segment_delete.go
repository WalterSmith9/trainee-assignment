package app

import (
	"github.com/WalterSmith9/trainee-assignment/src/models/segment"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/models"
	apiSegment "github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/segment"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) SegmentDeleteHandler(params apiSegment.SegmentDeleteParams) middleware.Responder {

	// Remove segment
	err := segment.DeleteSegment(params.SegmentInfo)
	if err != nil {
		return apiSegment.NewSegmentCreateInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
	}

	return apiSegment.NewSegmentCreateOK().WithPayload(&models.SuccessResponse{
		Code:    1,
		Message: "success",
	})
}
