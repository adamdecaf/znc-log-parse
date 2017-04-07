# znc-log-parse

A znc log parser for quickly catching up on a room's conversation.

### Usage

```
$ znc-log-parse -help
Usage of bin/znc-log-parse-osx-amd64:
  -keep string
    	The message types to keep. Options: join, part, quit, rename, msg (default "msg")
  -path string
    	The filepath or directory of logfile(s) to parse. (default ".")
```

Typical usage for me involves catching up on conversation from yesterday. This could be done with something like (note, you'll need to manage timezones).

```
$ znc-log-parse -path '/opt/znc/moddata/log/adam/*/*/'$(date -d yesterday +"%Y-%m-%d").log
From /opt/znc/moddata/log/adam/$server/$room/2017-04-06.log
[00:10:19] <person> Hi all! I'm new here
```
