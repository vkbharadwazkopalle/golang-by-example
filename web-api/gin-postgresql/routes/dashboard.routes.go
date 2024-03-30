package routes

import (
	"web-api/gin-postgresql/controllers"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.Engine) {

	router.GET("/api/dashboard", controllers.GetDashboard())

}
