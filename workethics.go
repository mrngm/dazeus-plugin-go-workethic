package main

import (
	"fmt"
	"time"

	"github.com/dazeus/dazeus-go"
)

type timeMessage struct {
	timedate time.Time
	message  string
}

func WorkEthics(network, channel string, dz *dazeus.DaZeus) {
	var minTimeID string
	var smallestDuration time.Duration
	times := make(map[string]timeMessage)
ethicLoop:
	for {
		now := time.Now()
		times["start"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 200, now.Location()),
			message:  "Het is weer tijd voor noeste arbeid!",
		}
		times["lunch"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 100, now.Location()),
			message:  "Is het al lunchtijd?",
		}
		times["stop"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 200, now.Location()),
			message:  "Het is weer gedaan met de pret. Op naar huis!",
		}
		times["stophx"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day(), 17, 30, 0, 200, now.Location()),
			message:  "]17:30",
		}
		times["nextday"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 1, 0, now.Location()),
		}
		times["nextnextday"] = timeMessage{
			timedate: time.Date(now.Year(), now.Month(), now.Day()+2, 0, 0, 1, 0, now.Location()),
		}
		weekday := now.Weekday() // Sunday = 0, ...

		if weekday >= time.Monday && weekday < time.Saturday {
			fmt.Println("Happy weekday!")
			smallestDuration, _ = time.ParseDuration("1337h")
			for id, t := range times {
				// Determine what time is closest to now
				if time.Until(t.timedate) >= 0 && time.Until(t.timedate) < smallestDuration {
					smallestDuration = time.Until(t.timedate)
					minTimeID = id
				}
			}

			// Start a timer
			fmt.Printf("Setting a timer for '%s' (%s): %s\n", minTimeID, times[minTimeID].message, time.Until(times[minTimeID].timedate).String())
			select {
			case <-time.After(time.Until(times[minTimeID].timedate)):
				if times[minTimeID].message != "" {
					dz.Message(network, channel, times[minTimeID].message)
				}
				continue ethicLoop
			}
		}
		// Suppose we start this program on a Saturday, set the timer for next next day (Monday)
		if weekday == time.Saturday {
			// Set a timer for next Monday 00:00:01
			fmt.Println("Today is weekend, namely " + weekday.String() + ", go relax!")
			fmt.Println("Setting a timer: " + time.Until(times["nextnextday"].timedate).String())
			select {
			case <-time.After(time.Until(times["nextnextday"].timedate)):
				continue ethicLoop
			}
		}
		// Suppose we start this program on a Sunday, set the timer for next day
		if weekday == time.Sunday {
			// Set a timer for next Monday 00:00:01
			fmt.Println("Today is weekend, namely " + weekday.String() + ", go relax!")
			fmt.Println("Setting a timer: " + time.Until(times["nextday"].timedate).String())
			select {
			case <-time.After(time.Until(times["nextday"].timedate)):
				continue ethicLoop
			}
		}
	}
}
