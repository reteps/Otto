package main

import (
	"github.com/reteps/Otto/ottolib"
	"os"
	"os/exec"
	"strings"
)

func send(message, chatid string) {
	exec.Command(ottolib.SendLocation(), message, chatid).Run()
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
