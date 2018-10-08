package strftime

import (
	"strings"
	"time"
)

// Parse a string with the specified strftime layout
func Parse(layout, value string) (t time.Time, err error) {
	return time.Parse(Layout(layout), value)
}

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC; ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it against the Local location; ParseInLocation uses the given location.
func ParseInLocation(layout, value string, loc *time.Location) (t time.Time, err error) {
	return time.ParseInLocation(layout, value, loc)
}

// Format returns a textual representation of the time value formatted according to layout
func Format(layout string, t time.Time) (s string) {
	return t.Format(Layout(layout))
}

// Layout the given strftime layout to Go time format
func Layout(layout string) (fmt string) {
	var s strings.Builder

	for i := 0; i < len(layout); i++ {

		if layout[i] != '%' {
			s.WriteByte(layout[i])
			continue
		}

		i++ // comsume after %
		switch layout[i] {
		case 'a':
			// %a	Weekday as locale’s abbreviated name.	Mon
			s.WriteString("Mon")
		case 'A':
			// %A	Weekday as locale’s full name.	Monday
			s.WriteString("Monday")
		case 'd':
			// %d	Day of the month as a zero-padded decimal number.	30
			s.WriteString("02")
		case 'b':
			// %b	Month as locale’s abbreviated name.	Sep
			s.WriteString("Jan")
		case 'B':
			// %B	Month as locale’s full name.	September
			s.WriteString("January")
		case 'm':
			// %m	Month as a zero-padded decimal number.	09
			s.WriteString("01")
		case 'y':
			// %y	Year without century as a zero-padded decimal number.	13
			s.WriteString("06")
		case 'Y':
			// %Y	Year with century as a decimal number.	2013
			s.WriteString("2006")
		case 'H':
			// %H	Hour (24-hour clock) as a zero-padded decimal number.	07
			s.WriteString("15")
		case 'I':
			// %I	Hour (12-hour clock) as a zero-padded decimal number.	07
			s.WriteString("03")
		case 'p':
			// %p	Locale’s equivalent of either AM or PM.	AM
			s.WriteString("PM")
		case 'M':
			// %M	Minute as a zero-padded decimal number.	06
			s.WriteString("04")
		case 'S':
			// %S	Second as a zero-padded decimal number.	05
			s.WriteString("05")
		case 'f':
			// %f	Microsecond as a decimal number, zero-padded on the left.	000000
			s.WriteString(".000000")
		case 'z':
			// %z	UTC offset in the form +HHMM or -HHMM (empty string if the the object is naive).
			s.WriteString("-0700")
		case 'Z':
			// %Z	Time zone name (empty string if the object is naive).
			s.WriteString("MST")
		case 'c':
			// %c	Date and time representation.	Mon Jan 30 15:04:05 2006
			s.WriteString("Mon Jan 02 15:04:05 2006")
		case 'x':
			// %x	Date representation.	2006-01-02
			s.WriteString("2006-01-02")
		case 'X':
			// %X	Time representation.	15:04:05
			s.WriteString("15:04:05")
		case '%':
			// %%	A literal '%' character.	%
			s.WriteString("%")
		case 'e':
			// %e	Millisecond as a decimal number, zero-padded on the left.	000
			s.WriteString(".000")
		case 'n':
			// %n	Nanosecond as a decimal number, zero-padded on the left.	000000000
			s.WriteString(".000000000")
		case 't':
			// %t	ISO8601 appropriate date and time representation. 2006-01-02T15:04:05.000Z0700
			s.WriteString("2006-01-02T15:04:05.000Z0700")
		case 'u':
			// %u	RFC3339 appropriate date and time representation. 2006-01-02T15:04:05Z07:00
			s.WriteString(time.RFC3339)
		case 'v':
			// %v	RFC3339Nano date and time representation. 2006-01-02T15:04:05.999999999Z07:00
			s.WriteString(time.RFC3339Nano)
		}

	}

	return s.String()
}
