package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"
	"github.com/pooulad/csv-2-sql-example/db"
	"github.com/pooulad/csv-2-sql-example/utils"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	exelAddressFlag := flag.String("address", "", "exelAddressFlag")

	flag.Parse()

	if *exelAddressFlag != "" {
		err = utils.ReadExcelAndInsertData(db, *exelAddressFlag)
		if err != nil {
			log.Fatal(err)
		}
		utils.Colorize(utils.ColorGreen, "Data inserted successfully!")
		utils.Colorize(utils.ColorBlue, "\u001b[36mMade by \U0001F9E1")
	} else {
		utils.Colorize(utils.ColorRed, "Please enter a valid address!")
	}

}
