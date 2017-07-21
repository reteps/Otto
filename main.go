package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

//--------------FUNCTIONS----------------------//
func Date(message, from string) string {
	t := time.Now()
	format := fmt.Sprintf("Today is %s, %s %d, %d", t.Weekday(), t.Month(), t.Day(), t.Year())
	return format
}
func Help(message, from string) string {
	keys := make([]string, len(keywords))

	i := 0
	for k := range keywords {
		keys[i] = k
		i++
	}
	newmessage := "Commands include:" + strings.Join(keys, ",")
	return newmessage
}
func Random(message, from string) string {
	return "work in progress"
}

var (
	funcmap  = map[string]func(string, string) string{"date": Date, "help": Help, "random": Random}
	keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.0",
		"date": "FUNCTION", "help": "FUNCTION", "random": "FUNCTION"}
)

//--------------DO NOT MODIFY------------------------//

func send(message, chatid string) {
	command := fmt.Sprintf("osascript -e 'tell application \"Messages\"' -e 'set mybuddy to a reference to text chat id \"%s\"' -e 'send \"%s\" to mybuddy' -e 'end tell'", chatid, message)
	exec.Command(command).Run()
}
func testsend(message, chatid string) {
	fmt.Println(message)
}

type Results struct {
	Lastperson   string `json:"lastperson"`
	Lastamount   int    `json:"lastamount"`
	Lasttext     string `json:"lasttext"`
	Errormessage string `json:"errormessage"`
}

func readandparsesettings(location string) Results {

	file, err := ioutil.ReadFile(location)

	if err != nil {
		panic(err)
	}
	data := Results{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	return data
}
func writesettings(location string, data Results) error {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(location, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid, settingslocation := fulltext[0], fulltext[1], fulltext[2], fulltext[3]
	data := readandparsesettings(settingslocation)
	ottomessage := false
	if len(message) >= 4 {
		section := strings.ToLower(message[:4])
		if section == "otto" {
			ottomessage = true
			//check if allowed
			allowedtorun := true
			if from == data.Lastperson {
				data.Lastamount += 1
				if data.Lastamount >= 5 {
					allowedtorun = false
				}

			} else {
				data.Lastperson = from
				data.Lastamount = 1
			}
			//send correct text
			if allowedtorun {
				phrase := message[4:]
				hasntBeenCalled := true
				for key, value := range keywords {
					if strings.Contains(phrase, key) {
						hasntBeenCalled = false
						var result string
						if value == "FUNCTION" {
							result = funcmap[key](message, from)
						} else {
							result = value
						}
						testsend(result, chatid)
						//send(result, chatid)
						break
					}
				}
				if hasntBeenCalled {
					testsend(data.Errormessage, chatid)
				}
			}
			err := writesettings(settingslocation, data)
			if err != nil {
				panic(err)
			}

		}
	}
	if ottomessage != true {
		data.Lasttext = message
		err := writesettings(settingslocation, data)
		if err != nil {
			panic(err)
		}
	}
}
