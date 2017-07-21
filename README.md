# Otto
an imessage bot

otto is a group chat imessages OSX handler. It runs in applescript with a golang parser.  


imessage chat syntax:  
`otto COMMAND ARGS`  
`(otto) (say) (that's cool!)`


set it up:
+ clone this repo and give a star.
+ run `setup.sh` to set up `otto.applescript`
+ select `otto.applescript` as your imessage handler.


how it works:
+ `otto.applescript` receives the message and sends it to `ottohandler` in the form of `MESSAGE|~|WHO|~|GROUPID|~|SETTINGSLOCATION`
+ `ottohandler` reads the settings, then gets the correct text to send (if any) and sends the text


DISCLAIMER:
+ parts of `main.go` and `otto.applescript` were taken from
[Jared](https://github.com/ZekeSnider/Jared). All credit goes to its owner.
