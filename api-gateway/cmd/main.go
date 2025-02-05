package main

import (
	"log"
	"olympy/api-gateway/api"
	"olympy/api-gateway/config"
	companyservice "olympy/api-gateway/genproto/company_service"
	"os"

	"google.golang.org/grpc"

	companyhandlers "olympy/api-gateway/api/handlers/company-handlers"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	connCompany, err := grpc.Dial(cfg.CompanyServiceHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to company service: %v", err)
	}
	defer connCompany.Close()

	// Creating clients for services
	companyClient := companyservice.NewCompanyServiceClient(connCompany)
	// Creating handler instances
	companyHandlers := companyhandlers.NewCompanyHandlers(companyClient, logger)
	// Creating API instance
	api := api.New(cfg, logger, companyHandlers)
	logger.Fatal(api.RUN())
}
