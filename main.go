package main

import "fmt"

type interval struct {
	u    string
	nano uint64 // in nanoseconds
}

type system []interval

func mk(u string, q uint64, dt interval) interval {
	return interval{u: u, nano: dt.nano * q}
}

var (
	decade        = interval{u: "decade", nano: 10 * 365 * 24 * 60 * 60 * 1e9}
	year          = interval{u: "year", nano: 365 * 24 * 60 * 60 * 1e9}
	month         = interval{u: "month", nano: 30 * 24 * 60 * 60 * 1e9}
	week          = interval{u: "week", nano: 7 * 24 * 60 * 60 * 1e9}
	day           = interval{u: "day", nano: 24 * 60 * 60 * 1e9}
	hour          = interval{u: "hour", nano: 60 * 60 * 1e9}
	minute        = interval{u: "minute", nano: 60 * 1e9}
	second        = interval{u: "second", nano: 1e9}
	human  system = []interval{
		decade, year, month, week, day, hour, minute, second,
	}

	moon            = interval{u: "moon", nano: 15 * 24 * 60 * 60 * 1e9}
	halfmoon        = interval{u: "half-moon", nano: 8 * 24 * 60 * 60 * 1e9}
	mooner   system = []interval{
		moon, halfmoon, day, hour, minute, second,
	}

	dogyear = mk("dog year", 5, month)

	dogtime system = []interval{
		dogyear,
	}

	marstime system = []interval{mk("mars year", 687, day)}
)

func say(nano uint64, system system) []string {
	var out []string
	for _, s := range system {
		q := nano / s.nano
		if q != 0 {
			out = append(out,
				fmt.Sprintf("%d %ss", q, s.u),
			)
		}
		nano = nano % s.nano
	}
	return out
}

func main() {
	fmt.Println(
		say(10000*601*390*1e9, human),
	)
	fmt.Println(
		say(15*601*390*1e9, mooner),
	)
	fmt.Println(
		say(1005*601*390*1e9, dogtime),
	)
	fmt.Println(
		say(1005*601*390*1e9, marstime),
	)
}
