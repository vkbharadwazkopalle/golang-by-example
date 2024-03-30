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

func CreateInterviewsTable() {

	db := database.Db

	q := `CREATE TABLE IF NOT EXISTS interviews (
		interview_id SERIAL PRIMARY KEY,
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

func GetInterviews() gin.HandlerFunc {

	return func(c *gin.Context) {
		var result []models.Interview

		db := database.Db

		query, err := db.Query("SELECT * FROM interviews")

		var rowsReadErr bool
		for query.Next() {

			var name, created_on string
			var interview_id, active int

			err = query.Scan(&interview_id, &name, &active, &created_on)

			if err != nil {
				log.Printf("error occurred while reading the database rows: %v", err)
				rowsReadErr = true
				break
			}
			result = append(result, models.NewInterview(interview_id, name, active, created_on))
		}

		if rowsReadErr {
			log.Println("we are not able to fetch few records")
		}

		c.JSON(http.StatusOK, responses.CommonResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}

func AddNewInterview() gin.HandlerFunc {

	return func(c *gin.Context) {

		db := database.Db

		fmt.Println("db", db)

		CreateInterviewsTable()

		body := models.Interview{}

		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(400, "Interview is not defined")
			return
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			c.AbortWithStatusJSON(400, "Bad Input")
			return
		}

		fmt.Println(body)

		q := `INSERT INTO interviews VALUES($1, $2, $3, $4)`

		insert, err := db.Exec(q, body.Interview_Id, body.Name, body.Active, time.Now())

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

		// if no record was inserted, let us say client has failed
		if n == 0 {
			e := "could not insert the record, please try again after sometime"
			fmt.Println("failed to insert", e)
			c.AbortWithStatusJSON(400, e)
		}

		var result models.Interview = body

		c.JSON(http.StatusCreated, responses.CommonResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}

func makeGinResponse(c *gin.Context, i int, s string) {
	panic("unimplemented")
}
