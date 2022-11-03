package app

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/", index) // Главная страница
	
	app.Get("/animal/create", createAnimalGet) // Страница создания нового животного
	app.Post("/animal/create", createAnimal) // Добавление нового животного

	app.Get("/animals/:id?", readAnimal) // Страница вывода данных о множестве животных, либо одном с сохранением шаблона

	app.Get("/animal/update/:id", updateAnimalGet) // Страница вывода данных для обновления данных о животном
	app.Post("/animal/update/:id", updateAnimal) // Обновление данных о животном

	app.Get("/animal/delete/:id", deleteAnimal) // Страница, которая вызывается для удаления данных о животном
}
