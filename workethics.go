package main

import (
	"fmt"
	"time"
        "math/rand"
	"github.com/dazeus/dazeus-go"
)

func WorkEthics(network, channel string, dz *dazeus.DaZeus) {
ethicLoop:
	for {
		now := time.Now()
		today0900 := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 200, now.Location())
		today1200 := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 100, now.Location())
		today1700 := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 200, now.Location())
		today1730 := time.Date(now.Year(), now.Month(), now.Day(), 17, 30, 0, 200, now.Location())
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 1, 0, now.Location())
		nextnextDay := time.Date(now.Year(), now.Month(), now.Day()+2, 0, 0, 1, 0, now.Location())
		weekday := now.Weekday() // Sunday = 0, ...
		
		// choose a random time between 9:00 and 16:00 for the coffee time notification
		rand.Seed(time.Now().UTC().UnixNano())
		randHour := rand.Intn(15 - 9) + 9
		randMin := rand.Intn(60)
		todayCoffeeRandom := time.Date(now.Year(), now.Month(), now.Day(), randHour, randMin, 0, 200, now.Location())

		if weekday >= time.Monday && weekday < time.Saturday {
			fmt.Println("Happy weekday!")
			if now.Before(today0900) {
				fmt.Println("Waiting until 09:00:00")
				fmt.Println("Setting a timer: " + time.Until(today0900).String())
				select {
				case <-time.After(time.Until(today0900)):
					dz.Message(network, channel, "Het is weer tijd voor noeste arbeid!")
					continue ethicLoop
				}
			}
			if now.Before(todayCoffeeRandom) {
				fmt.Println("Waiting until random coffee time")
				fmt.Println("Setting a timer: " + time.Until(todayCoffeeRandom).String())
				select {
				case <-time.After(time.Until(todayCoffeeRandom)):
					dz.Message(network, channel, "Hebben jullie ook zo'n zin in koffie?")
					continue ethicLoop
				}
			}
			if now.Before(today1200) {
				fmt.Println("Waiting until 12:00:00")
				fmt.Println("Setting a timer: " + time.Until(today1200).String())
				select {
				case <-time.After(time.Until(today1200)):
					dz.Message(network, channel, "Is het al lunchtijd?")
					continue ethicLoop
				}
			}
			if now.Before(today1700) {
				fmt.Println("Waiting until 17:00:00")
				fmt.Println("Setting a timer: " + time.Until(today1700).String())
				select {
				case <-time.After(time.Until(today1700)):
					dz.Message(network, channel, "Het is weer gedaan met de pret. Op naar huis!")
					continue ethicLoop
				}
			}
			if now.Before(today1730) {
				fmt.Println("Waiting until 17:30:00")
				fmt.Println("Setting a timer: " + time.Until(today1730).String())
				select {
				case <-time.After(time.Until(today1730)):
					dz.Message(network, channel, "]17:30")
					continue ethicLoop
				}
			}
			// We're done for today, set a timer for 1 second after midnight
			fmt.Println("It's time for bed. Good night.")
			fmt.Println("Setting a timer: " + time.Until(nextDay).String())
			select {
			case <-time.After(time.Until(nextDay)):
				continue ethicLoop
			}
		}
		// Suppose we start this program on a Saturday, set the timer for next next day (Monday)
		if weekday == time.Saturday {
			// Set a timer for next Monday 00:00:01
			fmt.Println("Today is weekend, namely " + weekday.String() + ", go relax!")
			fmt.Println("Setting a timer: " + time.Until(nextnextDay).String())
			select {
			case <-time.After(time.Until(nextnextDay)):
				continue ethicLoop
			}
		}
		// Suppose we start this program on a Sunday, set the timer for next day
		if weekday == time.Sunday {
			// Set a timer for next Monday 00:00:01
			fmt.Println("Today is weekend, namely " + weekday.String() + ", go relax!")
			fmt.Println("Setting a timer: " + time.Until(nextDay).String())
			select {
			case <-time.After(time.Until(nextDay)):
				continue ethicLoop
			}
		}
	}
}
