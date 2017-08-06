package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/alfredxing/calc/compute"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//FUNCTION MAP
var ottomap map[string]interface{}

func init() {
	ottomap = map[string]interface{}{"date": Date,
		"help":    Help,
		"random":  Random,
		"say":     Say,
		"roll":    Roll,
		"mock":    Mock,
		"flip":    Flip,
		"magic":   Magic,
		"will":    Magic,
		"tod":	   ToD,
		"weather": Weather,
		"calc":    Calc,
		"egg":     Egg, //eightball easter egg
		"hello":   "hello there!",
		"version": "I am currently version 1.1beta",
		"what":    "I am a imessage virtual assistant that runs when Peter's computer is on. Type 'otto help' to see all the commands I can do.",
		"hi":      "hi there!",
		"time":    Time, //running off host computer
		"thanks":  "you're welcome",
		"google":  Google,   //gets first span
		"wiki":    Wiki,     //link
		"info":    Wikitext, //intro paragraph
	}
}

//FUNCTIONS
func ToD(message, from string) string {
 	//Chooses a random truth or dare
 	dares := []string{
 		"Lick the floor.",
 		"Dance with no music for 1 minute.",
 		"Break two eggs on your head.",
 		"Do your best impression of a baby being born.",
 		"Put 4 ice cubes down your pants.",
 	}
 	truths := []string{
 		"What are you most self-conscious about?",
 		"What is your deepest darkest fear?",
 		"What is the scariest dream you have ever had?",
 		"What is the stupidest thing you have ever done?",
 		"What is the airspeed velocity of an unladen swallow?",
 	}
 	if message == "truth" {
 		rand.Seed(time.Now().Unix())
 		n := rand.Int() % len(test)
 		tod := fmt.Sprintln(truths[n])
 	}
 	if message == "dare" {
 		rand.Seed(time.Now().Unix())
 		n := rand.Int() % len(test)
 		tod := fmt.Sprintf(dares[n])
 	}
 }
func Wiki(message string) string {
	if message == "" {
		return "search wikipedia for what?"
	}
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&search=%s&limit=1&namespace=0&format=json", strings.Replace(message[1:], " ", "%20", -1))
	resp, err := http.Get(url)
	if err != nil {
		return "read error wikisearch" + err.Error()
	}
	var result1 []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result1)
	if err != nil {
		return "decode error wikisearch " + err.Error()
	}
	defer resp.Body.Close()
	urllist := result1[len(result1)-1].([]interface{})
	if len(urllist) == 0 {
		return "Wikipedia couldn't find that page"
	}
	newurl := urllist[0].(string)
	return newurl
}
func Wikitext(message string) string {
	newurl := Wiki(message)
	if strings.Contains(newurl, "Wikipedia couldn't find that page") || strings.Contains(newurl, "error wikisearch") {
		return newurl
	}
	page := strings.Split(newurl, "/")
	pageurl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&prop=extracts&exintro&explaintext&titles=%s&format=json", strings.Replace(page[len(page)-1], "_", "%20", -1))
	//PART 2
	type Wikipage struct {
		Pageid  int    `json:"pageid"`
		Ns      int    `json:"ns"`
		Title   string `json:"title"`
		Extract string `json:"extract"`
	}
	type Wikidata struct {
		Complete string                         `json:"batchcomplete"`
		Query    map[string]map[string]Wikipage `json:"query"`
	}
	resp, err := http.Get(pageurl)
	if err != nil {
		return "read error wikidata " + err.Error()
	}
	wikidata := &Wikidata{}
	err = json.NewDecoder(resp.Body).Decode(wikidata)
	if err != nil {
		return "decode error wikidata " + err.Error()
	}
	defer resp.Body.Close()
	var result string
	for k, _ := range wikidata.Query {
		for k1, _ := range wikidata.Query[k] {

			result = wikidata.Query[k][k1].Extract
			break
		}
		break
	}
	if len(result) > 500 {
		return result[:500] + "..."
	}
	return result

}

func Google(message string) string {
	if message == "" {
		return "google what? NOTE:this doesn't work perfectly"
	}
	url := "http://www.google.com/search?q=" + strings.Replace(strings.Replace(strings.Replace(message[1:], " ", "|~|", -1), "+", "%2B", -1), "|~|", "+", -1)
	response, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(io.Reader(response.Body))
	if err != nil {
		return err.Error()
	}
	var returntext string
	valid := true
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "days ago") == false && strings.Contains(text, "day ago") == false && text != "" && valid == true {
			valid = false
			returntext = text
		}
	})
	return returntext
}
func Date() string {
	t := time.Now()
	format := fmt.Sprintf("Today is %s, %s %d, %d", t.Weekday(), t.Month(), t.Day(), t.Year())
	return format
}
func Help() string {
	keys := make([]string, len(ottomap))

	i := 0
	for k := range ottomap {
		keys[i] = k
		i++
	}
	newmessage := "Commands include:\n" + strings.Join(keys, ", ")
	return newmessage
}
func Random(message string) string {
	if message == "" {
		return "usage:random low high"
	}
	values := strings.Split(message[1:], " ")
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
	return randint(low, high)
}
func Egg(message string) string {
	if message == "" {
		return "for example:`otto egg peace,hippie=Peace dude`"
	}
	parts := strings.Split(message[1:], "=")
	if len(parts) != 2 {
		return "Invalid easter egg for otto magic."
	}
	triggers := strings.Split(parts[0], ",")
	egg := parts[1]
	Data.Eightball.Eastereggs[egg] = triggers
	return fmt.Sprintf("added easter egg for '%s' that is triggered by %v", egg, triggers)
}
func Say(message string) string {
	if message == "" {
		return "say what?"
	}
	return message[1:]
}
func Roll(message string) string {
	if message == "" {
		return "roll what? ex. 2d20"
	}
	sections := strings.Split(message[1:], "d")
	dice, err := strconv.Atoi(sections[0])
	if err != nil {
		return "invalid dice amount"
	}
	num, err := strconv.Atoi(sections[1])
	if err != nil {
		return "invalid high roll"
	}
	if dice > 100 || num > 100 {
		return "highest number to roll / dice to have is 100"
	}
	var result []string
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < dice; i++ {
		result = append(result, randint(1, num))
	}
	return strings.Join(result, ",")
}
func Mock(message string) string {
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
func Flip(message string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	state := randbool()
	if state == true {
		return "heads"
	} else {
		return "tails"
	}
}
func Magic(message string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	for key, value := range Data.Eightball.Eastereggs {
		for _, keyword := range value {
			if strings.Contains(message, keyword) {
				return key
			}
		}
	}
	//normal, no secrets
	num, _ := strconv.Atoi(randint(0, len(Data.Eightball.Phrases)-1))
	return Data.Eightball.Phrases[num]
}
func Weather(message string) string {
	var location string
	if message == "" {
		location = Data.Weather.Default
	} else {
		location = strings.ToLower(message)
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", location, Data.Weather.Apikey)

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

	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	weather := &WeatherDecoder{}
	json.NewDecoder(resp.Body).Decode(weather)

	response := fmt.Sprintf("Right now in %s, it is %.2f degrees. The weather is %s and there is %.0f%% humidity.",
		weather.Name, weather.Main["temp"], weather.Weather[0]["main"], weather.Main["humidity"])
	return response
}

func Calc(message string) string {
	if message == "" {
		return "Calculate what?"
	}
	res, err := compute.Evaluate(message)
	if err != nil {
		return err.Error()
	}
	return strconv.FormatFloat(res, 'f', 4, 64)
}
func Time() string {
	return "It is " + time.Now().Format(time.Kitchen)
}

//SETTINGS
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
	Maxmessage   string            `json:"maxmessage"`
	Eightball    EightballSettings `json:"eightball"`
}

//helper functions
func randbool() bool {
	num := rand.Float64()
	if num > 0.5 {
		return true
	} else {
		return false
	}
}

func randint(low, high int) string {
	result := strconv.Itoa(rand.Intn((high+1)-low) + low)
	return result

}
