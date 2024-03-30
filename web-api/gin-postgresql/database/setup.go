package database

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"web-api/gin-postgresql/configs"
)

var (
	// DBCon is the connection handle
	// for the database
	Db *sql.DB
)

func ConnectDB() {

	connStr := configs.GetDBURI()

	fmt.Println("connStr", connStr)
	var err error
	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		panic("@#@#@#@#Error on Setup.GO. failed to connect database")
	}

	// defer Db.Close()

	var version string
	if err := Db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)

}
