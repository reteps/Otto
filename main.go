package main

import (
	"encoding/json"
	"fmt"
	"github.com/alfredxing/calc/compute"
	"io/ioutil"
	"math/rand"
	"net/http"
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
	if message == "" {
		return "usage:random low high"
	}
	values := strings.Split(message, " ")
	low, err := strconv.Atoi(values[0])
	if err != nil {
		return "invalid low number"
	}
	high, err := strconv.Atoi(values[1])
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
	if message == "" {
		return "say what?"
	}
	return message[1:]
}
func Roll(message, from string) string {
	if message == "" {
		return "roll what? ex. 2d20"
	}
	sections := strings.Split(message, "d")
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
	for _, v := range Data.Chat.Lasttext {
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
	num, _ := strconv.Atoi(Randint(0, len(Data.Eightball.Phrases)-1))
	return Data.Eightball.Phrases[num]
}
func Weather(message, from string) string {
	var location string
	if message == "" {
		location = Data.Weather.Default
	} else {
		location = strings.ToLower(message)
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", location, Data.Weather.Apikey)
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	type WeatherDecoder struct {
		Main          map[string]interface{}   `json:"main"`
		Name          string                   `json:"name"`
		Weather       []map[string]interface{} `json:"weather"` //description
		Coord         map[string]float64       `json:"coord"`
		Base          string                   `json:"base"`
		Visibility    int                      `json:"visibility"`
		Wind          map[string]float64       `json:"wind"`
		Clouds        map[string]int           `json:"clouds"`
		DateInSeconds int                      `json:"dt"`
		Sys           map[string]float64       `json:"sys"`
		Id            int                      `json:"id"`
		Cod           int                      `json:"cod"`
	}

	weather := &WeatherDecoder{}
	json.NewDecoder(resp.Body).Decode(weather)
	response := fmt.Sprintf("Right now in %s, it is %.2f degrees. The weather is %s and there is %.0f%% humidity.",
		weather.Name, weather.Main["temp"], weather.Weather[0]["main"], weather.Main["humidity"])
	return response
}

func Calc(message, from string) string {
	if message == "" {
		return "Calculate what?"
	}
	res, err := compute.Evaluate(message)
	if err != nil {
		return err.Error()
	}
	return strconv.FormatFloat(res, 'f', 4, 64)
}

var (
	funcmap = map[string]func(string, string) string{"date": Date, "help": Help, "random": Random, "say": Say, "roll": Roll, "mock": Mock,
		"flip": Flip, "magic": Magic, "will": Magic, "weather": Weather, "calc": Calc}
	keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.1beta",
		"date": "FUNCTION", "help": "FUNCTION", "random": "FUNCTION", "say": "FUNCTION",
		"what": "I am a imessage virtual assistant that runs when Peter's computer is on. Type 'otto help' to see all the commands I can do.",
		"roll": "FUNCTION", "mock": "FUNCTION", "thanks": "you're welcome", "flip": "FUNCTION", "magic": "FUNCTION", "will": "FUNCTION",
		"hi": "hi there!", "weather": "FUNCTION", "calc": "FUNCTION"}
)

type WeatherSettings struct {
	Default string `json:"default"`
	Apikey  string `json:"apikey"`
}

type EightballSettings struct {
	Phrases    []string            `json:"phrases"`
	Eastereggs map[string][]string `json:"eastereggs"`
}
type ChatSettings struct {
	Lastperson     string `json:"lastperson"`
	Lastamount     int    `json:"lastamount"`
	Lasttext       string `json:"lasttext"`
	Lasttextperson string `json:"lasttextperson"`
}
type Results struct {
	Weather      WeatherSettings   `json:"weather"`
	Chat         ChatSettings      `json:"chat"`
	Errormessage string            `json:"errormessage"`
	Eightball    EightballSettings `json:"eightball"`
}

//--------------DO NOT MODIFY------------------------//

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
			allowedtorun := true
			if from == Data.Chat.Lastperson {
				Data.Chat.Lastamount += 1
				if Data.Chat.Lastamount > 5 {
					allowedtorun = false
				}
				if Data.Chat.Lastamount == 5 {
					send("You have reached your 5 consecutive text limit", chatid)
				}

			} else {
				Data.Chat.Lastperson = from
				Data.Chat.Lastamount = 1
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
								result = funcmap[key](phrase[len(key)+1:], Data.Chat.Lasttextperson)
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
		Data.Chat.Lasttext = message
		Data.Chat.Lasttextperson = from
		err := writesettings(settingslocation, Data)
		if err != nil {
			panic(err)
		}
	}
}
