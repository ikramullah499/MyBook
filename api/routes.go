package api

import (
	"MyBook/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repo repository.BookRepository
}

func SetupRoutes(app *fiber.App, repo repository.BookRepository) {
	h := Handler{Repo: repo}
	app.Post("/books", h.CreateBook)
	app.Get("/books", h.GetAllBooks)
	app.Get("/books/:id", h.GetBookByID)
	app.Delete("/books/:id", h.DeleteBookByID)
}

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var req repository.AddBook
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input format"})
	}
	if req.Title == "" || req.Author == "" || req.Year <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title, Author, and valid Year are required"})
	}
	book := repository.Book{Title: req.Title, Author: req.Author, Year: req.Year}
	if err := h.Repo.Create(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *Handler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.Repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve books"})
	}
	return c.JSON(books)
}

func (h *Handler) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.Repo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func (h *Handler) DeleteBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
	}
	return c.JSON(fiber.Map{"message": "Book deleted successfully"})
}
