package helpers

import (
	"time"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CurrDate() int64 {

	curr := time.Now().Format("2006-01-02")

	resp, err := time.Parse("2006-01-02", curr)
	CheckErr(err)
	return resp.Unix()
}

func ToUnix(date string) int64 {

	resp, err := time.Parse("2006-01-02", date)
	CheckErr(err)
	return resp.Unix()
}
