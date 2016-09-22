# matt.daemon

A very skinny process daemon built in Go. At the moment only manages executable binaries, but could easy be modified to support script execution through `/bin/sh`.

This is a contribution to *Derivco Hackathon#6* and has been built in around 6 hours total.

`matt.daemon` can:
- spawn and restart processes that crash; os/exit > 0
- ignore processes that terminate on purpose; os/exit 0
- expose an rpc api for killing/starting processes over tcp.

`matt.daemon` cannot:
- run shell scripts
- make movies

Additionally there is a controller called `mattd` which uses the rpc api to control
proccesses in ``matt.daemon``. `mattd`lives in a separate project and will be uploaded
as soon as i have cleaned up the code :)

The code is a bit clunky and contains 0 tests because Hackathon.

Cheers!
