package companyhandlers

import (
	"log"
	companyservice "olympy/api-gateway/genproto/company_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyHandlers struct {
	client companyservice.CompanyServiceClient
	logger *log.Logger
}

func NewCompanyHandlers(client companyservice.CompanyServiceClient, logger *log.Logger) *CompanyHandlers {
	return &CompanyHandlers{
		client: client,
		logger: logger,
	}
}

// CreateCompany godoc
// @Summary Add a company
// @Description This endpoint adds a new company.
// @Tags Company
// @Accept json
// @Produce json
// @Param request body models.CreateCompanyRequest true "Company details to add"
// @Success 200 {object} models.CreateCompanyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /companies/add [post]
func (h *CompanyHandlers) CreateCompany(ctx *gin.Context) {
	var req companyservice.CreateCompanyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.CreateCompany(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// UpdateCompany godoc
// @Summary Edit a company
// @Description This endpoint edits an existing company.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Param request body models.UpdateCompanyRequest true "Company details to edit"
// @Success 200 {object} models.UpdateCompanyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /companies/{id} [put]
func (h *CompanyHandlers) UpdateCompany(ctx *gin.Context) {
	var req companyservice.UpdateCompanyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.UpdateCompany(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description This endpoint deletes a company by its ID.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {object} models.DeleteCompanyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /companies/{id} [delete]
func (h *CompanyHandlers) DeleteCompany(ctx *gin.Context) {
	idStr := ctx.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &companyservice.DeleteCompanyRequest{Id: id}

	resp, err := h.client.DeleteCompany(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetCompany godoc
// @Summary Get a company
// @Description This endpoint retrieves a company by its ID.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path int true "Company ID"
// @Success 200 {object} models.GetCompanyResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /companies/{id} [get]
func (h *CompanyHandlers) GetCompany(ctx *gin.Context) {
	idStr := ctx.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &companyservice.GetCompanyRequest{Id: id}

	resp, err := h.client.GetCompany(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ListCompanies godoc
// @Summary List companies
// @Description This endpoint retrieves all companies with pagination and search.
// @Tags Company
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param search query string false "Search term for company name or description"
// @Success 200 {object} models.ListCompaniesResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /companies [get]
func (h *CompanyHandlers) ListCompanies(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	search := ctx.Query("search")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid limit number"})
		return
	}

	req := &companyservice.ListCompaniesRequest{
		Page:   int32(page),
		Limit:  int32(limit),
		Search: search,
	}

	resp, err := h.client.ListCompanies(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
