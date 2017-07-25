# Otto

otto is a imessage OSX chat handler. It runs in applescript with a golang parser.  


imessage chat syntax:  
`otto COMMAND ARGS`  
`(otto) (say) (that's cool!)`


features:
+ modular commands
+ customizeable settings
+ settings backup (if your settings file gets corrupted)

bugs / annoyances:
+ There is only 1 chat room stored, so data from the other chat can show up
  + this is especially seen in `otto mock`
+ Double messages - otto replies twice. This happens when you have unread messages in 1 chat, and otto is called in another.
  + otto thinks it gets called twice

set it up:
+ clone this repo and give a star.
+ change the `ottohandler` and `settings.json` location 2x inside of `otto.applescript`
+ move `otto.applescript` to `~/Library/Application Scripts/com.apple.iChat` 
+ select `otto.applescript` as your imessage handler.

contributing:
+ clone this repo
+ add your function - make sure you do it in the format `function(message, from string) string`
+ add your function to the `funcList`
+ add your function to the `keywords`
+ if your function takes arguments, make sure it handles edge cases including:
  + no arguments called
  + space is first character of argument
  + incorrect arguments

DISCLAIMER:
+ parts of `main.go` and `otto.applescript` were taken from
[Jared](https://github.com/ZekeSnider/Jared). All credit goes to its owner.
