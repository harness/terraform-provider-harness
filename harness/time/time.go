package time

import (
	"regexp"
	"strconv"
	"time"
)

type Time time.Time

// Return time as UTC
func (t Time) Time() time.Time {
	return time.Time(t).UTC()
}

// String returns t as a formatted string
func (t Time) String() string {
	return t.Time().String()
}

// MarshalJSON is used to convert the timestamp to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

var replaceQuotes = regexp.MustCompile(`"`)

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	data := string(b)

	// Field may have already been a string so we need to remove quotes before parsing.
	// https://harness.atlassian.net/browse/PL-24090
	data = replaceQuotes.ReplaceAllString(data, "")

	parsedInt, err := strconv.ParseInt(data, 10, 64)

	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(parsedInt, 0)
	return nil
}
