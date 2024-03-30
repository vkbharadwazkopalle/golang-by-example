package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-api/gin-postgresql/database"
	"web-api/gin-postgresql/models"
	"web-api/gin-postgresql/responses"

	"github.com/gin-gonic/gin"
)

func CreateProjectsTable() {

	db := database.Db

	q := `CREATE TABLE IF NOT EXISTS projects (
		project_id SERIAL PRIMARY KEY,
		name varchar(100) NOT NULL,
		active INT NOT NULL DEFAULT 1,
		created_on timestamp NOT NULL DEFAULT NOW()
	)`

	create, err := db.Exec(q)

	if err != nil {
		// log.Printf("error occurred while inserting new record into artist table: %v", err)
		// makeGinResponse(c, http.StatusInternalServerError, err.Error())
		// return
		fmt.Print("failed to create")
		panic(err.Error())
	}

	fmt.Print("success", create)
}

func AddNewProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.Db

		CreateProjectsTable()
		body := models.Project{}

		data, err := c.GetRawData()

		if err != nil {
			c.AbortWithStatusJSON(400, "Project is not defined")
			return
		}
		err = json.Unmarshal(data, &body)
		if err != nil {
			c.AbortWithStatusJSON(400, "Bad Input")
			return
		}
		fmt.Println(body)
		q := `INSERT INTO projects VALUES($1, $2, $3, $4)`
		insert, err := db.Exec(q, body.Project_Id, body.Name, body.Active, time.Now())

		if err != nil {
			fmt.Print("failed to insert")
			c.AbortWithStatusJSON(400, "Bad Input")
			panic(err.Error())
		}

		fmt.Print("success", insert)

		// checking the number of rows affected
		n, err1 := insert.RowsAffected()
		if err1 != nil {
			// 	log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
			// 	makeGinResponse(c, http.StatusInternalServerError, err.Error())
			// 	return
			fmt.Println("failed to check", err1.Error())
		}

		if n == 0 {
			c.AbortWithStatusJSON(400, "Insert Failed")
			return
		}

		var result models.Project = body

		c.JSON(http.StatusCreated, responses.CommonResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}

func GetProjects() gin.HandlerFunc {

	return func(c *gin.Context) {
		var result []models.Project

		db := database.Db

		query, err := db.Query("SELECT * FROM projects")

		var rowsReadErr bool
		for query.Next() {

			var name, created_on string
			var project_id, active int

			err = query.Scan(&project_id, &name, &active, &created_on)

			if err != nil {
				log.Printf("error occurred while reading the database rows: %v", err)
				rowsReadErr = true
				break
			}
			result = append(result, models.NewProject(project_id, name, active, created_on))
		}

		if rowsReadErr {
			log.Println("we are not able to fetch few records")
		}

		c.JSON(http.StatusOK, responses.CommonResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}
