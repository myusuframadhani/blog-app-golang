package routes

import (
	"blog-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Auth Routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
	}

	// Post Routes
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", controllers.CreatePost)
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}

	// Category Routes
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("/", controllers.CreateCategory)
		categoryRoutes.GET("/", controllers.GetCategories)
		categoryRoutes.GET("/:id", controllers.GetCategory)
		categoryRoutes.PUT("/:id", controllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
	}

	// Comment Routes
	commentRoutes := router.Group("/comments")
	{
		commentRoutes.POST("/", controllers.CreateComment)
		commentRoutes.GET("/", controllers.GetComments)
		commentRoutes.GET("/:id", controllers.GetComment)
		commentRoutes.PUT("/:id", controllers.UpdateComment)
		commentRoutes.DELETE("/:id", controllers.DeleteComment)
	}

	// Tag Routes
	tagRoutes := router.Group("/tags")
	{
		tagRoutes.POST("/", controllers.CreateTag)
		tagRoutes.GET("/", controllers.GetTags)
		tagRoutes.GET("/:id", controllers.GetTag)
		tagRoutes.PUT("/:id", controllers.UpdateTag)
		tagRoutes.DELETE("/:id", controllers.DeleteTag)
	}
}
