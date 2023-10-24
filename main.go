package main

import (
	"exel2sql/db"
	"exel2sql/utils"
	"flag"
	_ "github.com/lib/pq"
	"log"
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
