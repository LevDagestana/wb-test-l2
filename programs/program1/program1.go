package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func timeNowNTP() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while getting time from ntp: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Current time: %v\n", time)
	}
}

func main() {
	timeNowNTP()
}
