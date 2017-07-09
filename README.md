# Otto
an imessage bot

otto is a group chat imessages OSX handler. It runs in applescript with a golang parser.


imessage syntax:  
`otto COMMAND ARGS`  
`(otto) (say) (that's cool!)`


features to watch for / things that need to be added
+ ~easy edit commands~
+ `help` and `say` command
+ read `sendlocation` from a file (`settings.txt`?) and also functionality turning on and off
+ only select certain groups to run in
+ duplication glitch (sends message twice)
+ 


set it up:
+ place `otto.applescript` inside of `~/Library/Application Scripts/com.apple.iChat`
+ select `otto.applescript` as the applescript handler in `imessages > preferences`
+ replace the location of `ottohandler` 3x inside of `otto.applescript`
  + line `19`,`31` and `43`
+ change the `sendlocation` inside of `settings.txt` *not possible yet*
+ in theory, this should work

how it works:
+ `otto.applescript` receives the message and sends it to `ottohandler` in the form of `MESSAGE|~|WHO|~|GROUPID`
+ `ottohandler` gets the correct text to send (if any) and calls `SendText.applescript` with a message and groupid
+ `SendText.applescript` sends the text.

contributing:
+ create an issue with your wanted feature 
+ write a command for `ottolib` and send a Pull Request

DISCLAIMER:
+ parts of `SendText.applescript` and `otto.applescript` were taken from
[Jared](https://github.com/ZekeSnider/Jared). All credit goes to its owner.
