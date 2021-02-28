package helpers

import (
	"log"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CurrDate() int64 {

	curr := time.Now().Format("2006-01-02")

	resp, err := time.Parse("2006-01-02", curr)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Unix()
}

func ToUnix(date string) int64 {

	resp, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Unix()
}
