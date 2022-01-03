package sheets

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

// Request a token from the web, then returns the retrieved token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)

// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		log.Fatalf("Unable to read authorization code: %v", err)
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	return tok
// }

// Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Unable to cache oauth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

func Auth(creds []byte) *sheets.Service {
	ctx := context.Background()
	b := creds
	// b, err := ioutil.ReadFile(credsPath)
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }
	// If modifying these scopes, delete your previously saved token.json.

	client, err := google.CredentialsFromJSON(ctx, b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	// client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithCredentials(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv
}

func GetSheetData(srv *sheets.Service, sheetId string, valueRange string) ([][]string, error) {
	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := sheetId
	readRange := valueRange
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		// log.Println("No data found.")
		return nil, nil
	} else {
		var data [][]string
		for _, row := range resp.Values {
			// Create an [][]string without empty cells
			if len(row) != 0 {
				data = append(data, []string{})
				realIndex := len(data) - 1
				for _, item := range row {
					data[realIndex] = append(data[realIndex], item.(string))
					// fmt.Printf("Item: %s, \n", item)
				}
			}
			// fmt.Println()
		}
		return data, nil

	}
}

func AppendSheet(srv *sheets.Service, sheetId string, headerRange string, values [][]interface{}) (res *sheets.AppendValuesResponse, err error) {
	vals := &sheets.ValueRange{MajorDimension: `ROWS`, Values: values} //, Range: headerRange
	resp, err := srv.Spreadsheets.Values.Append(sheetId, headerRange, vals).ValueInputOption("USER_ENTERED").Do()
	return resp, err
}
