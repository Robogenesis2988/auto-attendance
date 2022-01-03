package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/OutboundSpade/auto-attendance/sheets"
	"github.com/joho/godotenv"
)

//go:embed credentials.json
var creds []byte

func main() {
	log.SetFlags(log.Lshortfile)
	loadEnv()
	service := sheets.Auth(creds)
	sheetId := os.Getenv("SHEET_ID")
	log.Println("Successfully Authenticated!")
	data, err := sheets.GetSheetData(service, sheetId, "Form Responses 1!A1:B")
	if err != nil {
		log.Fatalf("there was a problem retrieving sheet data - CUSTOM:\n%v", err)
	}

	log.Printf("Data: %v", data)
	for i := 1; i < 20; i++ {
		data := []interface{}{i, i + 1}
		log.Printf("Appending data: %v, %v", data...)
		_, err = sheets.AppendSheet(service, sheetId, "Form Responses 1!A1:B1", [][]interface{}{data})
		if err != nil {
			log.Fatalf("problem appending:%v", err)
		}
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func must(err error, msg string) {

}
