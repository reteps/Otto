package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

//Peter Stenger @reteps 7-6-17
//if you modify or use my code, credit me

func date(message, from string) string {
	t := time.Now()
	format := fmt.Sprintf("%s, %s %s, %s", t.Weekday(), t.Month(), t.Day(), t.Year())
	return format
}
func send(message, chatid string) {
	exec.Command("/Users/Peter/repos/Otto/SendText.applescript", message, chatid).Run()
}
func main() {
	keywords := map[string]string{"hello": "hello there!", "version": "I am currently version 1.0", "date": "function"}
	funcMap := map[string]func(string, string) string{"date": date}
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid := fulltext[0], fulltext[1], fulltext[2]
	if len(message) >= 4 {
		section := strings.ToLower(message[:4])
		if section == "otto" {
			//otto is being invoked
			phrase := message[4:]
			hasBeenCalled := false
			for key, value := range keywords {
				if strings.Contains(phrase, key) {
					//do whatever that word maps to
					hasBeenCalled = true
					var result string
					if value == "function" {
						//assumes date title matches with keyword
						result = funcMap[key](message, from)
					} else {
						result = value
					}
					send(result, chatid)
					break
				}
			}
			if hasBeenCalled == false {
				send("Otto is at your service", chatid)
			}
		}
	}
}
