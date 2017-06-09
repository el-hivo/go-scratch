package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")

	today := time.Now()
	dubaiTz, err := time.LoadLocation("Asia/Dubai")

	if err != nil {
		fmt.Println("Error loading zone, err = ", err.Error())
		return
	}

	laTz, err := time.LoadLocation("America/Los_Angeles")

	if err != nil {
		fmt.Println("Error loading zone, err = ", err.Error())
		return
	}

	todayInDubai := today.In(dubaiTz)
	fmt.Println("Local: ", today)
	fmt.Println("Local: ", today.Weekday())

	fmt.Println("Dubai :", todayInDubai)
	fmt.Println("Dubai: ", todayInDubai.Weekday())

	//let's try with UTC
	dateStr := "2017-06-14T22:00:00.000Z"
	laDate, err := time.ParseInLocation(time.RFC3339, dateStr, laTz)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	laDate = laDate.In(laTz)

	fmt.Println("In LA: ", laDate)
	fmt.Println("In LA: ", laDate.Weekday())

	d, _ := time.ParseInLocation("2006-01-02", "2010-06-28", dubaiTz)
	l, _ := time.ParseInLocation("2006-01-02", "2010-06-28", laTz)

	fmt.Println("Dubai: ", d, " is a ", d.Weekday(), " in UTC ", d.UTC())
	fmt.Println("Los Angeles: ", l, " is a ", l.Weekday(), " in UTC ", l.UTC())

	apiaTz, err := time.LoadLocation("Pacific/Fiji")

	if err != nil {
		fmt.Println("Weird not recognized")
	}

	at, _ := time.ParseInLocation("2006-01-02", "2010-06-28", apiaTz)
	fmt.Println("Apia: ", at, " is a ", at.Weekday(), " in UTC ", at.UTC())

	//Pacific/Tahiti
	tahitiTz, err := time.LoadLocation("Pacific/Tahiti")

	if err != nil {
		fmt.Println("Weird not recognized")
	}

	tt, _ := time.ParseInLocation("2006-01-02", "2010-06-28", tahitiTz)
	fmt.Println("Tahiti: ", tt, " is a ", tt.Weekday(), " in UTC ", tt.UTC())
}
