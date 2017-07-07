package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Peter Stenger @reteps 7-6-17
//if you modify or use my code, credit me

func date(message, from string) {
	return fmt.Sprintf("%s, %s %s, %s", time.stdLongWeekDay, time.stdLongMonth, stdDay, stdLongYear)
}
func send(message, chatid string) {
	exec("./SendText.applescript", message, chatid)
}
func main() {
	keywords := map[string]string{"hello": "hello there!", "version": "I am currently version 1.0", "date": "function"}
	message, from, chatid := strings.split(os.Args[1:], "|~|")
	if len(message) >= 4 {
		section = strings.toLower(message[:4])
		if section == "otto" {
			//otto is being invoked
			phrase = message[4:]
			for key, value := range keywords {
				if strings.Contains(phrase, key) {
					//do whatever that word maps to
					if value == "function" {
						//assumes date title matches with keyword
						result := key(message, from)
					} else {
						result := action
					}
					send(result, chatid)
				}
			}
		}
	}
}
