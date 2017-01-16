package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
)

type TheatreWaitingList struct {
	ID                        int    `bson:"specialId" gorm:"ID"`
	StatusID                  int    `bson:"" gorm:"STATUS_ID"`
	OfferedOutcomeDescription string `gorm:"OFFERED_OUTCOME_DESCRIPTION"`
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
		// var twl TheatreWaitingList
		var id int
		var offerId string
		rows.Scan(&id, &offerId)
		log.Println("id: ", id)
		log.Println("offerId: ", offerId)
		// do something
	}

}
