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
	newmessage := "Commands include:\n" + strings.Join(keys, ", ")
	return newmessage
}
func Randint(low, high int) string {
	result := strconv.Itoa(rand.Intn((high+1)-low) + low)
	return result

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
	return Randint(low, high)
}
func Say(message, from string) string {
	if message[4:] == "" {
		return "say what?"
	}
	return message[5:]
}
func Roll(message, from string) string {
	if message[5:] == "" {
		return "roll what? ex. 2d20"
	}
	sections := strings.Split(message[6:], "d")
	dice, err := strconv.Atoi(sections[0])
	if err != nil {
		return "invalid dice amount"
	}
	num, err := strconv.Atoi(sections[1])
	if err != nil {
		return "invalid high roll"
	}
	if dice > 100 || num > 100 {
		return "max number is 100"
	}
	var result []string
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < dice; i++ {
		result = append(result, Randint(1, num))
	}
	return strings.Join(result, ",")
}
func randbool() bool {
	num := rand.Float64()
	if num > 0.5 {
		return true
	} else {
		return false
	}
}
func Mock(message, from string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	ftext := ""
	for _, v := range Data.Lasttext {
		sv := string(v)
		mybool := randbool()
		if mybool == true {
			ftext += strings.ToUpper(sv)
		} else {
			ftext += strings.ToLower(sv)
		}
	}
	return ftext
}
func Flip(message, from string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	state := randbool()
	if state == true {
		return "heads"
	} else {
		return "tails"
	}
}
func Magic(message, from string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	num, _ := strconv.Atoi(Randint(0, len(Data.Phrases)-1))
	return Data.Phrases[num]
}

var (
	funcmap = map[string]func(string, string) string{"date": Date, "help": Help, "random": Random, "say": Say, "roll": Roll, "mock": Mock,
		"flip": Flip, "magic": Magic, "will": Magic}
	keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.1beta",
		"date": "FUNCTION", "help": "FUNCTION", "random": "FUNCTION", "say": "FUNCTION",
		"what": "I am a imessage virtual assistant that runs when Peter's computer is on. Type 'otto help' to see all the commands I can do.",
		"roll": "FUNCTION", "mock": "FUNCTION", "thanks": "you're welcome", "flip": "FUNCTION", "magic": "FUNCTION", "will": "FUNCTION"}
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
	Lastperson     string   `json:"lastperson"`
	Lastamount     int      `json:"lastamount"`
	Lasttext       string   `json:"lasttext"`
	Lasttextperson string   `json:"lasttextperson"`
	Errormessage   string   `json:"errormessage"`
	Phrases        []string `json:"8ballwords"`
}

func readandparsesettings(location string) Results {

	file, err := ioutil.ReadFile(location)

	if err != nil {
		panic(err)
	}
	Data := Results{}
	err = json.Unmarshal(file, &Data)
	if err != nil {
		panic(err)
	}
	return Data
}
func writesettings(location string, Data Results) error {
	jsondata, err := json.Marshal(Data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(location, jsondata, 0644)
	if err != nil {
		return err
	}
	return nil
}

var Data Results

func main() {
	fulltext := strings.Split(os.Args[1:][0], "|~|")
	message, from, chatid, settingslocation := fulltext[0], fulltext[1], fulltext[2], fulltext[3]
	Data = readandparsesettings(settingslocation)
	ottomessage := false
	if len(message) >= 4 {
		section := strings.ToLower(message[:4])
		if section == "otto" {
			ottomessage = true
			//check if allowed
			allowedtorun := true
			if from == Data.Lastperson {
				Data.Lastamount += 1
				if Data.Lastamount > 5 {
					allowedtorun = false
				}
				if Data.Lastamount == 5 {
					send("You have reached your 5 consecutive text limit", chatid)
				}

			} else {
				Data.Lastperson = from
				Data.Lastamount = 1
			}
			//send correct text
			if allowedtorun {
				phrase := message[4:]
				hasntBeenCalled := true
				for key, value := range keywords {
					if len(phrase) > len(key) {
						if phrase[1:len(key)+1] == key {
							hasntBeenCalled = false
							var result string
							if value == "FUNCTION" {
								result = funcmap[key](phrase, Data.Lasttextperson)
							} else {
								result = value
							}
							send(result, chatid)
							break
						}
					}
				}
				if hasntBeenCalled {
					send(Data.Errormessage, chatid)
				}
			}
			err := writesettings(settingslocation, Data)
			if err != nil {
				panic(err)
			}

		}
	}
	if ottomessage != true {
		Data.Lasttext = message
		Data.Lasttextperson = from
		err := writesettings(settingslocation, Data)
		if err != nil {
			panic(err)
		}
	}
}
