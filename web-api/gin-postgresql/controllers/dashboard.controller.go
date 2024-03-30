package controllers

import (
	"net/http"
	"web-api/gin-postgresql/models"
	"web-api/gin-postgresql/responses"

	"github.com/gin-gonic/gin"
)

func GetDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {

		result := []models.Dashboard{
			{
				Id:   "123",
				Name: "Projects",
				Data: models.Snapshot{40, 23},
			},
			{
				Id:   "456",
				Name: "Interviews",
				Data: models.Snapshot{40, 23},
			},
			{
				Id:   "789",
				Name: "Positions",
				Data: models.Snapshot{40, 23},
			},
		}

		c.JSON(http.StatusCreated, responses.CommonResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"result": result}})
	}
}
