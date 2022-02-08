package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 2022-02-08 15:17:40.32394 +0800 CST m=+0.000105074
	now := time.Now()
	p(now)

	// 2009-11-17 20:34:58.651387237 +0000 UTC
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	// November
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 107170h42m41.672552763s
	diff := now.Sub(then)
	p(diff)

	// 107170.7115757091
	// 6.430242694542546e+06
	// 3.8581456167255276e+08
	// 385814561672552763
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 2022-02-08 07:17:40.32394 +0000 UTC
	p(then.Add(diff))
	// 1997-08-27 09:52:16.978834474 +0000 UTC
	p(then.Add(-diff))
}
