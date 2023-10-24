package reader

import (
	"context"
	"database/sql"
	"encoding/csv"
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
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	csvData := createList(data)

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, row := range csvData {
		cellValue1 := row.Username
		cellValue2 := row.Identifier
		cellValue3 := row.Firstname
		cellValue4 := row.Lastname

		_, err = tx.Exec(
			"INSERT INTO users (Username,Identifier,Firstname,Lastname) VALUES ($1,$2,$3,$4)",
			cellValue1,
			cellValue2,
			cellValue3,
			cellValue4,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func createList(data [][]string) []ListRecord {
	var finalList []ListRecord
	for i, line := range data {
		var rec ListRecord
		if i == 0 {
			continue
		} else {
			for j, field := range line {
				switch j {
				case 0:
					rec.Username = field
				case 1:
					rec.Identifier = field
				case 2:
					rec.Firstname = field
				case 3:
					rec.Lastname = field
				}
			}
		}
		finalList = append(finalList, rec)
	}
	return finalList
}
