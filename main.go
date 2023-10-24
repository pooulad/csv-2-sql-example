package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/savioxavier/termlink"

	"github.com/pooulad/csv-2-sql-example/database"
	"github.com/pooulad/csv-2-sql-example/reader"
	"github.com/pooulad/csv-2-sql-example/utils"
)

func main() {
	var file, cfg string
	flag.StringVar(&file, "file", "./username.csv", "CSV file path")
	flag.StringVar(&cfg, "config", "./config.json", "Config file path")

	flag.Parse()

	if file == "" {
		utils.Colorize(utils.ColorRed, "Please enter a valid file address!")
		return
	}

	config, err := database.ReadConfig(cfg)
	if err != nil {
		utils.Colorize(utils.ColorRed, err.Error())
		return
	}

	db, err := database.ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = reader.ReadExcelAndInsertData(db, file)
	if err != nil {
		utils.Colorize(utils.ColorRed, err.Error())
		return
	}

	utils.Colorize(utils.ColorGreen, "Data inserted successfully!")
	utils.Colorize(utils.ColorBlue, "\u001b[36mFeel free to show your appreciation by starring this GitHub repository if you found it helpful or valuable\U0001F9E1.")
	fmt.Println(termlink.ColorLink("Follow hossein1376", "https://github.com/hossein1376", "italic green"))
	fmt.Println(termlink.ColorLink("Follow pooulad", "https://github.com/pooulad", "italic green"))
}
