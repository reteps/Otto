package ottolib

import (
	"fmt"
	"math/rand"
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
	message := "Commands include:" + strings.Join(keys, ",")
	return message
}
func Random(message, from string) string {
	return "work in progress"
}
func ReadValues() (string, int) {
	sett, _ := ioutil.ReadFile(settings)
	var data []interface{}
	err := json.Unmarshal(sett, &data)
	return data[0].(string), data[1].(int)
}

//--------------FUNCTIONS----------------------//

//----------------CONSTANTS-------------//
var (
	settings = "/Users/Peter/go/src/Otto/settings.json"
	funcmap  = map[string]func(string, string) string{"date": Date, "help": Help, "random": Random}
	keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.0",
		"date": "FUNCTION", "help": "FUNCTION", "random": "FUNCTION"}
	errormessage = "Sorry, I don't understand"
)

//----------------CONSTANTS-------------//

//--------------DO NOT MODIFY------------------------//

func FuncMap() map[string]func(string, string) string {
	return funcmap
}
func Keywords() map[string]string {
	return keywords
}
func Errormessage() string {
	return errormessage
}

//--------------DO NOT MODIFY-----------------------//
