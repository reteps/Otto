package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
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
	values := strings.Split(message[7:], " ")
	fmt.Println(values)
	if len(values) != 3 { //space
		return "usage:random low high"
	}
	low, err := strconv.Atoi(values[1])
	if err != nil {
		return "invalid low number"
	}
	high, err := strconv.Atoi(values[2])
	if err != nil {
		return "invalid high number"
	}
	if high-low <= 0 {
		return "high must be bigger then low"
	}
	rand.Seed(time.Now().UTC().UnixNano())
	result := strconv.Itoa(rand.Intn((high+1)-low) + low)
	return result
}
func Say(message, from string) string {
	if message[4:] == "" {
		return "say what?"
	}
	return message[5:]
}

var (
	funcmap  = map[string]func(string, string) string{"date": Date, "help": Help, "random": Random, "say": Say}
	keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.1beta",
		"date": "FUNCTION", "help": "FUNCTION", "random": "FUNCTION", "say": "FUNCTION", "what": "I am a imessage virtual assistant that runs when Peter's computer is on. Type 'otto help' to see all the commands I can do."}
)

//--------------DO NOT MODIFY------------------------//

func send(message, chatid string) {
	mybuddy := fmt.Sprintf("set mybuddy to a reference to text chat id \"%s\"", chatid)
	send := fmt.Sprintf("send \"%s\" to mybuddy", message)
	exec.Command("/usr/bin/osascript", "-e", "tell application \"Messages\"", "-e", mybuddy, "-e", send, "-e", "end tell").Run()
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
				data.Lastamount = 0
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
							result = funcmap[key](phrase, from)
						} else {
							result = value
						}
						send(result, chatid)
						break
					}
				}
				if hasntBeenCalled {
					send(data.Errormessage, chatid)
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
