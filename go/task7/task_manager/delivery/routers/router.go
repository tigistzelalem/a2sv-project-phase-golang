package routers

import (
	"task-manager/delivery/controllers"
	"task-manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(ctrl *controllers.Controller, jwtSvc infrastructure.JWTService) *gin.Engine{
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/register", ctrl.RegisterUser)
		api.POST("/login", ctrl.LoginUser)

		protected := api.Group("/", infrastructure.AuthMiddleware(jwtSvc))
		{
			protected.PUT("/promote/:id", ctrl.PromoteUser)
			protected.GET("/tasks", ctrl.GetAllTasks)
			protected.GET("/tasks/:id", ctrl.GetTaskByID)
			protected.POST("/tasks", ctrl.CreateTask)
			protected.PUT("/tasks/:id", ctrl.UpdateTask)
			protected.DELETE("/tasks/:id", ctrl.DeleteTask)

		}
	}

	return r

}