package users

import (
	"log"
	"strings"

	"github.com/OutboundSpade/auto-attendance/sheets"
	"github.com/OutboundSpade/auto-attendance/timestamp"
)

func GetUserById(uid string, s *sheets.GoogleSheet) string {
	data, err := sheets.GetSheetData(s)
	mustSheet(err)

	var name string
	for _, row := range data {
		if len(row) == 2 && row[0] == uid {
			name = row[1]
		}
	}
	return name
}
func UserHasSignedInToday(fullname string, loginSheet *sheets.GoogleSheet) bool {
	data, err := sheets.GetSheetData(loginSheet)
	mustSheet(err)
	date := timestamp.GetDate()
	for _, row := range data {
		if len(row) == 2 && strings.Split(row[0], " ")[0] == date && row[1] == fullname {
			return true
		}
	}
	return false
}

func isIdUsed(uid string, s *sheets.GoogleSheet) bool {
	data, err := sheets.GetSheetData(s)
	mustSheet(err)
	for _, row := range data {
		if row[0] == uid {
			return true
		}
	}
	return false
}

func mustSheet(err error) {
	if err != nil {
		log.Fatalf("there was a problem getting sheet data: %v", err)
	}
}
