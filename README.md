# Go strftime
Package strftime provides strftime like functionality as an alternative format for parsing date and time in Go.

It doesn't follow the full specification for strftime and align with what is supported by the Go time package adding a few extensions.


## Examples
```go
package main

import (
	"fmt"

	"github.com/brunotm/rexon/value/strftime"
)

func main() {
	st, err := strftime.Parse("%c %Z", "Mon Oct 08 17:48:51 2018 GMT")
	if err != nil {
		panic(err)
	}
	fmt.Println(st)

	s := strftime.Format("%t", st)
	fmt.Println(s)
}
```


## Supported directives and extensions

| Code   | Meaning    | Example|
|-----------|---------|-------|
|%a|   Weekday as locale’s abbreviated name.|   Mon|
|%A|   Weekday as locale’s full name.|  Monday|
|%d|   Day of the month as a zero-padded decimal number.|       30|
|%b|   Month as locale’s abbreviated name.|     Sep|
|%B|   Month as locale’s full name.|    September|
|%m|   Month as a zero-padded decimal number.|  09|
|%y|   Year without century as a zero-padded decimal number.|   13|
|%Y|   Year with century as a decimal number.|  2013|
|%H|   Hour (24-hour clock) as a zero-padded decimal number.|   07|
|%I|   Hour (12-hour clock) as a zero-padded decimal number.|   07|
|%p|   Locale’s equivalent of either AM or PM.| AM|
|%M|   Minute as a zero-padded decimal number.| 06|
|%S|   Second as a zero-padded decimal number.| 05|
|%f|   Microsecond as a decimal number, zero-padded on the left.|       000000|
|%z|   UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).| -700|
|%Z|   Time zone name (empty string if the object is naive).| MST|
|%c|   Date and time representation.|   Mon Jan 30 15:04:05 2006|
|%x|   Date representation.|    2006-01-02|
|%X|   Time representation.|    15:04:05.000Z0700|
|%%|   A literal '%' character.|        %|
|%e|   Extension - Millisecond as a decimal number, zero-padded on the left.|       000|
|%n|   Extension - Nanosecond as a decimal number, zero-padded on the left.|        000000000|
|%t|   Extension - ISO8601 appropriate date and time representation.| 2006-01-02T15:04:05.000Z0700|
|%u|   Extension - RFC3339 appropriate date and time representation.| 2006-01-02T15:04:05Z07:00|
|%v|   Extension - RFC3339Nano date and time representation.| 2006-01-02T15:04:05.999999999Z07:00|

---------------------------
Written by Bruno Moura <brunotm@gmail.com>