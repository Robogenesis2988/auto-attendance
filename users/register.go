package users

import (
	"fmt"
	"log"

	"github.com/OutboundSpade/auto-attendance/sheets"
)

func PromptForUserRegistration(name string) {
	if len(name) == 0 {
		fmt.Print("First & Last Name:")
		fmt.Scanln(&name)
	}
	log.Println(name)
}

func RegisterUser(fullname string, uid string, usersSheet *sheets.GoogleSheet) error {
	if isIdUsed(uid, usersSheet) {
		return fmt.Errorf("code %v is already in use", uid)
	}
	return nil
}
