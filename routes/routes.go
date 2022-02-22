package routes

import (
	"risevest/auth"
	"risevest/controllers"
	"risevest/utils"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Base Api end point
	api := app.Group("/api/v1")
	// Authentication end points
	_auth := api.Group("/auth")
	private := api.Group("/drive")
	_auth.Post("/login", auth.Login)
	_auth.Post("/register", auth.Register)

	private.Use(utils.SecureAuth())

	private.Post("/create-folder", controllers.CreateFolder)

	private.Post("/upload/folder/:folder", controllers.StoreFileInFolder)

	private.Post("/upload", controllers.StoreFile)

	private.Get("/download/:filename", controllers.DownloadFile)

	private.Get("/view/files", controllers.GetFile)

	private.Delete("/:fileID", controllers.DeleteFile)

}
