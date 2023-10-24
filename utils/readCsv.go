package utils

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
)

type ListRecord struct {
	Username   string
	Identifier string
	Firstname  string
	Lastname   string
}

func ReadExcelAndInsertData(db *sql.DB, fileAddress string) error {
	f, err := os.Open(fileAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	csvData := createList(data)

	for _, row := range csvData {
		cellValue1 := row.Username
		cellValue2 := row.Identifier
		cellValue3 := row.Firstname
		cellValue4 := row.Lastname

		_, err := db.Exec("INSERT INTO users (Username,Identifier,Firstname,Lastname) VALUES ($1,$2,$3,$4)", cellValue1, cellValue2, cellValue3, cellValue4)
		if err != nil {
			return err
		}

	}

	return nil
}

func createList(data [][]string) []ListRecord {
	var fianlList []ListRecord
	for i, line := range data {
		if i > 0 {
			var rec ListRecord
			for j, field := range line {
				if j == 0 {
					rec.Username = field
				} else if j == 1 {
					rec.Identifier = field
				} else if j == 2 {
					rec.Firstname = field
				} else if j == 3 {
					rec.Lastname = field
				}
			}
			fianlList = append(fianlList, rec)
		}
	}
	return fianlList
}
