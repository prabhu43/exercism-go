// Package gigasecond is a library to manipulate time with gigasecond
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
