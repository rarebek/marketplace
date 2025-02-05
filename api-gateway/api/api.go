package api

import (
	"log"

	companyhandlers "olympy/api-gateway/api/handlers/company-handlers"
	"olympy/api-gateway/api/middleware/casbin"
	"olympy/api-gateway/config"
	_ "olympy/api-gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	logger         *log.Logger
	cfg            *config.Config
	companyhandler *companyhandlers.CompanyHandlers
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	companyhandler *companyhandlers.CompanyHandlers,
) *API {
	return &API{
		logger:         logger,
		cfg:            cfg,
		companyhandler: companyhandler,
	}
}

// NewRoute
// NewRoute
// @title API
// @description TEST
// @BasePath /api/v1
// @version 1.7
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (a *API) RUN() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(casbin.NewAuthorizer())

	api := router.Group("/api/v1")
	{
		api.POST("/companies/add", a.companyhandler.CreateCompany)
		api.PUT("/companies/update", a.companyhandler.UpdateCompany)
		api.DELETE("/companies/delete", a.companyhandler.DeleteCompany)
		api.GET("/companies/get", a.companyhandler.GetCompany)
		api.GET("/companies/getall", a.companyhandler.ListCompanies)
	}

	return router.Run(a.cfg.ServerAddress)
}
