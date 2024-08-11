package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(controller *controllers.Controller) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)

	// Protected routes
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.JWTAuthMiddleware()) // Applying JWT middleware to protected routes
	{
		authRoutes.GET("/tasks", controller.GetTasks)
		authRoutes.GET("/tasks/:id", controller.GetTask)

		admin := r.Group("/")
		admin.Use(middleware.RequireAdmin())
		{
			admin.POST("/promote/:id", controller.PromoteUser)
			admin.POST("/tasks", controller.CreateTask)
			admin.PUT("/tasks/:id", controller.UpdateTask)
			admin.DELETE("/tasks/:id", controller.DeleteTask)
			admin.POST("/promote/:id", controller.PromoteUser)
		}
	}

	return r
}
