package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-gorm/database"
	"github.com/sixfwa/fiber-gorm/models"
)

type Page struct {
	ID     uint   `json:"id"`
	Body   string `json:"body"`
	IdPage string `json:"id_page"`
}

func CreateResponsePage(pageModel models.Page) Page {
	return Page{ID: pageModel.ID, Body: pageModel.Body, IdPage: pageModel.IdPage}
}

func CreatePage(c *fiber.Ctx) error {
	var page models.Page

	if err := c.BodyParser(&page); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&page)
	responsePage := CreateResponsePage(page)

	return c.Status(400).JSON(responsePage)
}

func GetPages(c *fiber.Ctx) error {
	pages := []models.Page{}

	database.Database.Db.Find(&pages)
	responsePages := []Page{}

	for _, page := range pages {
		responsePage := CreateResponsePage(page)
		responsePages = append(responsePages, responsePage)
	}
	return c.Status(200).JSON(responsePages)
}

func findPage(id int, page *models.Page) error {
	database.Database.Db.Find(&page, "id = ?", id)
	if page.ID == 0 {
		return errors.New("page doesn't exist")
	}
	return nil
}
func GetPage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var page models.Page

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findPage(id, &page); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsePage := CreateResponsePage(page)

	return c.Status(200).JSON(responsePage)
}

func UpdatePage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var page models.Page

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findPage(id, &page); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdatePage struct {
		Body   string `json:"body"`
		IdPage string `json:"id_page"`
	}

	var updateData UpdatePage

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	page.Body = updateData.Body
	page.IdPage = updateData.IdPage

	database.Database.Db.Save(&page)

	responsePage := CreateResponsePage(page)
	return c.Status(200).JSON(responsePage)
}

func DeletePage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var page models.Page

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findPage(id, &page); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&page).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Deleted Page")
}
