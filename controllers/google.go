package controllers

import (
	"github.com/InfamousFreak/Tech-Task-24/config"
	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}
