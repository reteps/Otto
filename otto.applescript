using terms from application "Messages"
	on message sent theMessage with eventDescription
	end message sent
	
	on message received theText from theBuddy with eventDescription
		set getname to name of theBuddy as text
		try
			set myresult to get id of theBuddy
		on error errMsg
			set errMsgParts to splitText(errMsg, "\"")
			set errCount to count of errMsgParts
			set myresult to item (errCount - 1) of errMsgParts
		end try
		do shell script ("/usr/local/bin/python3 /Users/Peter/repos/imessage/message \"" & theText & "|~|" & getname & "|~|" & myresult & "\"")
	end message received
	
	on chat room message received theText with eventDescription from theBuddy for theChat
		set getname to name of theBuddy as text
		try
			set myresult to get id of theChat
		on error errMsg
			set errMsgParts to splitText(errMsg, "\"")
			set errCount to count of errMsgParts
			set myresult to item (errCount - 1) of errMsgParts
		end try
		do shell script ("/usr/local/bin/python3 /Users/Peter/repos/imessage/message \"" & theText & "|~|" & getname & "|~|" & myresult & "\"")
	end chat room message received
	
	on active chat message received theText with eventDescription from theBuddy for theChat
		set getname to name of theBuddy as text
		try
			set myresult to get id of theChat
		on error errMsg
			set errMsgParts to splitText(errMsg, "\"")
			set errCount to count of errMsgParts
			set myresult to item (errCount - 1) of errMsgParts
		end try
		do shell script ("/usr/local/bin/python3 /Users/Peter/repos/imessage/message \"" & theText & "|~|" & getname & "|~|" & myresult & "\"")
	end active chat message received
	
	on addressed message received theText with eventDescription from theBuddy for theChat
		set getname to name of theBuddy as text
		try
			set myresult to get id of theChat
		on error errMsg
			set errMsgParts to splitText(errMsg, "\"")
			set errCount to count of errMsgParts
			set myresult to item (errCount - 1) of errMsgParts
		end try
		do shell script ("/usr/local/bin/python3 /Users/Peter/repos/imessage/message \"" & theText & "|~|" & getname & "|~|" & myresult & "\"")
	end addressed message received
	
	on received text invitation with eventDescription
	end received text invitation
	
	on received audio invitation theText from theBuddy for theChat with eventDescription
	end received audio invitation
	
	on received video invitation theText from theBuddy for theChat with eventDescription
	end received video invitation
	
	on buddy authorization requested with eventDescription
	end buddy authorization requested
	
	on addressed chat room message received with eventDescription
	end addressed chat room message received
	
	on login finished with eventDescription
	end login finished
	
	on logout finished with eventDescription
	end logout finished
	
	on buddy became available with eventDescription
	end buddy became available
	
	on buddy became unavailable with eventDescription
	end buddy became unavailable
	
	on received file transfer invitation theFileTransfer with eventDescription
	end received file transfer invitation
	
	on av chat started with eventDescription
	end av chat started
	
	on av chat ended with eventDescription
	end av chat ended
	
	on completed file transfer with eventDescription
	end completed file transfer
	
end using terms from

on splitText(sourceText, textDelimiter)
	set AppleScript's text item delimiters to {textDelimiter}
	set messageParts to (every text item in sourceText) as list
	set AppleScript's text item delimiters to ""
	return messageParts
end splitText
