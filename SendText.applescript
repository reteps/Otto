#!/usr/bin/env osascript
on run argv
	tell application "Messages"
		set mychatid to item 2 of argv
		set mymessage to item 1 of argv
		set mybuddy to a reference to text chat id mychatid
		send mymessage to mybuddy
	end tell
end run
