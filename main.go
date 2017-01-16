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

	user := os.Getenv("MS_USER")
	host := os.Getenv("MS_HOST")
	password := os.Getenv("MS_PASSWORD")
	dbName := os.Getenv("MS_DB")
	dbConnectionStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password)
	log.Println("dbConnectionStr: ", dbConnectionStr)
	db, err := gorm.Open("mssql", dbConnectionStr)
	defer db.Close()
	if err != nil {
		log.log(err)
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
