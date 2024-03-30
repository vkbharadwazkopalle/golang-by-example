package routes

import (
	"web-api/gin-postgresql/controllers"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(router *gin.Engine) {

	router.GET("/api/projects", controllers.GetProjects())
	router.POST("/api/projects", controllers.AddNewProject())

}
