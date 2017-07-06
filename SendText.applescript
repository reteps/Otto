#!/usr/bin/env osascript
on run argv
	tell application "Messages"
		set myid to item 2 of argv
		set mymessage to item 1 of argv
		set theBuddy to a reference to text chat id myid
		send mymessage to theBuddy
	end tell
end run
