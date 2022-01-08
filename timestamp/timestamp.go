package timestamp

import (
	"fmt"
	"time"
)

func GetTimeStamp() string {
	now := time.Now()
	year, m, day := now.Date()
	month := int(m)
	hr, min, sec := now.Clock()

	return fmt.Sprintf("%d/%d/%d %d:%d:%d", month, day, year, hr, min, sec)
}

func GetDate() string {
	now := time.Now()
	year, m, day := now.Date()
	month := int(m)
	return fmt.Sprintf("%d/%d/%d", month, day, year)
}
