# Otto

otto is a imessage OSX chat handler. It runs in applescript with a golang parser.  


imessage chat syntax:  
`otto COMMAND ARGS`  
`(otto) (say) (that's cool!)`


features:
+ modular commands
+ customizeable settings
+ settings backup (if your settings file gets corrupted)

bugs:
+ There is only 1 chat room stored, so data from other chats can show up
+ Double messages - otto replies twice. This happens when you have unread messages in 1 chat, and otto is called in another.

set it up:
+ clone this repo and give a star.
+ change the `ottohandler` and `settings.json` location 2x inside of `otto.applescript`
+ move `otto.applescript` to `~/Library/Application Scripts/com.apple.iChat` 
+ select `otto.applescript` as your imessage handler.

contributing:
+ fork this repo
+ add your function to th
+ create your function inside `library.go`. Make sure it is one of these types:
  + `func(string, string) string` _calls a function with message and from as arguments_
  + `func(string) string` _calls a function with message as the argument_
  + `func() string` _calls a function that returns a string_
  + `string` _returns a message_

+ add your function and it's keyword to `ottomap` inside `library.go`
+ if your function takes arguments, make sure it handles edge cases including:
  + no arguments called
  + space is first character of argument
  + incorrect arguments

DISCLAIMER:
+ parts of `main.go` and `otto.applescript` were taken from
[Jared](https://github.com/ZekeSnider/Jared). All credit goes to its owner.
