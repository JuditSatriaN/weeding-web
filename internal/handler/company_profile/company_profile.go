package company_profile

import (
	"github.com/gofiber/fiber/v2"
	"github.com/weeding-web/internal/entity"
)

func CompanyProfileHandler(ctx *fiber.Ctx) error {

	return ctx.Render("company_profile/index", entity.WebData{
		Title:       entity.CompanyProfileTitle,
		TemplateURL: entity.TemplateUrl,
	})

}
