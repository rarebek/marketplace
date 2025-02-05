package models

// CompanySwagger represents the company model for swagger documentation
type CompanySwagger struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// CreateCompanyRequest represents the create company request for swagger
type CreateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

// CreateCompanyResponse represents the create company response for swagger
type CreateCompanyResponse struct {
	Company CompanySwagger `json:"company"`
	Message string         `json:"message"`
}

// UpdateCompanyRequest represents the update company request for swagger
type UpdateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

// UpdateCompanyResponse represents the update company response for swagger
type UpdateCompanyResponse struct {
	Company CompanySwagger `json:"company"`
	Message string         `json:"message"`
}

// DeleteCompanyResponse represents the delete company response for swagger
type DeleteCompanyResponse struct {
	Message string `json:"message"`
}

// GetCompanyResponse represents the get company response for swagger
type GetCompanyResponse struct {
	Company CompanySwagger `json:"company"`
}

// ListCompaniesResponse represents the list companies response for swagger
type ListCompaniesResponse struct {
	Companies []CompanySwagger `json:"companies"`
	Total     int64            `json:"total"`
	Page      int32            `json:"page"`
	Limit     int32            `json:"limit"`
}
