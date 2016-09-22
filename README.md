# matt.daemon

```
77777777777777777777777777777777777~.,..:=77  7777777777777777777777777777777777
77777777777777777777777777777777,.............77 7777777777777777777777777777777
777777777777777777777777777 I+,...,.............=  77777777777777777777777777777
7777777777777777777777777 7:......................7 7777777777777777777777777777
77777777777777777777777777.,..,.,:.,...,..,........7  77777777777777777777777777
7777777777777777777777777=:,,,.:~:.~.,......,:......  77777777777777777777777777
77777777777777777777777 7,::=++??+===++~::=~=~:...,.~7  777777777777777777777777
7777777777777777777777 7=~:+????IIIIII??III??+++::~,.7  777777777777777777777777
7777777777777777777777 7I==+????IIIIIIIIIIII????=,=~:=  777777777777777777777777
7777777777777777777777 7I~????IIIIIIIIIIIIIII???+=,.,,   77777777777777777777777
7777777777777777777777 7+=+?IIIIIIIIIIIIIIIII???++:,.?     777777777777777777777
7777777777777777777777  I:II?IIIIIIIIIIIIIIII???+~:.:7     777777777777777777777
7777777777777777777777  7+II??IIIIIIIIIIIIIII???+:,,:7      7 777777777777777777
77777777777777777777  7I+?I+~==+?IIIIIIIIIIII???+~.,I7       7777777777777777777
7777777777777777777   7?+I????I?==??IIII?+++=??+++..77      77777777777777777777
77777777777777777     7I+I?+~=~.I+=?II==+??I?=~~+~.?I?    7   777777777777777777
77777777777777777     7IIII??+??????II+I?=~=,=+++~=I??7     77777777777777777777
777777777777          7?IIIIIIIII???I?+II?+++???+~~++7      77777777777777777777
77777777777           7I?I?I?IIII???I?+IIIIIII?+==I??7 77 7777777777777777777777
777777777             7II???III???III???IIIII?+++=?+7  7777777777777777777777777
777777777              7I?????I??==?+~~+IIII?++==??I  77777777777777777777777777
77777777               77??I?I??????II??III??+==~??7 777777777777777777777777777
7777777                  ???I????IIIIII?+???++==7    7 7777777777777777777777777
77777777                 7??II??+?++++~~+???++==7     77777777777777777777777777
77777777                 7???IIII?I?????II?++==I    7777777777777777777777777777
77777777                 77?+?IIII?+?III???====7    7777777777777777777777777777
77777  7                   ?????IIIIIIII?+====++7   7777777777777777777777777777
```
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

Example config file:

```
pidDir="/Users/park/matt.daemon/pids"
logDir="/Users/park/matt.daemon/logs"

[processes]

  [process.Program1]
  description="My first program"
  script="/path/to/bin/program1"

  [process.Program2]
  description="My second program"
  script="/path/to/bin/program2"
```
