package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func send(message, chatid string) {
	mybuddy := fmt.Sprintf("set mybuddy to a reference to text chat id \"%s\"", chatid)
	send := fmt.Sprintf("send \"%s\" to mybuddy", message)
	exec.Command("/usr/bin/osascript", "-e", "tell application \"Messages\"", "-e", mybuddy, "-e", send, "-e", "end tell").Run()
}
func testsend(message, chatid string) {
	fmt.Println(message)
}

func readandparsesettings(location string) Results {

	file, err := ioutil.ReadFile(location)

	if err != nil {
		panic(err)
	}
	Data := Results{}
	err = json.Unmarshal(file, &Data)
	if err != nil {
		newlocation := fmt.Sprintf("%sbackup.json", location[:len(location)-5])
		//try backup
		newfile, err := ioutil.ReadFile(newlocation)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(newfile, &Data)
		if err != nil {
			panic(err)
		}
	}
	return Data
}
func writesettings(location string, Data Results) error {
	jsondata, err := json.MarshalIndent(Data, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(location, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}

func checkandwriteallowed(from, chatid string) bool {
	allowedtorun := true
	if from == Data.Chat.Lastperson {
		Data.Chat.Lastamount += 1
		if Data.Chat.Lastamount > 5 {
			allowedtorun = false
		}
		if Data.Chat.Lastamount == 5 {
			send(Data.Maxmessage, chatid)
		}

	} else {
		Data.Chat.Lastperson = from
		Data.Chat.Lastamount = 1
	}
	return allowedtorun
}

var Data Results

func main() {
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid, settingslocation := fulltext[0], fulltext[1], fulltext[2], fulltext[3]
	Data = readandparsesettings(settingslocation)
	ottomessage := false
	if len(message) >= 4 {
		if strings.ToLower(message[:4]) == "otto" {
			ottomessage = true
			//check if allowed
			//send correct text
			allowedtorun := checkandwriteallowed(from, chatid)
			if allowedtorun {
				phrase := message[4:]
				hasntBeenCalled := true
				for key, value := range ottomap {
					if len(phrase) > len(key) {
						if phrase[1:len(key)+1] == key {
							hasntBeenCalled = false
							var result string

							switch value.(type) {
							case string:
								result = value.(string)
							case func(string, string) string:
								result = value.(func(string, string) string)(phrase[len(key)+1:], Data.Chat.Lasttextperson)
							case func(string) string:
								result = value.(func(string) string)(phrase[len(key)+1:])
							case func() string:
								result = value.(func() string)()
							default:
								result = "This function was not created properly."
							}
							testsend(result, chatid)
							break
						}
					}
				}
				if hasntBeenCalled {
					testsend(Data.Errormessage, chatid)
				}
			}
			err := writesettings(settingslocation, Data)
			if err != nil {
				panic(err)
			}

		}
	}
	if ottomessage != true {
		Data.Chat.Lasttext = message
		Data.Chat.Lasttextperson = from
		err := writesettings(settingslocation, Data)
		if err != nil {
			panic(err)
		}
	}
}
