package service

import (
	"context"
	"log"
	companyService "olympy/company-service/genproto/company_service"
	"olympy/company-service/internal/storage"
	"os"
)

type CompanyService struct {
	companyService.UnimplementedCompanyServiceServer
	companyStorage *storage.Company
	logger         *log.Logger
}

func NewCompanyService(companyStorage *storage.Company) *CompanyService {
	return &CompanyService{
		companyStorage: companyStorage,
		logger:         log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *companyService.CreateCompanyRequest) (*companyService.CreateCompanyResponse, error) {
	s.logger.Println("Create Company Request")
	return s.companyStorage.CreateCompany(ctx, req)
}

func (s *CompanyService) GetCompany(ctx context.Context, req *companyService.GetCompanyRequest) (*companyService.GetCompanyResponse, error) {
	s.logger.Println("Get Company Request")
	return s.companyStorage.GetCompany(ctx, req)
}

func (s *CompanyService) UpdateCompany(ctx context.Context, req *companyService.UpdateCompanyRequest) (*companyService.UpdateCompanyResponse, error) {
	s.logger.Println("Update Company Request")
	return s.companyStorage.UpdateCompany(ctx, req)
}

func (s *CompanyService) DeleteCompany(ctx context.Context, req *companyService.DeleteCompanyRequest) (*companyService.DeleteCompanyResponse, error) {
	s.logger.Println("Delete Company Request")
	return s.companyStorage.DeleteCompany(ctx, req)
}

func (s *CompanyService) ListCompanies(ctx context.Context, req *companyService.ListCompaniesRequest) (*companyService.ListCompaniesResponse, error) {
	s.logger.Println("List Companies Request")
	return s.companyStorage.ListCompanies(ctx, req)
}
