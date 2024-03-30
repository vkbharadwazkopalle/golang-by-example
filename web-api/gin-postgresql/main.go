package main

import (
	"web-api/gin-postgresql/database"
	"web-api/gin-postgresql/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {

	database.ConnectDB()

	router := gin.Default()

	router.Use(cors.New(CORSConfig()))

	routes.DashboardRoutes(router)
	routes.InterviewsRoutes(router)
	routes.ProjectsRoutes(router)

	router.Run()
}
