package duration

import (
	"strings"
	"time"
	"unicode"
)

func Parse(duration string) time.Duration {
	now := time.Now()
	add := Add(now, duration)
	return add.Sub(now)
}

func Add(t time.Time, duration string) time.Time {
	duration = strings.ToUpper(duration)
	for i, r := range duration {
		switch r {
		case 'P':
			_, t = addDate(duration[i+1:], t)
		case 'T':
			_, t = addTime(duration[i+1:], t)
		default:
			return t
		}
		break
	}
	return t
}

func addDate(duration string, t time.Time) (string, time.Time) {
	var num int
	var index int
	for i, r := range duration {
		switch r {
		case 'Y':
			t = t.AddDate(num, 0, 0)
			num = 0
		case 'M':
			t = t.AddDate(0, num, 0)
			num = 0
		case 'W':
			num *= 7
			t = t.AddDate(0, 0, num)
			num = 0
		case 'D':
			t = t.AddDate(0, 0, num)
			num = 0
		case 'T':
			return addTime(duration[i+1:], t)
		default:
			if unicode.IsDigit(r) {
				num *= 10
				num += int(r - '0')
			} else {
				return duration[index:], t
			}
		}
		index = i
	}
	return duration[index+1:], t
}

func addTime(duration string, t time.Time) (string, time.Time) {
	var num time.Duration
	var index int
	for i, r := range duration {
		switch r {
		case 'P':
			return addDate(duration[i:], t)
		case 'H':
			t = t.Add(num * time.Hour)
			num = 0
		case 'M':
			t = t.Add(num * time.Minute)
			num = 0
		case 'S':
			t = t.Add(num * time.Second)
			num = 0
		default:
			if unicode.IsDigit(r) {
				num *= 10
				num += time.Duration(r - '0')
			} else {
				return duration[index:], t
			}
		}
		index = i
	}
	return duration[index+1:], t
}
