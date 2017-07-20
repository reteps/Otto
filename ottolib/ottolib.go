package ottolib

import (
	"fmt"
	"strings"
	"time"
	"math/rand"
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
func Mock(message, from string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	text :=	Readvalue("2nd to last text")
	ftext := ""
	for _, v := range text {
		sv := string(v)
		num := rand.Float64()
		if num > 0.5 {
			ftext += strings.ToUpper(sv)
		} else {
			ftext += strings.ToLower(sv)
		}
	}
	return ftext
}		
func ReadValues() {
	sett, _ := ioutil.ReadFile(settings)
	var data interface{}
	err := json.Unmarshal(sett, &data)
		
//--------------FUNCTIONS----------------------//

//----------------CONSTANTS-------------//
var settings = "/Users/Peter/repos/Otto/settings.json"
var sendlocation = "/Users/Peter/repos/Otto/SendText.applescript"
var funcmap = map[string]func(string, string) string{"date": Date,"help":Help,"mock":Mock}
var keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.0", 
"date": "FUNCTION","help":"FUNCTION","turn":"FUNCTION","tell":"FUNCTION","mock":"FUNCTION"}
var errormessage = "Sorry, I don't understand"

//----------------CONSTANTS-------------//

//--------------DO NOT MODIFY------------------------//

func FuncMap() map[string]func(string, string) string {
	return funcmap
}
func Keywords() map[string]string {
	return keywords
}
func SendLocation() string {
	return sendlocation
}
func Errormessage() string {
	return errormessage
}

//--------------DO NOT MODIFY-----------------------//
