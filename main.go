package main

import (
	"fmt"
	"github.com/reteps/Otto/ottolib"
	"os"
	"os/exec"
	"strings"
)

func send(message, chatid string) {
	command := fmt.Sprintf("osascript -e 'tell application \"Messages\"' -e 'set mybuddy to a reference to text chat id \"%s\"' -e 'send \"%s\" to mybuddy' -e 'end tell'", chatid, message)
	exec.Command(command).Run()
}
func main() {
	keywords := ottolib.Keywords()
	funcMap := ottolib.FuncMap()
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid := fulltext[0], fulltext[1], fulltext[2]

	if len(message) >= 4 {
		section := strings.ToLower(message[:4])
		if section == "otto" {
			phrase := message[4:]
			hasntBeenCalled := true
			for key, value := range keywords {
				if strings.Contains(phrase, key) {
					hasntBeenCalled = false
					var result string
					if value == "FUNCTION" {
						result = funcMap[key](message, from)
					} else {
						result = value
					}
					send(result, chatid)
					break
				}
			}
			if hasntBeenCalled {
				send(ottolib.Errormessage(), chatid)
			}
		}
	}
}
