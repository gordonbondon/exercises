# cron_parser1

Command line application that will parse [cron](https://man7.org/linux/man-pages/man5/crontab.5.html)
string and expand each field to show the time at which it will run.

## Usage

```
Usage of cron_parser1:
  -crontab string
    	Provide cron string in crontab format
```

### Example

```
~$ cron_parser1 -crontab "*/15 0 1,15 * 1-5 /usr/bin/find"

minute       [0 15 30 45]
hour         [0]
day of month [1 15]
month        [1 2 3 4 5 6 7 8 9 10 11 12]
day of week  [1 2 3 4 5]
command      /usr/bin/find
```

## Installation

Installation requires [`go`](https://golang.org/doc/install)

```
$ go install github.com/gordonbondon/exercises/takehome/cron_parser1@latest

$ cron_parser1 -help
```

## Testing

To run unit tests do this:

```
go test ./...
```
