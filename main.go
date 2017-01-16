package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
)

type TheatreWaitingList struct {
	gorm.Model
	ID                        int    `gorm:"ID"`
	StatusID                  int    `gorm:"STATUS_ID"`
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
		log.Fatal(err)
		return
	}

	rows, err := db.Model(&TheatreWaitingList{}).Raw("SELECT TOP 10 * FROM dbo.TheatreWaitingList").Rows() // (*sql.Rows, error)
	for rows.Next() {
		var twl TheatreWaitingList
		db.ScanRows(rows, &twl)
		log.Println(twl)
		// do something
	}

}
