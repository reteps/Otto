# Otto
an imessage bot

otto is a group chat imessages OSX handler. It runs in applescript with a golang parser.  


imessage chat syntax:  
`otto COMMAND ARGS`  
`(otto) (say) (that's cool!)`


set it up:
+ clone this repo and give a star.
+ change the `ottohandler` and `settings.json` location 3x inside of `otto.applescript`
+ move `otto.applescript` to `~/Library/Application Scripts/com.apple.iChat` 
+ select `otto.applescript` as your imessage handler.

DISCLAIMER:
+ parts of `main.go` and `otto.applescript` were taken from
[Jared](https://github.com/ZekeSnider/Jared). All credit goes to its owner.
