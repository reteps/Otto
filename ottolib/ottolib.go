package ottolib

import (
	"fmt"
	"time"
)

//--------------FUNCTIONS----------------------//
func Date(message, from string) string {
	t := time.Now()
	format := fmt.Sprintf("Today is %s, %s %d, %d", t.Weekday(), t.Month(), t.Day(), t.Year())
	return format
}

//--------------FUNCTIONS----------------------//

//----------------CONSTANTS-------------//
var sendlocation = "/Users/Peter/repos/Otto/SendText.applescript"
var funcmap = map[string]func(string, string) string{"date": Date}
var keywords = map[string]string{"hello": "hello there!", "version": "I am currently version 1.0", "date": "FUNCTION"}
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
