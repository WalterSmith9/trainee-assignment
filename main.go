package main

import (
	"log"

	"github.com/WalterSmith9/trainee-assignment/src/helpers/configs"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi"
	"github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations"

	apiSegment "github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/segment"

	apiUser "github.com/WalterSmith9/trainee-assignment/src/swagger/generated/restapi/operations/user"

	"github.com/WalterSmith9/trainee-assignment/src/swagger/app"
	"github.com/go-openapi/loads"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewTraineeAssignmentAPI(swaggerSpec)

	//api.APIKeyHeaderAuth = srv.APIKeyHeaderAuth

	api.UserGetSegmentsHandler = apiUser.GetSegmentsHandlerFunc(srv.GetSegmentsHandler)
	api.SegmentSegmentCreateHandler = apiSegment.SegmentCreateHandlerFunc(srv.SegmentCreateHandler)
	api.SegmentSegmentDeleteHandler = apiSegment.SegmentDeleteHandlerFunc(srv.SegmentDeleteHandler)
	api.UserUserEditHandler = apiUser.UserEditHandlerFunc(srv.UserEditHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	configs.InitConfig()

	server.ConfigureAPI()

	server.Port = configs.GetApiPort()
	server.EnabledListeners = configs.GetEnabledListeners()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
