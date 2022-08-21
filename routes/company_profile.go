package routes

import (
	"github.com/gofiber/fiber/v2"

	companyProfileSvc "github.com/weeding-web/internal/handler/company_profile"
)

func BuildCompanyProfileRoutes(service fiber.Router) {
	service.Get("/", companyProfileSvc.CompanyProfileHandler)
}
