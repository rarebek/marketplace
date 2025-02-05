package api

import (
	"log"
	"net"

	"olympy/company-service/internal/config"
	companyService "olympy/company-service/genproto/company_service"

	"google.golang.org/grpc"
)

type (
	API struct {
		service companyService.CompanyServiceServer
	}
)

func New(service companyService.CompanyServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	companyService.RegisterCompanyServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
