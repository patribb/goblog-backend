package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patribb/blogbackend/controllers"
	"github.com/patribb/blogbackend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controllers.CreatePost)
	app.Get("/api/allposts", controllers.AllPosts)
	app.Get("/api/allposts/:id", controllers.DetailPost)
	app.Put("/api/updatepost/:id", controllers.UpdatePost)
	app.Get("api/uniquepost", controllers.UniquePost)
	app.Delete("api/deletepost/:id", controllers.DeletePost)
	app.Post("/api/upload-image", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
