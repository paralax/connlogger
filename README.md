# connlogger
simple server-side connection logger, useful as a simple honeypot

# Usage

See `start.sh` for a simple startup script (e.g. this is what I use for my RDP honeypot), but basically:
```
$ connlogger -port N -logfile /path/to/logfile
```

# Log format

Logs are line delineated, timestamp in format `YYYY/MM/DD HH:MM:SS` and then a report where the connection came from, including client IP and source port. Example:
```
2019/01/06 14:33:32 Connection received from:  115.236.70.116:53951
```
