package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"ms-sql-demo/parser"
	"os"
)

type TheatreWaitingList struct {
	ID                        int    `gorm:"column:ID"`
	StatusID                  int    `gorm:"column:STATUS_ID"`
	OfferedOutcomeDescription string `gorm:"column:OFFERED_OUTCOME_DESCRIPTION"`
}

func main() {

	USER := os.Getenv("MS_USER")
	SERVER := os.Getenv("MS_SERVER")
	PORT := os.Getenv("MS_PORT")
	PASSWORD := os.Getenv("MS_PASSWORD")
	DB_NAME := os.Getenv("MS_DB")

	dbConnectionStr := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%s", SERVER, DB_NAME, USER, PASSWORD, PORT)
	log.Println("dbConnectionStr: ", dbConnectionStr)
	db, err := gorm.Open("mssql", dbConnectionStr)
	defer db.Close()
	if err != nil {
		log.Println("DB Connection Error ")
		log.Fatal(err)
		return
	}

	rows, err := db.Raw("SELECT TOP 10 * FROM dbo.TheatreWaitingList").Rows() // (*sql.Rows, error)
	defer rows.Close()

	if err != nil {
		log.Println("Query error")
		log.Fatal(err)
		return
	}

	for rows.Next() {
		var twl TheatreWaitingList
		db.ScanRows(rows, &twl)
		log.Println(twl)
		rowMap, err := parser.NewS(rows, nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		rowMap.Scan()
		objectMap := rowMap.Map()
		log.Println("objectMap: ", objectMap)
		// do something
	}

	// RowToMap()

}

func RowToMap() {
	USER := os.Getenv("MS_USER")
	SERVER := os.Getenv("MS_SERVER")
	PORT := os.Getenv("MS_PORT")
	PASSWORD := os.Getenv("MS_PASSWORD")
	DB_NAME := os.Getenv("MS_DB")
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%s", SERVER, DB_NAME, USER, PASSWORD, PORT)
	log.Println("connString: ", connString)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT TOP 10 * FROM dbo.TheatreWaitingList")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		rowMap, err := parser.NewS(rows, nil)
		if err != nil {
			log.Fatal(err)
			return
		}

		objectMap := rowMap.Map()
		log.Println(objectMap)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
