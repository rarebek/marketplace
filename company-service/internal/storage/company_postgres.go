package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	companyService "olympy/company-service/genproto/company_service"
	"olympy/company-service/internal/config"

	"github.com/Masterminds/squirrel"
	//"github.com/google/uuid"
)

type Company struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func NewCompanyService(config *config.Config) (*Company, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Company{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (a *Company) AddCompany(ctx context.Context, req *companyService.Company) (*companyService.Company, error) {
	data := map[string]interface{}{
		"id":         req.Id,
		"name":       req.Name,
		"created_at": time.Now(),
	}

	query, args, err := a.queryBuilder.Insert("companies").
		SetMap(data).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var id int64

	if err := a.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &companyService.Company{
		Id:        id,
		Name:      req.Name,
		CreatedAt: data["created_at"].(time.Time).String(),
	}, nil
}

func (a *Company) EditCompany(ctx context.Context, req *companyService.Company) (*companyService.Company, error) {
	data := map[string]interface{}{
		"name":       req.Name,
		"updated_at": time.Now(),
	}

	query, args, err := a.queryBuilder.Update("companies").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := a.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	var updatedCompany companyService.Company
	err = a.db.QueryRowContext(ctx, "SELECT id, name, created_at, updated_at FROM companies WHERE id = $1", req.Id).
		Scan(&updatedCompany.Id, &updatedCompany.Name, &updatedCompany.CreatedAt, &updatedCompany.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated company: %v", err)
	}

	return &updatedCompany, nil
}

func (a *Company) DeleteCompany(ctx context.Context, req *companyService.DeleteCompanyRequest) (*companyService.DeleteCompanyResponse, error) {
	query, args, err := a.queryBuilder.Delete("companies").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("athlete with ID %d not found", req.Id)
	}

	return &companyService.DeleteCompanyResponse{
		Message: fmt.Sprintf("Company with ID %d deleted successfully", req.Id),
	}, nil
}

func (a *Company) CreateCompany(ctx context.Context, req *companyService.CreateCompanyRequest) (*companyService.CreateCompanyResponse, error) {
	data := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"address":     req.Address,
		"website":     req.Website,
		"phone":       req.Phone,
		"email":       req.Email,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	query, args, err := a.queryBuilder.Insert("companies").
		SetMap(data).
		Suffix("RETURNING id, name, description, address, website, phone, email, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var company companyService.Company
	err = a.db.QueryRowContext(ctx, query, args...).Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.Address,
		&company.Website,
		&company.Phone,
		&company.Email,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &companyService.CreateCompanyResponse{
		Company: &company,
		Message: "Company created successfully",
	}, nil
}

func (a *Company) GetCompany(ctx context.Context, req *companyService.GetCompanyRequest) (*companyService.GetCompanyResponse, error) {
	query, args, err := a.queryBuilder.Select(
		"id", "name", "description", "address", "website", "phone", "email", "created_at", "updated_at",
	).From("companies").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var company companyService.Company
	err = a.db.QueryRowContext(ctx, query, args...).Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.Address,
		&company.Website,
		&company.Phone,
		&company.Email,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("company with ID %d not found", req.Id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &companyService.GetCompanyResponse{
		Company: &company,
	}, nil
}

func (a *Company) UpdateCompany(ctx context.Context, req *companyService.UpdateCompanyRequest) (*companyService.UpdateCompanyResponse, error) {
	data := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"address":     req.Address,
		"website":     req.Website,
		"phone":       req.Phone,
		"email":       req.Email,
		"updated_at":  time.Now(),
	}

	query, args, err := a.queryBuilder.Update("companies").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, name, description, address, website, phone, email, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var company companyService.Company
	err = a.db.QueryRowContext(ctx, query, args...).Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.Address,
		&company.Website,
		&company.Phone,
		&company.Email,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("company with ID %d not found", req.Id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &companyService.UpdateCompanyResponse{
		Company: &company,
		Message: "Company updated successfully",
	}, nil
}

func (a *Company) ListCompanies(ctx context.Context, req *companyService.ListCompaniesRequest) (*companyService.ListCompaniesResponse, error) {
	baseQuery := a.queryBuilder.Select(
		"id", "name", "description", "address", "website", "phone", "email", "created_at", "updated_at",
	).From("companies")

	if req.Search != "" {
		baseQuery = baseQuery.Where(squirrel.Or{
			squirrel.Like{"name": "%" + req.Search + "%"},
			squirrel.Like{"description": "%" + req.Search + "%"},
		})
	}

	countQuery, countArgs, err := baseQuery.Column("COUNT(*)").ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build count query: %v", err)
	}

	var total int64
	err = a.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %v", err)
	}

	offset := (req.Page - 1) * req.Limit
	baseQuery = baseQuery.Limit(uint64(req.Limit)).Offset(uint64(offset))

	query, args, err := baseQuery.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	rows, err := a.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var companies []*companyService.Company
	for rows.Next() {
		var company companyService.Company
		err := rows.Scan(
			&company.Id,
			&company.Name,
			&company.Description,
			&company.Address,
			&company.Website,
			&company.Phone,
			&company.Email,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		companies = append(companies, &company)
	}

	return &companyService.ListCompaniesResponse{
		Companies: companies,
		Total:     total,
		Page:      req.Page,
		Limit:     req.Limit,
	}, nil
}
