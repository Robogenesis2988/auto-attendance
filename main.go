package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/OutboundSpade/auto-attendance/serial"
	"github.com/OutboundSpade/auto-attendance/sheets"
	"github.com/joho/godotenv"
)

//go:embed credentials.json
var creds []byte

func main() {
	// log.SetFlags(log.Lshortfile)

	var portArg string
	flag.StringVar(&portArg, "port", "", "The name/path of the port")
	flag.Parse()
	if portArg == "" {
		fmt.Println("-port flag required!")
		allPorts := serial.GetUSBPorts()
		if len(allPorts) == 0 {
			fmt.Println("No USB ports found! Please make sure the device is plugged in")
		} else {
			fmt.Println("Detected ports:")
			for _, p := range allPorts {
				fmt.Printf("\t%s\n", p)
			}
		}
		os.Exit(1)
		return
	}

	loadEnv()
	sheetId := os.Getenv("SHEET_ID")
	if sheetId == "" {
		log.Fatal("the environment variable 'SHEET_ID' must be set")
	}
	service := sheets.Auth(creds)
	log.Println("Successfully Authenticated via Google!")

	port := serial.Open(portArg, 9600)
	serialData := make(chan string)
	go serial.ListenForData(&port, serialData)

	for {
		data := <-serialData
		log.Printf("Recieved SerialData: %v", data)
	}

	data, err := sheets.GetSheetData(service, sheetId, "Form Responses 1!A2:B")
	must(err, "there was a problem retrieving sheet data - CUSTOM")
	log.Println("Data:")
	for _, row := range data {
		log.Printf("%v", row)
	}

	// for i := 1; i < 20; i++ {
	// 	data := []interface{}{i, i + 1}
	// 	log.Printf("Appending data: %v, %v", data...)
	// 	_, err = sheets.AppendSheet(service, sheetId, "Form Responses 1!A1:B1", [][]interface{}{data})
	// 	must(err, "problem appending")
	// }
}

func loadEnv() {
	should(godotenv.Load(), "error loading .env file")
}

func should(err error, msg string) {
	if err := recover(); err != nil {
		log.Fatalf("%v:\n%v", msg, err)
	}
}
func must(err error, msg string) {
	if err != nil {
		log.Fatalf("%v:\n%v", msg, err)
	}
}
