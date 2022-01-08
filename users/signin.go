package users

import (
	"fmt"
	"log"

	"github.com/OutboundSpade/auto-attendance/sheets"
	"github.com/OutboundSpade/auto-attendance/timestamp"
)

func SignIn(uid string, usersSheet *sheets.GoogleSheet, loginSheet *sheets.GoogleSheet) error {
	user := GetUserById(uid, usersSheet)
	if len(user) == 0 {
		return fmt.Errorf("user with id %s doesn't exist! make sure the user is registered", uid)
	}
	log.Printf("Found user: %s", user)
	if UserHasSignedInToday(user, loginSheet) {
		return fmt.Errorf("user %s has already signed in today", user)
	}
	data := [][]interface{}{{timestamp.GetTimeStamp(), user}}
	_, err := sheets.AppendSheet(loginSheet, data)
	if err != nil {
		return fmt.Errorf("there was a problem appending to the sheet with data: %v", data)
	}

	log.Printf("Successfully signed in: %s", user)
	return nil
}
