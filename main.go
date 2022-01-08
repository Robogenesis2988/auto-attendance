package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/OutboundSpade/auto-attendance/serial"
	"github.com/OutboundSpade/auto-attendance/sheets"
	"github.com/OutboundSpade/auto-attendance/users"
	"github.com/joho/godotenv"
)

//go:embed credentials.json
var creds []byte

func main() {
	// log.SetFlags(log.Lshortfile)
	// scanSound, _ := sound.InitSound("scan.wav")

	var portArg string
	var verify bool
	var register bool
	flag.StringVar(&portArg, "port", "", "The name/path of the port")
	flag.BoolVar(&verify, "verify", false, "Used to find out which person a tag belongs to")
	flag.BoolVar(&register, "register", false, "Used to register a tag with a name")
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
	usersSheet := sheets.GoogleSheet{
		Service:    service,
		Id:         sheetId,
		ValueRange: "users!A2:B",
	}
	loginSheet := sheets.GoogleSheet{
		Service:    service,
		Id:         sheetId,
		ValueRange: "Form Responses 1!A2:B",
	}

	port := serial.Open(portArg, 9600)
	log.Printf("Listening for data on port %s", portArg)

	// sound.PlayStartupSound()

	for {
		serialData := make(chan string)
		go serial.ListenForData(&port, serialData)
		uid := <-serialData //Listen for serial data
		// sound.PlayScanSound()
		log.Printf("Recieved SerialData: %v", uid)
		err := users.SignIn(uid, &usersSheet, &loginSheet)
		if err != nil {
			log.Print(err)
		}
	}

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
