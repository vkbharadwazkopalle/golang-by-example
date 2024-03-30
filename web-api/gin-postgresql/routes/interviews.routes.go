package routes

import (
	"web-api/gin-postgresql/controllers"

	"github.com/gin-gonic/gin"
)

func InterviewsRoutes(router *gin.Engine) {
	router.GET("/api/interviews", controllers.GetInterviews())
	router.POST("/api/interviews", controllers.AddNewInterview())
}
